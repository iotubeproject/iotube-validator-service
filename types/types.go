package types

import (
	"time"

	"github.com/crosschained-io/crosschained-service/contract"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/lib/pq"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
)

var (
	ErrNotExist             = errors.New("record does not exist")
	ErrNoTransaction        = errors.New("no transaction")
	ErrTwoTransactions      = errors.New("one transaction at a time")
	ErrDuplicateRecord      = errors.New("record already exists")
	ErrValidatorIsOutOfDate = errors.New("validator is out of date")
)

const (
	Unknown = iota
	New
	Ready
	Processed
	Settled
	Rejected
	Inactive
	InReview
	OutOfTheDuty
)

type (
	Validator struct {
		ID          uint   `gorm:"primaryKey;autoIncrement"`
		TubeID      uint64 `gorm:"UNIQUE_INDEX:u_log;unsigned;index"`
		Address     string `gorm:"size:42;not null;index"`
		BlockHeight uint64 `gorm:"UNIQUE_INDEX:u_log;unsigned;index"`
		TxHash      string `gorm:"size:66;not null;index"`
		LogIndex    uint   `gorm:"UNIQUE_INDEX:u_log;unsigned;index"`
		Active      bool   `gorm:"type:bool;default:false"`
	}

	VerifierStatus struct {
		ID          uint   `gorm:"primaryKey;autoIncrement"`
		TubeID      uint64 `gorm:"UNIQUE_INDEX:u_log;unsigned;index"`
		TxHash      string `gorm:"size:66;not null;index"`
		BlockHeight uint64 `gorm:"UNIQUE_INDEX:u_log;unsigned;index"`
		LogIndex    uint   `gorm:"UNIQUE_INDEX:u_log;unsigned;index"`
		Paused      bool   `gorm:"type:bool;default:false"`
	}

	Deposit struct {
		ID                string          `gorm:"primary_key;size:66;not null;"`
		BlockHeight       uint64          `gorm:"unsigned;index"`
		TxHash            string          `gorm:"size:66;not null;uniqueIndex"`
		SourceTubeID      uint64          `gorm:"unsigned;index" sql:"type:bigint"`
		SourceTubeNonce   uint64          `gorm:"unsigned;index" sql:"type:bigint"`
		SourceToken       string          `gorm:"size:42;not null;index"`
		Sender            string          `gorm:"size:42;not null;index"`
		Recipient         string          `gorm:"size:42;not null;index"`
		DestinationTubeID uint64          `gorm:"unsigned;index" sql:"type:bigint"`
		DestinationToken  string          `gorm:"size:42;not null;index"`
		Amount            decimal.Decimal `gorm:"type:decimal(60,0);not null;"`
		Fee               decimal.Decimal `gorm:"type:decimal(60,0);not null;"`
		CreatedAt         time.Time
		UpdatedAt         time.Time
	}

	Testimony struct {
		ID        uint64 `gorm:"primary_key;unsigned"`
		DepositID string `gorm:"UNIQUE_INDEX:deposit_validator;size:66;not null;index"`
		Validator string `gorm:"UNIQUE_INDEX:deposit_validator;size:42;not null;index"`
		Signature string `gorm:"size:132;not null"`
		CreatedAt time.Time
		UpdatedAt time.Time
	}

	Withdrawal struct {
		ID          string         `gorm:"primary_key;size:66;not null;"`
		BlockHeight uint64         `gorm:"unsigned;index"`
		TubeID      uint64         `gorm:"unsigned"`
		TxHash      string         `gorm:"size:66;not null;uniqueIndex"`
		Validators  pq.StringArray `gorm:"type:text[]"`
		CreatedAt   time.Time
		UpdatedAt   time.Time
	}

	RelayTask struct {
		ID          string          `gorm:"primary_key;size:66;not null;"`
		Deposit     Deposit         `gorm:"foreignKey:ID;"`
		BackedUp    bool            `gorm:"type:bool;default:false"`
		Status      int             `gorm:"default:0;index"`
		Fee         decimal.Decimal `gorm:"type:decimal(60,0);not null;"`
		Relayer     string          `gorm:"size:42;index;default:'';"`
		Nonce       uint64          `gorm:"unsigned;index;" sql:"type:bigint"`
		TxHash      string          `gorm:"size:66;index;default:'';"`
		GasPrice    decimal.Decimal `gorm:"type:decimal(60,0);default:0"`
		Testimonies []Testimony     `gorm:"foreignKey:DepositID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
		CreatedAt   time.Time
		UpdatedAt   time.Time
	}

	Tube struct {
		ID               uint64 `gorm:"primary_key;unsigned"`
		TipHeight        uint64 `gorm:"unsigned"`
		SyncedHeight     uint64 `gorm:"unsigned"`
		VerifierUpToDate bool   `gorm:"type:bool"`
	}

	SignatureTask struct {
		ID        string  `gorm:"primary_key;size:66;not null;"`
		Deposit   Deposit `gorm:"foreignKey:ID;"`
		Validator string  `gorm:"size:42;index;"`
		Status    int     `gorm:"default:0;index;"`
		Value     string  `gorm:"size:132;not null"`
		CreatedAt time.Time
		UpdatedAt time.Time
	}
)

func (Deposit) TableName() string {
	return "deposits"
}

func (Testimony) TableName() string {
	return "testimonies"
}

func (Withdrawal) TableName() string {
	return "withdrawals"
}

func (RelayTask) TableName() string {
	return "relay_tasks"
}

func (rt *RelayTask) TaskID() common.Hash {
	return common.HexToHash(rt.ID)
}

func (Tube) TableName() string {
	return "tubes"
}

func (Validator) TableName() string {
	return "validators"
}

func (SignatureTask) TableName() string {
	return "signature_tasks"
}

func NewWithdrawal(e *contract.ERC20TubeSettled) *Withdrawal {
	validators := []string{}
	for _, validator := range e.Validators {
		validators = append(validators, validator.Hex())
	}
	return &Withdrawal{
		ID:          hexutil.Encode(e.Key[:]),
		BlockHeight: e.Raw.BlockNumber,
		TxHash:      e.Raw.TxHash.Hex(),
		Validators:  validators,
	}
}

func NewDeposit(e *contract.ERC20TubeReceipt) *Deposit {
	return &Deposit{
		BlockHeight:       e.Raw.BlockNumber,
		TxHash:            e.Raw.TxHash.Hex(),
		SourceTubeNonce:   e.Nonce.Uint64(),
		SourceToken:       e.Token.Hex(),
		Sender:            e.Sender.Hex(),
		Recipient:         e.Recipient.Hex(),
		DestinationTubeID: e.TargetTubeID.Uint64(),
		Amount:            decimal.NewFromBigInt(e.Amount, 0),
		Fee:               decimal.NewFromBigInt(e.Fee, 0),
	}
}

func NewActiveValidator(e *contract.VerifierValidatorAdded) *Validator {
	return &Validator{
		Address:     e.Validator.Hex(),
		BlockHeight: e.Raw.BlockNumber,
		TxHash:      e.Raw.TxHash.Hex(),
		LogIndex:    e.Raw.Index,
		Active:      true,
	}
}

func NewInactiveValidator(e *contract.VerifierValidatorRemoved) *Validator {
	return &Validator{
		Address:     e.Validator.Hex(),
		BlockHeight: e.Raw.BlockNumber,
		TxHash:      e.Raw.TxHash.Hex(),
		LogIndex:    e.Raw.Index,
		Active:      false,
	}
}

func NewPausedVerifier(e *contract.VerifierPaused) *VerifierStatus {
	return &VerifierStatus{
		TxHash:      e.Raw.TxHash.Hex(),
		BlockHeight: e.Raw.BlockNumber,
		LogIndex:    e.Raw.Index,
		Paused:      true,
	}
}

func NewUnpausedVerifier(e *contract.VerifierUnpaused) *VerifierStatus {
	return &VerifierStatus{
		TxHash:      e.Raw.TxHash.Hex(),
		BlockHeight: e.Raw.BlockNumber,
		LogIndex:    e.Raw.Index,
		Paused:      false,
	}
}

func NewRelayTask(deposit *Deposit) *RelayTask {
	return &RelayTask{
		ID:      deposit.ID,
		Deposit: *deposit,
		Fee:     decimal.New(0, 0),
		Status:  Unknown,
	}
}

func NewSignatureTask(deposit *Deposit) *SignatureTask {
	return &SignatureTask{
		ID:      deposit.ID,
		Deposit: *deposit,
		Status:  Unknown,
		Value:   "",
	}
}
