package monitor

import (
	"context"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"

	"github.com/crosschained-io/crosschained-service/contract"
	"github.com/crosschained-io/crosschained-service/registry"
	"github.com/crosschained-io/crosschained-service/types"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
)

type (
	RelayerMonitor struct {
		MonitorBase
		// recorder         *Recorder
		testimonyFetcher TestimonyDAO
		relayAll         bool
		tubeContract     *contract.ERC20Tube
	}
)

func NewRelayerMonitorForAllTasks(
	chainAPI string,
	tubeAddr common.Address,
	startHeight, heightToConfirm, batchSize uint64,
	assetRegistry AssetRegistry,
	recorder *Recorder,
	testimonyFetcher TestimonyDAO,
) (*RelayerMonitor, error) {
	client, err := ethclient.DialContext(context.Background(), chainAPI)
	if err != nil {
		return nil, err
	}
	monitor, err := newRelayerMonitor(
		client,
		tubeAddr,
		startHeight, heightToConfirm, batchSize,
		recorder,
		assetRegistry,
		testimonyFetcher,
	)
	if err != nil {
		return nil, err
	}
	monitor.relayAll = true
	return monitor, err
}

func NewRelayerMonitorForRouter(
	chainAPI string,
	tubeAddr common.Address,
	routerAddr common.Address,
	startHeight, heightToConfirm, batchSize uint64,
	assetRegistry AssetRegistry,
	recorder *Recorder,
	testimonyFetcher TestimonyDAO,
) (*RelayerMonitor, error) {
	client, err := ethclient.DialContext(context.Background(), chainAPI)
	if err != nil {
		return nil, err
	}
	monitor, err := newRelayerMonitor(
		client,
		tubeAddr,
		startHeight, heightToConfirm, batchSize,
		recorder,
		assetRegistry,
		testimonyFetcher,
	)
	if err != nil {
		return nil, err
	}
	monitor.query.Addresses = append(monitor.query.Addresses, routerAddr)
	monitor.parsers[routerAddr] = contract.NewERC20RouterEventParser()

	return monitor, err
}

func newRelayerMonitor(
	client *ethclient.Client,
	tubeAddr common.Address,
	startHeight, heightToConfirm, batchSize uint64,
	recorder *Recorder,
	assetRegistry AssetRegistry,
	testimonyFetcher TestimonyDAO,
) (*RelayerMonitor, error) {
	if testimonyFetcher == nil {
		return nil, errors.New("testimony fetcher is nil")
	}
	query := ethereum.FilterQuery{
		Addresses: []common.Address{tubeAddr},
	}
	parsers := map[common.Address]*contract.EventParser{}
	parsers[tubeAddr] = contract.NewERC20TubeEventParser()

	tubeContract, err := contract.NewERC20Tube(tubeAddr, client)
	if err != nil {
		return nil, err
	}
	verifierContractAddr, err := tubeContract.Verifier(nil)
	if err != nil {
		return nil, err
	}
	fmt.Println("verifier:", verifierContractAddr)
	parsers[verifierContractAddr] = contract.NewVerifierEventParser()
	query.Addresses = append(query.Addresses, verifierContractAddr)
	id, err := tubeContract.TubeID(nil)
	if err != nil {
		return nil, err
	}

	return &RelayerMonitor{
		MonitorBase: MonitorBase{
			client:        client,
			confirmHeight: heightToConfirm,
			startHeight:   startHeight,
			batchSize:     batchSize,
			tubeID:        id.Uint64(),
			recorder:      recorder,
			assetRegistry: assetRegistry,
			query:         query,
			parsers:       parsers,
		},
		testimonyFetcher: testimonyFetcher,
		tubeContract:     tubeContract,
	}, nil
}

func (monitor *RelayerMonitor) Start(ctx context.Context) error {
	return nil
}

func (monitor *RelayerMonitor) Stop(ctx context.Context) error {
	monitor.client.Close()
	return nil
}

func (monitor *RelayerMonitor) TubeID() uint64 {
	return monitor.tubeID
}

func (monitor *RelayerMonitor) TaskID(_srcTubeID *big.Int, _nonce *big.Int, _token common.Address, _recipient common.Address, _amount *big.Int) (common.Hash, error) {
	return monitor.tubeContract.GenKey(nil, _srcTubeID, _nonce, _token, _recipient, _amount)
}

