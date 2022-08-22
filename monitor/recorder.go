package monitor

import (
	"encoding/hex"
	"math/big"
	"sync"

	"github.com/crosschained-io/crosschained-service/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Recorder struct {
	db         *gorm.DB
	explorerDB *ExplorerDB
	tx         *gorm.DB
	mutex      sync.RWMutex
}

func NewRecorderForRelayer(db *gorm.DB, explorerDB *ExplorerDB) (*Recorder, error) {
	if err := db.AutoMigrate(
		&types.Deposit{},
		&types.Withdrawal{},
		&types.RelayTask{},
		&types.Testimony{},
		&types.Validator{},
		&types.Tube{},
		&types.VerifierStatus{},
	); err != nil {
		return nil, err
	}
	if explorerDB != nil {
		if err := explorerDB.Init(); err != nil {
			return nil, err
		}
	}
	return &Recorder{db: db, explorerDB: explorerDB}, nil
}

func NewRecorderForValidator(db *gorm.DB) (*Recorder, error) {
	if err := db.AutoMigrate(
		&types.Deposit{},
		&types.SignatureTask{},
		&types.Validator{},
		&types.Tube{},
	); err != nil {
		return nil, err
	}
	return &Recorder{db: db}, nil
}

func (recorder *Recorder) Begin() error {
	recorder.mutex.Lock()
	defer recorder.mutex.Unlock()
	if recorder.tx != nil {
		return types.ErrTwoTransactions
	}
	recorder.tx = recorder.db.Begin()

	return recorder.tx.Error
}

func (recorder *Recorder) Create(value interface{}) error {
	recorder.mutex.Lock()
	defer recorder.mutex.Unlock()
	if recorder.tx == nil {
		return recorder.db.Create(value).Error
	}

	return recorder.tx.Clauses(clause.OnConflict{DoNothing: true}).Create(value).Error
}

func (recorder *Recorder) Upsert(value interface{}) error {
	recorder.mutex.Lock()
	defer recorder.mutex.Unlock()
	if recorder.tx == nil {
		return types.ErrNoTransaction
	}

	return recorder.tx.Clauses(clause.OnConflict{UpdateAll: true}).Create(value).Error
}

func (recorder *Recorder) Commit() error {
	recorder.mutex.Lock()
	defer recorder.mutex.Unlock()
	if recorder.tx == nil {
		return types.ErrNoTransaction
	}
	if err := recorder.tx.Commit().Error; err != nil {
		return err
	}
	recorder.tx = nil

	return nil
}

func (recorder *Recorder) Rollback() {
	recorder.mutex.Lock()
	defer recorder.mutex.Unlock()
	if recorder.tx != nil {
		recorder.tx.Rollback()
		recorder.tx = nil
	}
}

func (recorder *Recorder) TubeSyncedHeight(id uint64) (uint64, error) {
	var t types.Tube
	switch err := recorder.db.First(&t, id).Error; err {
	case gorm.ErrRecordNotFound:
		return 0, nil
	case nil:
		return t.SyncedHeight, nil
	default:
		return 0, err
	}
}

func (recorder *Recorder) IsValidator(tubeID uint64, validator common.Address) (bool, error) {
	var v types.Validator
	switch err := recorder.db.
		Table("validators").
		Where("tube_id = ? AND address = ?", tubeID, validator.Hex()).
		Order("block_height desc, log_index desc").First(&v).Error; errors.Cause(err) {
	case gorm.ErrRecordNotFound:
		return false, nil
	case nil:
		return v.Active, nil
	default:
		return false, err
	}
}

func (recorder *Recorder) Validators(tubeID uint64) ([]common.Address, error) {
	validatorAddrs := []common.Address{}
	if err := recorder.db.Transaction(func(tx *gorm.DB) error {
		var t types.Tube
		if err := tx.First(&t, tubeID).Error; err != nil {
			return err
		}
		if !t.VerifierUpToDate {
			return errors.Wrapf(types.ErrValidatorIsOutOfDate, "verifier is not up to tip %d", t.TipHeight)
		}
		validators := []types.Validator{}
		if err := tx.Table("validators").Where("tube_id = ?", tubeID).Order("block_height, log_index").Find(&validators).Error; err != nil {
			return err
		}
		statuses := map[string]bool{}
		for _, validator := range validators {
			statuses[validator.Address] = validator.Active
		}
		for addr, active := range statuses {
			if !active {
				continue
			}
			validatorAddrs = append(validatorAddrs, common.HexToAddress(addr))
		}
		return nil
	}); err != nil {
		return nil, err
	}

	return validatorAddrs, nil
}

func (recorder *Recorder) SettleDeposits(tubeID uint64) error {
	return recorder.db.Transaction(func(tx *gorm.DB) error {
		withdrawals := []*types.Withdrawal{}
		if result := tx.Table("withdrawals").
			Select("withdrawals.*").
			Joins("INNER JOIN relay_tasks ON relay_tasks.ID = withdrawals.ID").
			Where("relay_tasks.status <> ? AND withdrawals.tube_id = ?", types.Settled, tubeID).
			Find(&withdrawals); result.Error != nil {
			return result.Error
		}
		if len(withdrawals) == 0 {
			return nil
		}
		ids := make([]string, len(withdrawals))
		for i, withdrawal := range withdrawals {
			ids[i] = withdrawal.ID
		}

		return tx.Table("relay_tasks").Where("id IN ?", ids).Update("status", types.Settled).Update("backed_up", false).Error
	})
}

