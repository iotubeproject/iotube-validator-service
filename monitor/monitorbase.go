package monitor

import (
	"context"
	"log"
	"math/big"

	"github.com/crosschained-io/crosschained-service/contract"
	"github.com/crosschained-io/crosschained-service/types"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
)

const (
	ValidatorMode MonitorMode = iota + 1
	RelayRouterTask
	RelayAllTask
)

type (
	Monitor interface {
		TubeID() uint64
		Run() error
		Start(context.Context) error
		Stop(context.Context) error
	}

	MonitorMode uint8

	MonitorConfig struct {
		ChainAPI           string `json:"chainAPI" yaml:"chainAPI"`
		ContractAddr       string `json:"contractAddr" yaml:"contractAddr"`
		RouterContractAddr string `json:"routerContractAddr" yaml:"routerContractAddr"`
		BatchSize          uint16 `json:"batchSize" yaml:"batchSize"`
		HeightToConfirm    uint8  `json:"heightToConfirm" yaml:"heightToConfirm"`
		StartHeight        uint64 `json:"startHeight" yaml:"startHeight"`
		Mode               uint8  `json:"mode" yaml:"mode"`
		Validator          string `json:"validator" yaml:"validator"`
	}

	AssetRegistry interface {
		AssetOnTube(uint64, common.Address, uint64) (uint64, common.Address, bool, error)
	}

	TestimonyDAO interface {
		Testimony(common.Address, common.Hash) ([]byte, error)
	}

	MonitorBase struct {
		client           *ethclient.Client
		assetRegistry    AssetRegistry
		startHeight      uint64
		confirmHeight    uint64
		batchSize        uint64
		patrolSize       uint64
		lastPatrolHeight uint64
		tubeID           uint64
		recorder         *Recorder
		query            ethereum.FilterQuery
		parsers          map[common.Address]*contract.EventParser
	}
)

func (monitor *MonitorBase) TubeID() uint64 {
	return monitor.tubeID
}

func (monitor *MonitorBase) taskInfoFetcher(
	nonce uint64,
	token common.Address,
	targetTubeID uint64,
	recipient common.Address,
	amount *big.Int,
) (common.Hash, uint64, common.Address, bool, error) {
	assetID, dstAsset, active, err := monitor.assetRegistry.AssetOnTube(monitor.tubeID, token, targetTubeID)
	if err != nil {
		return common.Hash{}, 0, common.Address{}, false, err
	}
	if assetID == 0 {
		return common.Hash{}, 0, dstAsset, false, nil
	}

	return crypto.Keccak256Hash(
		math.U256Bytes(new(big.Int).SetUint64(monitor.tubeID)),
		math.U256Bytes(new(big.Int).SetUint64(nonce)),
		math.U256Bytes(new(big.Int).SetUint64(targetTubeID)),
		dstAsset.Bytes(),
		math.U256Bytes(amount),
		recipient.Bytes(),
	), assetID, dstAsset, active, nil
}

func (monitor *MonitorBase) fetchEvents(startHeight, endHeight uint64) (uint64, []*contract.Event, error) {
	tip, err := monitor.client.BlockNumber(context.Background())
	if err != nil {
		return 0, nil, errors.Wrap(err, "failed to get tip height")
	}
	if endHeight > tip {
		endHeight = tip
	}
	if endHeight < startHeight {
		return tip, nil, nil
	}
	monitor.query.FromBlock = new(big.Int).SetUint64(startHeight)
	monitor.query.ToBlock = new(big.Int).SetUint64(endHeight)
	logs, err := monitor.client.FilterLogs(context.Background(), monitor.query)
	if err != nil {
		return 0, nil, err
	}
	events := make([]*contract.Event, 0, len(logs))
	for _, log := range logs {
		parser, ok := monitor.parsers[log.Address]
		if !ok {
			return tip, nil, errors.Errorf("unexpected contract address %s", log.Address)
		}
		event, err := parser.Parse(log)
		if err != nil {
			return tip, nil, err
		}
		if event != nil {
			events = append(events, event)
		}
	}

	return tip, events, nil
}

