package api

import (
	"encoding/hex"

	"github.com/crosschained-io/crosschained-service/types"

	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type (
	Recorder struct {
		db *gorm.DB
	}
)

func NewRecorder(databaseURL string) (*Recorder, error) {
	db, err := gorm.Open(postgres.Open(databaseURL), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		return nil, errors.Wrap(err, "failed to create db")
	}
	if err := db.AutoMigrate(&types.SignatureTask{}); err != nil {
		return nil, err
	}
	return &Recorder{db: db}, nil
}

func (recorder *Recorder) Signature(id common.Hash) ([]byte, error) {
	var record types.SignatureTask
	if err := recorder.db.Where("id = ?", id.String()).First(&record).Error; err != nil {
		return nil, err
	}
	if record.Value == "" {
		return nil, errors.New("not signed yet")
	}

	return hex.DecodeString(record.Value)
}
