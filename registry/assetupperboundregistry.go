package registry

import (
	"context"
	"math/big"
	"sync"

	"github.com/crosschained-io/crosschained-service/contract"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type (
	AssetUpperBoundRegistry struct {
		batchSize        uint64
		startHeight      uint64
		client           *ethclient.Client
		contractAddr     common.Address
		db               *gorm.DB
		dirty            bool
		filterer         *contract.AssetUpperBoundFilterer
		mutex            sync.RWMutex
		upperBounds      map[uint64]*big.Int
		localUpperBounds map[uint64]*big.Int
		mode             int
	}
)

func NewAssetUpperBoundRegistry(
	cfg AssetUpperBoundRegistryConfig,
	client *ethclient.Client,
	db *gorm.DB,
) (*AssetUpperBoundRegistry, error) {
	contractAddr := common.HexToAddress(cfg.ContractAddr)
	filterer, err := contract.NewAssetUpperBoundFilterer(contractAddr, client)
	if err != nil {
		return nil, err
	}
	localUpperBounds := map[uint64]*big.Int{}
	for assetID, value := range cfg.UpperBounds {
		ub, ok := new(big.Int).SetString(value, 10)
		if !ok {
			return nil, errors.Errorf("invalid upper found value %s", value)
		}
		localUpperBounds[assetID] = ub
	}
	return &AssetUpperBoundRegistry{
		batchSize:        cfg.BatchSize,
		startHeight:      cfg.StartHeight,
		client:           client,
		contractAddr:     contractAddr,
		db:               db,
		dirty:            true,
		filterer:         filterer,
		localUpperBounds: localUpperBounds,
		mode:             cfg.Mode,
	}, nil
}

func (ar *AssetUpperBoundRegistry) Start(_ context.Context) error {
	if err := ar.db.AutoMigrate(&Asset{}, &Tube{}, &AssetOnTube{}, &Meta{}); err != nil {
		return err
	}
	return ar.db.Clauses(clause.OnConflict{DoNothing: true}).Create(&Meta{
		Key:    "r_asset_upper_bound_registry_synced_height",
		Height: 0,
	}).Error
}

func (ar *AssetUpperBoundRegistry) Stop(_ context.Context) error {
	return nil
}

func (ar *AssetUpperBoundRegistry) SyncedHeight() (uint64, error) {
	ar.mutex.Lock()
	defer ar.mutex.RUnlock()

	return ar.syncedHeight()
}

func (ar *AssetUpperBoundRegistry) syncedHeight() (uint64, error) {
	return readMeta(ar.db, "r_asset_upper_bound_registry_synced_height")
}

func (ar *AssetUpperBoundRegistry) TipHeight() (uint64, error) {
	ar.mutex.Lock()
	defer ar.mutex.RUnlock()

	return readMeta(ar.db, "r_asset_upper_bound_registry_tip_height")
}

func (ar *AssetUpperBoundRegistry) FetchEvents() error {
	ar.mutex.Lock()
	defer ar.mutex.Unlock()
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
	iter, err := ar.filterer.FilterUpperBoundSet(&bind.FilterOpts{
		Start: startHeight,
		End:   &endHeight,
	}, nil)
	if err != nil {
		return errors.Wrap(err, "failed to fetch Registration logs")
	}
	events := []*contract.AssetUpperBoundUpperBoundSet{}
	for iter.Next() {
		events = append(events, iter.Event)
	}
	return ar.db.Transaction(func(tx *gorm.DB) error {
		if len(events) != 0 {
			ar.dirty = true
			for _, event := range events {
				if err := tx.Clauses(clause.OnConflict{DoNothing: true}).Create(&AssetUpperBound{
					BlockHeight: event.Raw.BlockNumber,
					TxHash:      event.Raw.TxHash.Hex(),
					LogIndex:    event.Raw.Index,
					AssetID:     event.AssetID.Uint64(),
					UpperBound:  event.UpperBound.String(),
				}).Error; err != nil {
					return err
				}
			}
		}
		if err := tx.Clauses(clause.OnConflict{UpdateAll: true}).Create(&Meta{
			Key:    "r_asset_upper_bound_registry_synced_height",
			Height: syncedHeight,
		}).Error; err != nil {
			return err
		}
		return tx.Clauses(clause.OnConflict{UpdateAll: true}).Create(&Meta{
			Key:    "r_asset_upper_bound_registry_tip_height",
			Height: tip,
		}).Error
	})
}

func (ar *AssetUpperBoundRegistry) UpperBound(assetID uint64) (*big.Int, error) {
	ar.mutex.RLock()
	defer ar.mutex.RUnlock()
	switch ar.mode {
	case 3:
		return ar.upperBoundInConract(assetID)
	case 2:
		return ar.upperBoundInLocal(assetID)
	case 1:
		ub, err := ar.upperBoundInConract(assetID)
		if err != nil {
			return nil, err
		}
		if ub != nil {
			return ub, nil
		}
		return ar.upperBoundInLocal(assetID)
	default:
		ub, err := ar.upperBoundInLocal(assetID)
		if err != nil {
			return nil, err
		}
		if ub != nil {
			return ub, nil
		}
		return ar.upperBoundInConract(assetID)
	}
}

func (ar *AssetUpperBoundRegistry) upperBoundInLocal(assetID uint64) (*big.Int, error) {
	if upperBound, ok := ar.localUpperBounds[assetID]; ok {
		return upperBound, nil
	}

	return nil, nil
}

func (ar *AssetUpperBoundRegistry) upperBoundInConract(assetID uint64) (*big.Int, error) {
	if ar.dirty {
		upperBounds := map[uint64]*big.Int{}
		if err := ar.db.Transaction(func(tx *gorm.DB) error {
			records := []*AssetUpperBound{}
			if err := tx.Model(&AssetUpperBound{}).Order("block_height desc, log_index desc").Find(&records).Error; err != nil {
				return err
			}
			for _, record := range records {
				if _, exist := upperBounds[record.AssetID]; !exist {
					upperBound, ok := new(big.Int).SetString(record.UpperBound, 10)
					if !ok {
						return errors.Errorf("failed to parse %s to big.Int", record.UpperBound)
					}
					upperBounds[record.AssetID] = upperBound
				}
			}

			return nil
		}); err != nil {
			return nil, err
		}
		ar.upperBounds = upperBounds
		ar.dirty = false
	}
	upperBound, ok := ar.upperBounds[assetID]
	if !ok {
		return nil, nil
	}

	return upperBound, nil
}