func (monitor *MonitorBase) createDeposit(e *contract.ERC20TubeReceipt) (*types.Deposit, uint64, bool, error) {
	id, assetID, dstAsset, active, err := monitor.taskInfoFetcher(
		e.Nonce.Uint64(),
		e.Token,
		e.TargetTubeID.Uint64(),
		e.Recipient,
		e.Amount,
	)
	if err != nil {
		return nil, 0, false, errors.Wrap(err, "failed to generate task id")
	}
	deposit := types.NewDeposit(e)
	deposit.ID = id.Hex()
	deposit.SourceTubeID = monitor.tubeID
	deposit.DestinationToken = dstAsset.Hex()

	return deposit, assetID, active, nil
}

func (monitor *MonitorBase) createWithdrawal(e *contract.ERC20TubeSettled) *types.Withdrawal {
	withdrawal := types.NewWithdrawal(e)
	withdrawal.TubeID = monitor.tubeID
	return withdrawal
}

func (monitor *MonitorBase) createActiveValidator(e *contract.VerifierValidatorAdded) *types.Validator {
	validator := types.NewActiveValidator(e)
	validator.TubeID = monitor.tubeID
	return validator
}

func (monitor *MonitorBase) createInactiveValidator(e *contract.VerifierValidatorRemoved) *types.Validator {
	validator := types.NewInactiveValidator(e)
	validator.TubeID = monitor.tubeID
	return validator
}

func (monitor *MonitorBase) createPausedVerifierStatus(e *contract.VerifierPaused) *types.VerifierStatus {
	verifier := types.NewPausedVerifier(e)
	verifier.TubeID = monitor.tubeID
	return verifier
}

func (monitor *MonitorBase) createUnpausedVerifierStatus(e *contract.VerifierUnpaused) *types.VerifierStatus {
	verifier := types.NewUnpausedVerifier(e)
	verifier.TubeID = monitor.tubeID
	return verifier
}

func (monitor *MonitorBase) sync() (bool, uint64, uint64, []*contract.Event, error) {
	syncedHeight, err := monitor.recorder.TubeSyncedHeight(monitor.tubeID)
	if err != nil {
		return false, 0, 0, nil, errors.Wrapf(err, "failed to get synced height for tube %d", monitor.tubeID)
	}
	startHeight := syncedHeight + 1
	if startHeight < monitor.startHeight {
		startHeight = monitor.startHeight
	}
	nextSyncedHeight := startHeight + monitor.batchSize - 1
	if monitor.patrolSize > 0 {
		if monitor.lastPatrolHeight == 0 {
			if startHeight > monitor.patrolSize {
				monitor.lastPatrolHeight = startHeight - monitor.patrolSize
				if monitor.lastPatrolHeight < monitor.startHeight {
					monitor.lastPatrolHeight = monitor.startHeight
				}
			}
		} else {
			if nextSyncedHeight >= monitor.lastPatrolHeight+monitor.patrolSize {
				monitor.lastPatrolHeight, startHeight = startHeight, monitor.lastPatrolHeight
			}
		}
	}
	tipHeight, events, err := monitor.fetchEvents(startHeight, nextSyncedHeight)
	if err != nil {
		return false, 0, 0, nil, err
	}
	log.Printf("fetch %d logs in [%d, %d] for tube %d (tip %d)\n", len(events), startHeight, nextSyncedHeight, monitor.tubeID, tipHeight)
	isUpToDate := nextSyncedHeight >= tipHeight
	confirmedHeight := tipHeight - monitor.confirmHeight
	if nextSyncedHeight > confirmedHeight {
		nextSyncedHeight = confirmedHeight
	}
	if nextSyncedHeight < startHeight {
		return false, 0, 0, nil, errors.New("confirmed height is behind latest synced height")
	}

	return isUpToDate, tipHeight, nextSyncedHeight, events, nil
}

func (monitor *MonitorBase) process(tipHeight, syncedHeight uint64, upToDate bool, objs []interface{}) error {
	if err := monitor.recorder.Begin(); err != nil {
		return errors.Wrap(err, "failed to start a new transaction")
	}
	defer monitor.recorder.Rollback()
	for _, obj := range objs {
		if err := monitor.recorder.Create(obj); err != nil {
			return err
		}
	}
	if err := monitor.recorder.Upsert(&types.Tube{
		ID:               monitor.tubeID,
		TipHeight:        tipHeight,
		SyncedHeight:     syncedHeight,
		VerifierUpToDate: upToDate,
	}); err != nil {
		return err
	}
	return monitor.recorder.Commit()
}
