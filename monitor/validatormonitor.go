package monitor

import (
	"context"
	"fmt"
	"log"

	"github.com/crosschained-io/crosschained-service/contract"
	"github.com/crosschained-io/crosschained-service/registry"
	"github.com/crosschained-io/crosschained-service/types"
	"github.com/pkg/errors"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type (
	ValidatorMonitor struct {
		MonitorBase
		validator       common.Address
		assetUpperBound *registry.AssetUpperBoundRegistry
	}
)

func NewValidatorMonitor(
	chainAPI string,
	validator common.Address,
	tubeAddr common.Address,
	startHeight, heightToConfirm, batchSize uint64,
	assetRegistry AssetRegistry,
	assetUpperBound *registry.AssetUpperBoundRegistry,
	recorder *Recorder,
) (*ValidatorMonitor, error) {
	client, err := ethclient.DialContext(context.Background(), chainAPI)
	if err != nil {
		return nil, err
	}
	tubeContract, err := contract.NewERC20Tube(tubeAddr, client)
	if err != nil {
		return nil, errors.Wrap(err, "failed create erc20 tube")
	}
	verifierContractAddr, err := tubeContract.Verifier(nil)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get verifier address")
	}
	log.Println("verifier:", verifierContractAddr)
	query := ethereum.FilterQuery{
		Addresses: []common.Address{tubeAddr, verifierContractAddr},
	}
	id, err := tubeContract.TubeID(nil)
	if err != nil {
		return nil, err
	}
	patrolSize := batchSize * 6
	if patrolSize > 5000 {
		patrolSize = 5000 / batchSize * batchSize
	}

	return &ValidatorMonitor{
		MonitorBase: MonitorBase{
			client:        client,
			assetRegistry: assetRegistry,
			confirmHeight: heightToConfirm,
			startHeight:   startHeight,
			batchSize:     batchSize,
			patrolSize:    patrolSize,
			tubeID:        id.Uint64(),
			recorder:      recorder,
			query:         query,
			parsers: map[common.Address]*contract.EventParser{
				verifierContractAddr: contract.NewVerifierEventParser(),
				tubeAddr:             contract.NewERC20TubeEventParser(),
			},
		},
		validator:       validator,
		assetUpperBound: assetUpperBound,
	}, nil
}

func (monitor *ValidatorMonitor) Start(ctx context.Context) error {
	return nil
}

func (monitor *ValidatorMonitor) Stop(ctx context.Context) error {
	monitor.client.Close()
	return nil
}

func (monitor *ValidatorMonitor) Run() error {
	verifierUpToDate, tipHeight, nextSyncedHeight, events, err := monitor.sync()
	if err != nil {
		return err
	}

	objs := []interface{}{}
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
			deposit, assetID, active, err := monitor.createDeposit(e)
			switch errors.Cause(err) {
			case nil:
				task := types.NewSignatureTask(deposit)
				if !active {
					task.Status = types.Inactive
				} else {
					if monitor.assetUpperBound != nil {
						upperBound, err := monitor.assetUpperBound.UpperBound(assetID)
						if err != nil {
							fmt.Printf("failed to fetch upper bound for asset, %d\n", assetID)
							task.Status = types.InReview
						} else {
							if upperBound == nil || deposit.Amount.BigInt().Cmp(upperBound) <= 0 {
								task.Status = types.New
							} else {
								fmt.Printf("manual review is needed for %s\n", task.ID)
								task.Status = types.InReview
							}
						}
					} else {
						task.Status = types.New
					}
				}
				objs = append(objs, deposit, task)
			case registry.ErrAssetNotFound:
				log.Printf("ignore transfer from %d to %d whose asset %s is not found, %+v\n", monitor.tubeID, e.TargetTubeID, e.Token, err)
			default:
				return err
			}
		case *contract.VerifierValidatorAdded:
			objs = append(objs, monitor.createActiveValidator(e))
		case *contract.VerifierValidatorRemoved:
			objs = append(objs, monitor.createInactiveValidator(e))
		}
	}

	if err := monitor.process(tipHeight, nextSyncedHeight, verifierUpToDate, objs); err != nil {
		return errors.Wrap(err, "failed to process")
	}
	isValidator, err := monitor.recorder.IsValidator(monitor.tubeID, monitor.validator)
	if err != nil {
		return err
	}
	if isValidator {
		return monitor.recorder.UpdateSignatureTasks(monitor.tubeID, monitor.validator, types.Ready)
	}

	return monitor.recorder.UpdateSignatureTasks(monitor.tubeID, monitor.validator, types.OutOfTheDuty)
}
