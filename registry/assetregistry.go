package registry

import (
	"context"
	"math/big"
	"strings"
	"sync"

	"github.com/crosschained-io/crosschained-service/contract"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var ErrAssetNotFound = errors.New("asset not found")

type (
	tubeInfo struct {
		addrToID   map[common.Address]uint64
		idToAddr   map[uint64]common.Address
		idToStatus map[uint64]bool
	}

	AssetRegistry struct {
		abi            abi.ABI
		batchSize      uint64
		startHeight    uint64
		client         *ethclient.Client
		contractAddr   common.Address
		db             *gorm.DB
		dirty          bool
		filterer       *contract.AssetRegistryFilterer
		mode           int
		mutex          sync.RWMutex
		tubeInfos      map[uint64]*tubeInfo
		localTubeInfos map[uint64]*tubeInfo
	}
)

func newTubeInfo() *tubeInfo {
	return &tubeInfo{
		addrToID:   map[common.Address]uint64{},
		idToAddr:   map[uint64]common.Address{},
		idToStatus: map[uint64]bool{},
	}
}

func (ti *tubeInfo) AddAsset(id uint64, addr common.Address, deactivated bool) error {
	if _, exist := ti.addrToID[addr]; exist {
		return errors.New("duplicate address")
	}
	if _, exist := ti.idToAddr[id]; exist {
		return errors.New("duplicate id")
	}
	ti.addrToID[addr] = id
	ti.idToAddr[id] = addr
	ti.idToStatus[id] = !deactivated

	return nil
}

func (ti *tubeInfo) AssetID(addr common.Address) (uint64, bool) {
	id, ok := ti.addrToID[addr]
	if !ok {
		return 0, false
	}

	return id, true
}

func (ti *tubeInfo) AssetAddr(id uint64) (common.Address, error) {
	addr, ok := ti.idToAddr[id]
	if !ok {
		return common.Address{}, errors.Wrapf(ErrAssetNotFound, "%d", id)
	}

	return addr, nil
}

func (ti *tubeInfo) IsAssetActive(id uint64) bool {
	status, ok := ti.idToStatus[id]
	if !ok {
		return false
	}

	return status
}

func NewAssetRegistry(
	cfg AssetRegistryConfig,
	client *ethclient.Client,
	db *gorm.DB,
) (*AssetRegistry, error) {
	parsed, err := abi.JSON(strings.NewReader(contract.AssetRegistryABI))
	if err != nil {
		return nil, err
	}
	contractAddr := common.HexToAddress(cfg.ContractAddr)
	filterer, err := contract.NewAssetRegistryFilterer(contractAddr, client)
	if err != nil {
		return nil, err
	}
	localTubeInfos := map[uint64]*tubeInfo{}
	for _, asset := range cfg.Assets {
		ti, ok := localTubeInfos[asset.TubeID]
		if !ok {
			ti = newTubeInfo()
			localTubeInfos[asset.TubeID] = ti
		}
		if err := ti.AddAsset(asset.ID, common.HexToAddress(asset.Address), false); err != nil {
			return nil, err
		}
	}
	return &AssetRegistry{
		abi:            parsed,
		batchSize:      cfg.BatchSize,
		startHeight:    cfg.StartHeight,
		client:         client,
		contractAddr:   contractAddr,
		db:             db,
		dirty:          true,
		filterer:       filterer,
		localTubeInfos: localTubeInfos,
		mode:           cfg.Mode,
	}, nil
}

func (ar *AssetRegistry) Start(_ context.Context) error {
	return ar.db.AutoMigrate(&Asset{}, &Tube{}, &AssetOnTube{}, &Meta{})
}

func (ar *AssetRegistry) Stop(_ context.Context) error {
	return nil
}

func (ar *AssetRegistry) SyncedHeight() (uint64, error) {
	ar.mutex.Lock()
	defer ar.mutex.Unlock()
	return ar.syncedHeight()
}

func (ar *AssetRegistry) syncedHeight() (uint64, error) {
	return readMeta(ar.db, "r_asset_registry_synced_height")
}

func readMeta(db *gorm.DB, key string) (uint64, error) {
	var meta Meta
	switch err := db.Where("key = ?", key).First(&meta).Error; err {
	case gorm.ErrRecordNotFound:
		return 0, nil
	case nil:
		return meta.Height, nil
	default:
		return 0, err
	}
}

func (ar *AssetRegistry) TipHeight() (uint64, error) {
	ar.mutex.RLock()
	defer ar.mutex.RUnlock()

	return ar.tipHeight()
}

func (ar *AssetRegistry) tipHeight() (uint64, error) {
	return readMeta(ar.db, "r_asset_registry_tip_height")
}

func (ar *AssetRegistry) putAssetOnTubeRecord(
	tx *gorm.DB,
	blockHeight uint64,
	txHash common.Hash,
	logIndex uint,
	assetID *big.Int,
	tubeID *big.Int,
	assetAddr common.Address,
	active bool,
) error {
	return tx.Clauses(clause.OnConflict{DoNothing: true}).Create(&AssetOnTube{
		BlockHeight:  blockHeight,
		TxHash:       txHash.Hex(),
		LogIndex:     logIndex,
		AssetID:      assetID.Uint64(),
		TubeID:       tubeID.Uint64(),
		AssetAddress: assetAddr.Hex(),
		Active:       active,
	}).Error
}

func (ar *AssetRegistry) putTubeRecord(
	tx *gorm.DB,
	blockHeight uint64,
	txHash common.Hash,
	logIndex uint,
	tubeID *big.Int,
	active bool,
) error {
	return tx.Clauses(clause.OnConflict{DoNothing: true}).Create(&Tube{
		BlockHeight: blockHeight,
		TxHash:      txHash.Hex(),
		LogIndex:    logIndex,
		ID:          tubeID.Uint64(),
		Deactive:    !active,
	}).Error
}

func (ar *AssetRegistry) putAssetRecord(
	tx *gorm.DB,
	blockHeight uint64,
	txHash common.Hash,
	logIndex uint,
	assetID *big.Int,
	active bool,
) error {
	return tx.Clauses(clause.OnConflict{DoNothing: true}).Create(&Asset{
		BlockHeight: blockHeight,
		TxHash:      txHash.Hex(),
		LogIndex:    logIndex,
		ID:          assetID.Uint64(),
		Deactive:    !active,
	}).Error
}

func (ar *AssetRegistry) put(tipHeight, syncedHeight uint64, logs ...types.Log) error {
	return ar.db.Transaction(func(tx *gorm.DB) error {
		if len(logs) != 0 {
			ar.dirty = true
			for _, log := range logs {
				switch log.Topics[0] {
				case common.BytesToHash(ar.abi.Methods["AssetActivated"].ID):
					event, err := ar.filterer.ParseAssetActivated(log)
					if err != nil {
						return err
					}
					if err := ar.putAssetRecord(
						tx,
						event.Raw.BlockNumber,
						event.Raw.TxHash,
						event.Raw.Index,
						event.Id,
						true,
					); err != nil {
						return err
					}
				case common.BytesToHash(ar.abi.Methods["AssetDeactivated"].ID):
					event, err := ar.filterer.ParseAssetDeactivated(log)
					if err != nil {
						return err
					}
					if err := ar.putAssetRecord(
						tx,
						event.Raw.BlockNumber,
						event.Raw.TxHash,
						event.Raw.Index,
						event.Id,
						false,
					); err != nil {
						return err
					}
				case common.BytesToHash(ar.abi.Methods["NewAsset"].ID):
					event, err := ar.filterer.ParseNewAsset(log)
					if err != nil {
						return err
					}
					if err := ar.putAssetOnTubeRecord(
						tx,
						event.Raw.BlockNumber,
						event.Raw.TxHash,
						event.Raw.Index,
						event.AssetID,
						event.TubeID,
						event.AssetAddress,
						true,
					); err != nil {
						return err
					}
				case common.BytesToHash(ar.abi.Methods["AssetSetOnTube"].ID):
					event, err := ar.filterer.ParseAssetSetOnTube(log)
					if err != nil {
						return err
					}
					if err := ar.putAssetOnTubeRecord(
						tx,
						event.Raw.BlockNumber,
						event.Raw.TxHash,
						event.Raw.Index,
						event.AssetID,
						event.TubeID,
						event.AssetAddress,
						true,
					); err != nil {
						return err
					}
				case common.BytesToHash(ar.abi.Methods["AssetRemovedOnTube"].ID):
					event, err := ar.filterer.ParseAssetRemovedOnTube(log)
					if err != nil {
						return err
					}
					if err := ar.putAssetOnTubeRecord(
						tx,
						event.Raw.BlockNumber,
						event.Raw.TxHash,
						event.Raw.Index,
						event.AssetID,
						event.TubeID,
						event.AssetAddress,
						false,
					); err != nil {
						return err
					}
				case common.BytesToHash(ar.abi.Methods["AssetOnTubeActivated"].ID):
					event, err := ar.filterer.ParseAssetOnTubeActivated(log)
					if err != nil {
						return err
					}
					if err := ar.putAssetOnTubeRecord(
						tx,
						event.Raw.BlockNumber,
						event.Raw.TxHash,
						event.Raw.Index,
						event.AssetID,
						event.TubeID,
						event.AssetAddress,
						true,
					); err != nil {
						return err
					}
				case common.BytesToHash(ar.abi.Methods["AssetOnTubeDeactivated"].ID):
					event, err := ar.filterer.ParseAssetOnTubeDeactivated(log)
					if err != nil {
						return err
					}
					if err := ar.putAssetOnTubeRecord(
						tx,
						event.Raw.BlockNumber,
						event.Raw.TxHash,
						event.Raw.Index,
						event.AssetID,
						event.TubeID,
						event.AssetAddress,
						false,
					); err != nil {
						return err
					}
				case common.BytesToHash(ar.abi.Methods["TubeActivated"].ID):
					event, err := ar.filterer.ParseTubeActivated(log)
					if err != nil {
						return err
					}
					if err := ar.putTubeRecord(
						tx,
						event.Raw.BlockNumber,
						event.Raw.TxHash,
						event.Raw.Index,
						event.Id,
						true,
					); err != nil {
						return err
					}
				case common.BytesToHash(ar.abi.Methods["TubeDeactivated"].ID):
					event, err := ar.filterer.ParseTubeDeactivated(log)
					if err != nil {
						return err
					}
					if err := ar.putTubeRecord(
						tx,
						event.Raw.BlockNumber,
						event.Raw.TxHash,
						event.Raw.Index,
						event.Id,
						true,
					); err != nil {
						return err
					}
				}
			}
		}
		if err := tx.Clauses(clause.OnConflict{UpdateAll: true}).Create(&Meta{
			Key:    "r_asset_registry_synced_height",
			Height: syncedHeight,
		}).Error; err != nil {
			return err
		}
		return tx.Clauses(clause.OnConflict{UpdateAll: true}).Create(&Meta{
			Key:    "r_asset_registry_tip_height",
			Height: tipHeight,
		}).Error
	})
}

func (ar *AssetRegistry) FetchEvents() error {
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
	config := ethereum.FilterQuery{
		Addresses: []common.Address{ar.contractAddr},
		Topics:    nil,
		FromBlock: new(big.Int).SetUint64(startHeight),
		ToBlock:   new(big.Int).SetUint64(endHeight),
	}
	logs, err := ar.client.FilterLogs(context.Background(), config)
	if err != nil {
		return errors.Wrap(err, "failed to fetch AssetActivated event")
	}

	return ar.put(tip, endHeight, logs...)
}

func (ar *AssetRegistry) AssetOnTube(
	sourceTubeID uint64,
	sourceAssetAddr common.Address,
	targetTubeID uint64,
) (uint64, common.Address, bool, error) {
	ar.mutex.RLock()
	defer ar.mutex.RUnlock()
	switch ar.mode {
	case 3:
		return ar.assetOnTubeInContract(
			sourceTubeID,
			sourceAssetAddr,
			targetTubeID,
		)
	case 2:
		return ar.assetOnTubeInLocal(
			sourceTubeID,
			sourceAssetAddr,
			targetTubeID,
		)
	case 1:
		assetID, assetAddress, assetStatus, err := ar.assetOnTubeInContract(
			sourceTubeID,
			sourceAssetAddr,
			targetTubeID,
		)
		if errors.Cause(err) != ErrAssetNotFound {
			return assetID, assetAddress, assetStatus, err
		}
		return ar.assetOnTubeInLocal(
			sourceTubeID,
			sourceAssetAddr,
			targetTubeID,
		)
	default:
		assetID, assetAddress, assetStatus, err := ar.assetOnTubeInLocal(
			sourceTubeID,
			sourceAssetAddr,
			targetTubeID,
		)
		if errors.Cause(err) != ErrAssetNotFound {
			return assetID, assetAddress, assetStatus, err
		}
		return ar.assetOnTubeInContract(
			sourceTubeID,
			sourceAssetAddr,
			targetTubeID,
		)
	}
}

func (ar *AssetRegistry) assetOnTubeInLocal(
	sourceTubeID uint64,
	sourceAssetAddr common.Address,
	targetTubeID uint64,
) (uint64, common.Address, bool, error) {
	ti, ok := ar.localTubeInfos[sourceTubeID]
	if !ok {
		return 0, common.Address{}, false, errors.Wrapf(ErrAssetNotFound, "no asset on tube %d", sourceTubeID)
	}
	assetID, ok := ti.AssetID(sourceAssetAddr)
	if !ok {
		return 0, common.Address{}, false, errors.Errorf("failed to fetch asset id for %s", sourceAssetAddr)
	}
	if assetID == 0 {
		return 0, common.Address{}, false, nil
	}
	ti, ok = ar.localTubeInfos[targetTubeID]
	if !ok {
		return 0, common.Address{}, false, errors.Wrapf(ErrAssetNotFound, "no asset on tube %d", targetTubeID)
	}
	addr, err := ti.AssetAddr(assetID)
	if err != nil {
		return assetID, addr, false, err
	}

	return assetID, addr, ti.IsAssetActive(assetID), nil
}

func (ar *AssetRegistry) assetOnTubeInContract(
	sourceTubeID uint64,
	sourceAssetAddr common.Address,
	targetTubeID uint64,
) (uint64, common.Address, bool, error) {
	if ar.dirty {
		tubeInfos := map[uint64]*tubeInfo{}
		if err := ar.db.Transaction(func(tx *gorm.DB) error {
			tubes := []*Tube{}
			if err := tx.Model(&Tube{}).Order("block_height desc, log_index desc").Find(&tubes).Error; err != nil {
				return err
			}
			tubeStatus := map[uint64]bool{}
			for _, record := range tubes {
				if _, exist := tubeStatus[record.ID]; !exist {
					tubeStatus[record.ID] = record.Deactive
				}
			}
			assets := []*Asset{}
			if err := tx.Model(&Asset{}).Order("block_height desc, log_index desc").Find(&assets).Error; err != nil {
				return err
			}
			assetStatus := map[uint64]bool{}
			for _, record := range assets {
				if _, exist := assetStatus[record.ID]; !exist {
					assetStatus[record.ID] = record.Deactive
				}
			}
			aont := []*AssetOnTube{}
			if err := tx.Model(&AssetOnTube{}).Order("block_height desc, log_index desc").Find(&aont).Error; err != nil {
				return err
			}
			assetMap := map[uint64]map[uint64]struct{}{}
			for _, record := range aont {
				if _, exist := assetMap[record.AssetID]; !exist {
					assetMap[record.AssetID] = map[uint64]struct{}{}
				}
				if _, exist := assetMap[record.AssetID][record.TubeID]; !exist {
					assetMap[record.AssetID][record.TubeID] = struct{}{}
					if assetStatus[record.AssetID] || tubeStatus[record.TubeID] {
						record.Active = false
					}
					if _, exist := tubeInfos[record.TubeID]; !exist {
						tubeInfos[record.TubeID] = newTubeInfo()
					}
					if err := tubeInfos[record.TubeID].AddAsset(
						record.AssetID,
						common.HexToAddress(record.AssetAddress),
						record.Active,
					); err != nil {
						return err
					}
				}
			}

			return nil
		}); err != nil {
			return 0, common.Address{}, false, err
		}
		ar.tubeInfos = tubeInfos
		ar.dirty = false
	}
	syncedHeight, err := ar.syncedHeight()
	if err != nil {
		return 0, common.Address{}, false, errors.Wrap(err, "failed to read synced height")
	}
	tip, err := ar.tipHeight()
	if err != nil {
		return 0, common.Address{}, false, errors.Wrap(err, "failed to read tip height")
	}
	if syncedHeight != tip {
		return 0, common.Address{}, false, errors.Errorf("asset registry is not up to date")
	}
	tr, ok := ar.tubeInfos[sourceTubeID]
	if !ok {
		return 0, common.Address{}, false, errors.Wrapf(ErrAssetNotFound, "no asset on source tube %d", sourceTubeID)
	}
	assetID, ok := tr.AssetID(sourceAssetAddr)
	if !ok {
		return 0, common.Address{}, false, errors.Errorf("failed to fetch asset id for %s", sourceAssetAddr)
	}
	tr, ok = ar.tubeInfos[targetTubeID]
	if !ok {
		return 0, common.Address{}, false, errors.Wrapf(ErrAssetNotFound, "no asset on target tube %d", targetTubeID)
	}
	addr, err := tr.AssetAddr(assetID)
	if err != nil {
		return assetID, addr, false, err
	}

	return assetID, addr, tr.IsAssetActive(assetID), nil
}
