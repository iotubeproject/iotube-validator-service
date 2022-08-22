package validator

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"log"
	"math/big"
	"time"

	"github.com/crosschained-io/crosschained-service/types"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type (
	item struct {
		task *types.SignatureTask
		ts   time.Time
	}
	TubeStats struct {
		processedItems []*item
		processingList map[common.Hash]*item
		waitList       map[common.Hash]*item
		quota          *big.Int
		duration       time.Duration
		amount         *big.Int
		id             uint64
	}
	Service struct {
		privateKey *ecdsa.PrivateKey
		stats      []*TubeStats
		db         *gorm.DB
		url        string
		terminate  chan struct{}
		interval   time.Duration
	}
)

func (ts *TubeStats) refresh() {
	now := time.Now()
	for i := len(ts.processedItems) - 1; i >= 0; i-- {
		if ts.processedItems[i].ts.Before(now) {
			for j := 0; j <= i; j++ {
				ts.quota = ts.quota.Add(ts.quota, ts.processedItems[j].task.Deposit.Amount.BigInt())
			}
			ts.processedItems = ts.processedItems[i+1:]
			return
		}
	}
}

func (ts *TubeStats) delay(amount *big.Int) time.Duration {
	return ts.duration * time.Duration(new(big.Int).Div(amount, ts.amount).Int64())
}

func (ts *TubeStats) hasSufficientQuota(amount *big.Int) bool {
	return ts.quota.Cmp(amount) > 0
}

func (ts *TubeStats) addToProcessingList(id common.Hash, it *item) {
	ts.quota = ts.quota.Sub(ts.quota, it.task.Deposit.Amount.BigInt())
	ts.processingList[id] = it
}

func (ts *TubeStats) moveToProcessedList(now time.Time) {
	for id, task := range ts.processingList {
		delete(ts.waitList, id)
		task.ts = now
		ts.processedItems = append(ts.processedItems, task)
	}
}

func NewService(
	privateKey *ecdsa.PrivateKey,
	tubeIDs []uint64,
	quotas []*big.Int,
	durations []time.Duration,
	amounts []*big.Int,
	databaseURL string,
	interval time.Duration,
) *Service {
	stats := make([]*TubeStats, 0, len(tubeIDs))
	for i := range tubeIDs {
		stats = append(stats, &TubeStats{
			processedItems: []*item{},
			processingList: map[common.Hash]*item{},
			waitList:       map[common.Hash]*item{},
			id:             tubeIDs[i],
			quota:          quotas[i],
			duration:       durations[i],
			amount:         amounts[i],
		})
	}
	return &Service{
		privateKey: privateKey,
		stats:      stats,
		url:        databaseURL,
		interval:   interval,
		terminate:  make(chan struct{}),
	}
}

func (s *Service) Start(ctx context.Context) error {
	db, err := gorm.Open(postgres.Open(s.url), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		return errors.Wrap(err, "failed to create db")
	}
	if err := db.AutoMigrate(&types.Deposit{}, &types.SignatureTask{}); err != nil {
		return err
	}
	s.db = db

	go func() {
		for {
			select {
			case <-s.terminate:
				return
			default:
				if err := s.run(); err != nil {
					log.Printf("failed to run sign service", err)
				}
			}
			time.Sleep(s.interval)
		}
	}()
	return nil
}

func (s *Service) Stop(ctx context.Context) error {
	close(s.terminate)
	time.Sleep(s.interval * 2)
	db, err := s.db.DB()
	if err != nil {
		return err
	}

	return db.Close()
}

func (s *Service) run() error {
	now := time.Now()
	for _, stat := range s.stats {
		stat.refresh()
		signatures := map[common.Hash][]byte{}
		for id, item := range stat.waitList {
			if item.ts.After(now) || !stat.hasSufficientQuota(item.task.Deposit.Amount.BigInt()) {
				continue
			}
			signature, err := crypto.Sign(id[:], s.privateKey)
			if err != nil {
				return err
			}
			signatures[id] = signature
			stat.addToProcessingList(id, item)
		}
		tasks, err := s.tasksToSign(stat.id)
		if err != nil {
			return err
		}
		for i := range tasks {
			task := tasks[i]
			id := common.HexToHash(task.ID)
			if _, ok := stat.waitList[id]; ok {
				continue
			}
			if delay := stat.delay(task.Deposit.Amount.BigInt()); delay > 0 {
				stat.waitList[id] = &item{
					task: &task,
					ts:   now.Add(delay),
				}
				continue
			}
			if !stat.hasSufficientQuota(task.Deposit.Amount.BigInt()) {
				stat.waitList[id] = &item{
					task: &task,
					ts:   now,
				}
				continue
			}
			signature, err := crypto.Sign(id[:], s.privateKey)
			if err != nil {
				return err
			}
			signatures[id] = signature
			stat.addToProcessingList(id, &item{
				task: &task,
				ts:   now,
			})
		}
		log.Printf("Signing %d tasks for tube %d", len(signatures), stat.id)
		if err := s.putSignatures(signatures); err != nil {
			return err
		}
		stat.moveToProcessedList(now)
	}
	return nil
}

func (s *Service) putSignatures(signatures map[common.Hash][]byte) error {
	return s.db.Transaction(func(tx *gorm.DB) error {
		for id, signature := range signatures {
			if err := tx.Model(&types.SignatureTask{}).Where("id = ? AND value = ?", id.Hex(), "").Updates(&types.SignatureTask{
				Value:  hex.EncodeToString(signature),
				Status: types.Processed,
			}).Error; err != nil {
				return err
			}
		}
		return nil
	})
}

func (s *Service) tasksToSign(tubeID uint64) ([]types.SignatureTask, error) {
	var tasks []types.SignatureTask
	if err := s.db.
		Table("signature_tasks").
		Preload("Deposit").
		Select("signature_tasks.*").
		Limit(10).
		Order("deposits.created_at").
		Joins("INNER JOIN deposits ON deposits.id = signature_tasks.id").
		Where("signature_tasks.status = ? AND (signature_tasks.value = ? OR signature_tasks.value IS NULL) AND deposits.destination_tube_id = ?", types.Ready, "", tubeID).
		Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}
