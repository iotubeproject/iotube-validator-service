package registry

import (
	"context"
	"log"
	"net/http"
	"sync"

	"github.com/crosschained-io/crosschained-service/contract"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

const (
	NotRegistered = iota
	CONTRACT
	GRPC_API
	HTTP
)

var ErrUnknownValidator = errors.New("unknown validator")

type (
	ValidatorRegistry struct {
		batchSize           uint64
		startHeight         uint64
		client              *ethclient.Client
		contractAddr        common.Address
		db                  *gorm.DB
		dirty               bool
		filterer            *contract.ValidatorRegistryFilterer
		mutex               sync.RWMutex
		validatorInfos      map[common.Address]ValidatorInfo
		localValidatorInfos map[common.Address]ValidatorInfo
		mode                int
	}
)

func NewValidatorRegistry(
	cfg ValidatorRegistryConfig,
	client *ethclient.Client,
	db *gorm.DB,
) (*ValidatorRegistry, error) {
	contractAddr := common.HexToAddress(cfg.ContractAddr)
	filterer, err := contract.NewValidatorRegistryFilterer(contractAddr, client)
	if err != nil {
		return nil, err
	}
	localValidatorInfos := map[common.Address]ValidatorInfo{}
	for validator, info := range cfg.Validators {
		localValidatorInfos[common.HexToAddress(validator)] = info
	}

	return &ValidatorRegistry{
		batchSize:           cfg.BatchSize,
		startHeight:         cfg.StartHeight,
		client:              client,
		contractAddr:        contractAddr,
		db:                  db,
		dirty:               true,
		filterer:            filterer,
		localValidatorInfos: localValidatorInfos,
		mode:                cfg.Mode,
	}, nil
}

func (ar *ValidatorRegistry) Start(_ context.Context) error {
	if err := ar.db.AutoMigrate(&Asset{}, &Tube{}, &AssetOnTube{}, &Meta{}); err != nil {
		return err
	}
	return ar.db.Clauses(clause.OnConflict{DoNothing: true}).Create(&Meta{
		Key:    "r_validator_registry_synced_height",
		Height: 0,
	}).Error
}

func (ar *ValidatorRegistry) Stop(_ context.Context) error {
	return nil
}

func (ar *ValidatorRegistry) syncedHeight() (uint64, error) {
	return readMeta(ar.db, "r_validator_registry_synced_height")
}

func (ar *ValidatorRegistry) TipHeight() (uint64, error) {
	ar.mutex.Lock()
	defer ar.mutex.RUnlock()

	return readMeta(ar.db, "r_validator_registry_tip_height")
}

func (ar *ValidatorRegistry) FetchEvents() error {
	ar.mutex.Lock()
	defer ar.mutex.Unlock()
	for addr, lvi := range ar.localValidatorInfos {
		if lvi.Genre == HTTP {
			resp, err := http.Get(lvi.URI + "/../health")
			if err != nil {
				log.Printf("local validator %s is not reachable, %+v\n", addr, err)
				continue
			}
			if resp.StatusCode != 200 {
				log.Printf("local validator %s has invalid status %d\n", addr, resp.StatusCode)
				continue
			}
		}
	}
	tip, err := ar.client.BlockNumber(context.Background())
	if err != nil {
		return errors.Wrap(err, "failed to get tip height for asset registry")
	}
	syncedHeight, err := ar.syncedHeight()
	if err != nil {
		return errors.Wrap(err, "failed to get synced height of asset registry")
	}
	startHeight := syncedHeight + 1
	if startHeight < ar.startHeight {
		startHeight = ar.startHeight
	}
	endHeight := startHeight + ar.batchSize - 1
	if tip < endHeight {
		endHeight = tip
	}
	if startHeight > endHeight {
		return nil
	}
	iter, err := ar.filterer.FilterRegistration(&bind.FilterOpts{
		Start: startHeight,
		End:   &endHeight,
	}, nil, nil)
	if err != nil {
		return errors.Wrap(err, "failed to fetch Registration logs")
	}
	events := []*contract.ValidatorRegistryRegistration{}
	for iter.Next() {
		events = append(events, iter.Event)
	}
	return ar.db.Transaction(func(tx *gorm.DB) error {
		if len(events) != 0 {
			ar.dirty = true
			for _, event := range events {
				if err := tx.Clauses(clause.OnConflict{DoNothing: true}).Create(&Validator{
					BlockHeight: event.Raw.BlockNumber,
					TxHash:      event.Raw.TxHash.Hex(),
					LogIndex:    event.Raw.Index,
					Address:     event.Raw.Address.Hex(),
					Genre:       event.Genre.Uint64(),
					URI:         event.Uri,
				}).Error; err != nil {
					return err
				}
			}
		}
		if err := tx.Clauses(clause.OnConflict{UpdateAll: true}).Create(&Meta{
			Key:    "r_validator_registry_synced_height",
			Height: syncedHeight,
		}).Error; err != nil {
			return err
		}
		return tx.Clauses(clause.OnConflict{UpdateAll: true}).Create(&Meta{
			Key:    "r_validator_registry_tip_height",
			Height: tip,
		}).Error
	})
}

func (ar *ValidatorRegistry) Validator(validator common.Address) (uint64, string, error) {
	ar.mutex.RLock()
	defer ar.mutex.RUnlock()
	switch ar.mode {
	case 3:
		return ar.validatorInContract(validator)
	case 2:
		return ar.validatorInLocal(validator)
	case 1:
		genre, uri, err := ar.validatorInContract(validator)
		if errors.Cause(err) != ErrUnknownValidator {
			return genre, uri, nil
		}
		return ar.validatorInLocal(validator)
	default:
		genre, uri, err := ar.validatorInLocal(validator)
		if errors.Cause(err) != ErrUnknownValidator {
			return genre, uri, nil
		}
		return ar.validatorInContract(validator)
	}
}

func (ar *ValidatorRegistry) validatorInLocal(validator common.Address) (uint64, string, error) {
	if info, ok := ar.localValidatorInfos[validator]; ok {
		return info.Genre, info.URI, nil
	}
	return 0, "", ErrUnknownValidator
}

func (ar *ValidatorRegistry) validatorInContract(validator common.Address) (uint64, string, error) {
	if ar.dirty {
		validatorInfos := map[common.Address]ValidatorInfo{}
		if err := ar.db.Transaction(func(tx *gorm.DB) error {
			records := []*Validator{}
			if err := tx.Model(&Validator{}).Order("block_height desc, log_index desc").Find(&records).Error; err != nil {
				return err
			}
			for _, record := range records {
				addr := common.HexToAddress(record.Address)
				if _, exist := validatorInfos[addr]; !exist {
					validatorInfos[addr] = ValidatorInfo{Genre: record.Genre, URI: record.URI}
				}
			}

			return nil
		}); err != nil {
			return 0, "", err
		}
		ar.validatorInfos = validatorInfos
		ar.dirty = false
	}
	info, ok := ar.validatorInfos[validator]
	if !ok {
		return 0, "", errors.Wrapf(ErrUnknownValidator, "validator %s has not registered", validator)
	}

	return info.Genre, info.URI, nil
}
