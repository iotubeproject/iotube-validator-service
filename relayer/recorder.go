package relayer

import (
	"context"
	"math/big"

	"github.com/crosschained-io/crosschained-service/types"

	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Recorder struct {
	url string
	db  *gorm.DB
}

func NewRecorder(databaseURL string) *Recorder {
	return &Recorder{url: databaseURL}
}

func (recorder *Recorder) Start(context.Context) error {
	db, err := gorm.Open(postgres.Open(recorder.url), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		return err
	}
	if err := db.AutoMigrate(&types.Deposit{}, &types.RelayTask{}); err != nil {
		return err
	}
	recorder.db = db

	return nil
}

func (recorder *Recorder) Stop(context.Context) error {
	db, err := recorder.db.DB()
	if err != nil {
		return err
	}

	return db.Close()
}

func (recorder *Recorder) UpdateRelayTask(id common.Hash, relayer common.Address, nonce uint64, gasPrice *big.Int, txHash common.Hash) error {
	deposit := &types.Deposit{}
	if result := recorder.db.Where("id = ?", id.Hex()).First(deposit); result.Error != nil {
		return result.Error
	}
	return recorder.db.Transaction(func(tx *gorm.DB) error {
		if result := tx.Updates(&types.RelayTask{
			ID:       id.Hex(),
			Relayer:  relayer.Hex(),
			Nonce:    nonce,
			GasPrice: decimal.NewFromBigInt(gasPrice, 0),
			TxHash:   txHash.Hex(),
		}); result.Error != nil {
			return result.Error
		}
		if result := tx.Model(&types.Deposit{}).Where("id = ?", id.Hex()).Update("status", types.Processed).Update("backed_up", false); result.Error != nil {
			return result.Error
		}
		return nil
	})
}

func (recorder *Recorder) LatestTask(tubeID uint64, relayer common.Address) (*types.RelayTask, error) {
	relayTask := types.RelayTask{}
	if result := recorder.db.
		Preload("Deposit").
		Preload("Testimonies").
		Table("relay_tasks").
		Order("relay_tasks.nonce desc").
		Joins("INNER JOIN deposits ON deposits.id = relay_tasks.id").
		Where("deposits.destination_tube_id = ? AND relay_tasks.relayer = ?", tubeID, relayer.Hex()).
		First(&relayTask); result.Error != nil {
		return nil, result.Error
	}
	return &relayTask, nil
}

func (recorder *Recorder) TasksToRelay(tubeID uint64, offset, limit int) ([]types.RelayTask, error) {
	tasks := []types.RelayTask{}
	if result := recorder.db.
		Limit(limit).
		Offset(offset).
		Preload("Testimonies").
		Preload("Deposit").
		Table("relay_tasks").
		Select("relay_tasks.*").
		Order("deposits.created_at").
		Joins("INNER JOIN deposits ON deposits.id = relay_tasks.id").
		Where("relay_tasks.status = ? AND deposits.destination_tube_id = ?", types.Ready, tubeID).
		Find(&tasks); result.Error != nil {
		return nil, result.Error
	}
	return tasks, nil
}