func (recorder *Recorder) UpdateSignatureTasks(tubeID uint64, validator common.Address, status int) error {
	return recorder.db.Transaction(func(tx *gorm.DB) error {
		tasks := []types.SignatureTask{}
		if err := tx.Table("signature_tasks").
			Preload("Deposit").
			Select("signature_tasks.id").
			Joins("INNER JOIN deposits ON deposits.id = signature_tasks.id").
			Where("signature_tasks.status = ? AND deposits.destination_tube_id = ?", types.New, tubeID).
			Find(&tasks).
			Error; err != nil {
			return err
		}
		ids := make([]string, len(tasks))
		for i, task := range tasks {
			ids[i] = task.ID
		}
		return tx.Table("signature_tasks").
			Where("id IN ?", ids).
			Update("status", status).
			Update("validator", validator.Hex()).
			Error
	})
}

func (recorder *Recorder) UpdateRelayTask(id common.Hash, relayer common.Address, nonce uint64, gasPrice *big.Int, txHash common.Hash) error {
	deposit := &types.Deposit{}
	if result := recorder.db.Where("id = ?", id.Hex()).First(deposit); result.Error != nil {
		return result.Error
	}
	return recorder.db.Transaction(func(tx *gorm.DB) error {
		if result := tx.Select("relayer, nonce, gas_price, tx_hash, status, backed_up").Updates(&types.RelayTask{
			ID:       id.Hex(),
			Relayer:  relayer.Hex(),
			Nonce:    nonce,
			GasPrice: decimal.NewFromBigInt(gasPrice, 0),
			TxHash:   txHash.Hex(),
			Status:   types.Processed,
			BackedUp: false,
		}); result.Error != nil {
			return result.Error
		}
		if err := tx.Model(&types.Deposit{}).Where("id = ?", id.Hex()).Update("status", types.Processed).Update("backed_up", false).Error; err != nil {
			return err
		}
		return nil
	})
}

func (recorder *Recorder) WaitingForTestimonies(tubeID uint64, offset int, limit int) ([]*types.RelayTask, error) {
	tasks := []*types.RelayTask{}
	if err := recorder.db.
		Preload("Deposit").
		Preload("Testimonies").
		Limit(int(limit)).
		Offset(int(offset)).
		Select("relay_tasks.*").
		Table("relay_tasks").
		Joins("INNER JOIN deposits ON deposits.id = relay_tasks.id").
		Where("relay_tasks.status = ? AND deposits.destination_tube_id = ?", types.New, tubeID).
		Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}

func (recorder *Recorder) MarkAsReady(id common.Hash) error {
	return recorder.db.Model(&types.RelayTask{}).Where("id = ? AND status = ?", id.String(), types.New).Update("status", types.Ready).Update("backed_up", false).Error
}

func (recorder *Recorder) MarkAsRejected(id common.Hash) error {
	return recorder.db.Model(&types.RelayTask{}).Where("id = ? AND status = ?", id.String(), types.New).Update("status", types.Rejected).Update("backed_up", false).Error
}

func (recorder *Recorder) AddTestimony(id common.Hash, validator common.Address, signature []byte) error {
	return recorder.db.Clauses(clause.OnConflict{
		DoNothing: true,
	}).Create(&types.Testimony{
		DepositID: id.String(),
		Validator: validator.String(),
		Signature: hex.EncodeToString(signature),
	}).Error
}

func (recorder *Recorder) BackupInExplorerDB() error {
	if recorder.explorerDB == nil {
		return nil
	}
	var tasks []types.RelayTask
	if err := recorder.db.Preload("Deposit").Table("relay_tasks").Where("backed_up = ?", false).Find(&tasks).Error; err != nil {
		return err
	}
	ids := make([]string, len(tasks))
	txHashes := make([]string, len(tasks))
	statuses := make([]string, len(tasks))
	for i, task := range tasks {
		txHash := ""
		status := "new"
		switch task.Status {
		case types.Ready:
			status = "processing"
		case types.Processed:
			status = "validated"
		case types.Settled:
			var withdrawal *types.Withdrawal
			if err := recorder.db.Table("withdrawals").Where("id = ?", task.ID).Find(&withdrawal).Error; err != nil {
				return err
			}
			txHash = withdrawal.TxHash
			status = "settled"
		case types.Rejected:
			status = "rejected"
		case types.Inactive:
			status = "failed"
		}
		txHashes[i] = txHash
		statuses[i] = status
		ids[i] = task.ID
	}
	if err := recorder.explorerDB.Backup(tasks, statuses, txHashes); err != nil {
		return err
	}

	return recorder.db.Transaction(func(tx *gorm.DB) error {
		for _, task := range tasks {
			if err := tx.Table("relay_tasks").Where("id = ? AND status = ?", task.ID, task.Status).Update("backed_up", true).Error; err != nil {
				return err
			}
		}
		return nil
	})
}