func (monitor *RelayerMonitor) Run() error {
	verifierUpToDate, tipHeight, nextSyncedHeight, events, err := monitor.sync()
	if err != nil {
		return err
	}
	objs := []interface{}{}
	taskMap := map[common.Hash]*types.RelayTask{}
	for _, event := range events {
		if event.Height > nextSyncedHeight {
			switch event.Parsed.(type) {
			case *contract.VerifierValidatorAdded, *contract.VerifierValidatorRemoved:
				verifierUpToDate = false
			}
			continue
		}
		switch e := event.Parsed.(type) {
		case *contract.ERC20TubeReceipt:
			deposit, _, active, err := monitor.createDeposit(e)
			switch errors.Cause(err) {
			case nil:
				task := types.NewRelayTask(deposit)
				if _, exist := taskMap[e.Raw.TxHash]; exist {
					return errors.Errorf("multiple deposits in one transaction %s", deposit.ID)
				}
				switch {
				case !active:
					task.Status = types.Inactive
				case monitor.relayAll:
					task.Status = types.New
				}
				taskMap[e.Raw.TxHash] = task
				objs = append(objs, deposit, task)
			case registry.ErrAssetNotFound:
				log.Printf("ignore transfer from %d to %d whose asset %s is not found, %+v\n", monitor.tubeID, e.TargetTubeID, e.Token, err)
			default:
				return err
			}
		case *contract.ERC20TubeSettled:
			objs = append(objs, monitor.createWithdrawal(e))
		case *contract.VerifierValidatorAdded:
			objs = append(objs, monitor.createActiveValidator(e))
		case *contract.VerifierValidatorRemoved:
			objs = append(objs, monitor.createInactiveValidator(e))
		case *contract.VerifierPaused:
			objs = append(objs, monitor.createPausedVerifierStatus(e))
		case *contract.VerifierUnpaused:
			objs = append(objs, monitor.createUnpausedVerifierStatus(e))
		case *contract.ERC20TubeRouterRelayFeeReceipt:
			if monitor.relayAll {
				return errors.New("unexpected event in relay all mode")
			}
			txHash := e.Raw.TxHash
			task, exist := taskMap[txHash]
			if !exist {
				return errors.Errorf("failed to find deposit for task in transaction %s", txHash)
			}
			if task.Status == types.Unknown {
				task.Status = types.New
				task.Fee = task.Fee.Add(decimal.NewFromBigInt(e.Amount, 0))
			}
		}
	}
	if err := monitor.process(tipHeight, nextSyncedHeight, verifierUpToDate, objs); err != nil {
		return errors.Wrap(err, "failed to process")
	}
	if err := monitor.recorder.BackupInExplorerDB(); err != nil {
		log.Printf("failed to backup into explorer db, %+v\n", err)
	}
	if err := monitor.recorder.SettleDeposits(monitor.tubeID); err != nil {
		return errors.Wrapf(err, "failed to settle deposits for monitor %d", monitor.tubeID)
	}

	return monitor.collectTestimonies()
}

func (monitor *RelayerMonitor) collectTestimonies() error {
	validators, err := monitor.recorder.Validators(monitor.tubeID)
	switch errors.Cause(err) {
	case nil:
	case types.ErrValidatorIsOutOfDate:
		return nil
	default:
		return err
	}
	if len(validators) == 0 {
		return nil
	}
	tasks, err := monitor.recorder.WaitingForTestimonies(monitor.tubeID, 0, 100)
	if err != nil {
		return err
	}
	for _, task := range tasks {
		if err := monitor.collectTestimoniesForOne(task, validators); err != nil {
			log.Printf("failed to collect signatures for %s: %+v\n", task.ID, err)
		}
	}
	return nil
}

func (monitor *RelayerMonitor) collectTestimoniesForOne(task *types.RelayTask, validators []common.Address) error {
	id := common.HexToHash(task.ID)
	ts := map[common.Address][]byte{}
	numOfRejected := 0
	for _, testimony := range task.Testimonies {
		signature, err := hex.DecodeString(testimony.Signature)
		if err != nil {
			return err
		}
		ts[common.HexToAddress(testimony.Validator)] = signature
	}
	validSignatures := [][]byte{}
	for _, validator := range validators {
		sig, ok := ts[validator]
		if !ok || len(sig) == 0 {
			var err error
			sig, err = monitor.testimonyFetcher.Testimony(validator, id)
			if err != nil {
				log.Printf(
					"failed to fetch testimony from %s for %s {from: (%d, %s, %s), to: (%d, %s, %s), %s\n",
					validator,
					id,
					task.Deposit.SourceTubeID,
					task.Deposit.SourceToken,
					task.Deposit.Sender,
					task.Deposit.DestinationTubeID,
					task.Deposit.DestinationToken,
					task.Deposit.Recipient,
					err,
				)
				continue
			}
			if err := monitor.recorder.AddTestimony(id, validator, sig); err != nil {
				return err
			}
		}
		if len(sig) > 0 {
			validSignatures = append(validSignatures, sig)
		} else {
			numOfRejected++
		}
	}
	if numOfRejected*3 >= len(validators) {
		if err := monitor.recorder.MarkAsRejected(id); err != nil {
			return err
		}
		return nil
	}
	if len(validSignatures)*3 > len(validators)*2 {
		if err := monitor.recorder.MarkAsReady(id); err != nil {
			return err
		}
	}
	return nil
}
