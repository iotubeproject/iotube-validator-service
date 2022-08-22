package relayer

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"log"
	"math/big"
	"time"

	"github.com/crosschained-io/crosschained-service/tubeclient"
	"github.com/crosschained-io/crosschained-service/types"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type (
	Account struct {
		PrivateKey   *ecdsa.PrivateKey
		Address      common.Address
		PendingNonce uint64
		Nonce        uint64
	}

	// Config is the config of a relay service
	Config struct {
		ChainAPI     string `json:"chainAPI" yaml:"chainAPI"`
		ContractAddr string `json:"contractAddr" yaml:"contractAddr"`

		PrivateKeys        string `json:"privateKeys" yaml:"privateKeys"`
		GasLimit           uint64 `json:"gasLimit" yaml:"gasLimit"`
		GasPriceUpperBound string `json:"gasPriceUpperBound" yaml:"gasPriceUpperBound"`
	}

	Relayer struct {
		recorder           *Recorder
		client             *ethclient.Client
		gasLimit           uint64
		gasPriceUpperBound *big.Int
		tube               *tubeclient.EthereumClient
		tubeID             uint64
		accounts           []*Account
	}
)

func newAccount(pk *ecdsa.PrivateKey) *Account {
	addr := crypto.PubkeyToAddress(pk.PublicKey)
	relayer := &Account{
		PrivateKey: pk,
		Address:    addr,
	}
	return relayer
}

func (account *Account) transactionOpts(chainID *big.Int, nonce uint64, gasLimit uint64, gasPrice *big.Int) (*bind.TransactOpts, error) {
	opts, err := bind.NewKeyedTransactorWithChainID(account.PrivateKey, chainID)
	if err != nil {
		return nil, err
	}
	opts.Value = big.NewInt(0)
	opts.GasLimit = gasLimit
	opts.GasPrice = gasPrice
	opts.Nonce = new(big.Int).SetUint64(nonce)

	return opts, nil
}

func NewRelayer(
	cfg Config,
	recorder *Recorder,
	privateKeys []*ecdsa.PrivateKey,
) (*Relayer, error) {
	client, err := ethclient.DialContext(context.Background(), cfg.ChainAPI)
	if err != nil {
		return nil, err
	}
	tube, err := tubeclient.NewEthereumClient(
		client,
		cfg.ContractAddr,
		"",
	)
	if err != nil {
		return nil, err
	}
	id, err := tube.ID()
	if err != nil {
		return nil, err
	}
	gasPriceUpperBound, ok := new(big.Int).SetString(cfg.GasPriceUpperBound, 10)
	if !ok {
		return nil, errors.Errorf("failed to cast %s to big.Int", cfg.GasPriceUpperBound)
	}

	accounts := []*Account{}
	for _, privateKey := range privateKeys {
		accounts = append(accounts, newAccount(privateKey))
	}
	return &Relayer{
		client:             client,
		recorder:           recorder,
		gasLimit:           cfg.GasLimit,
		gasPriceUpperBound: gasPriceUpperBound,
		tube:               tube,
		tubeID:             id.Uint64(),
		accounts:           accounts,
	}, nil
}

func (relayer *Relayer) Size() int {
	return len(relayer.accounts)
}

func (relayer *Relayer) Start(ctx context.Context) error {
	for _, account := range relayer.accounts {
		if err := relayer.reloadNonce(account); err != nil {
			return err
		}
	}

	return nil
}

func (relayer *Relayer) Stop(ctx context.Context) error {
	relayer.client.Close()
	return nil
}

func (relayer *Relayer) TubeID() uint64 {
	return relayer.tubeID
}

func (relayer *Relayer) Run() error {
	tasks, err := relayer.recorder.TasksToRelay(relayer.tubeID, 0, relayer.Size())
	if err != nil {
		return err
	}
	if len(tasks) == 0 {
		return nil
	}
	idx := 0
	for _, account := range relayer.accounts {
		if err := relayer.reloadNonce(account); err != nil {
			return err
		}
		ready, err := relayer.checkAccount(account)
		if err != nil {
			return err
		}
		if !ready {
			continue
		}
		if idx >= len(tasks) {
			continue
		}
		tx, err := relayer.relay(account, &tasks[idx], false)
		if err != nil {
			return err
		}
		if tx == nil {
			// already settled
			continue
		}
		if err := relayer.recorder.UpdateRelayTask(
			tasks[idx].TaskID(),
			account.Address,
			tx.Nonce(),
			tx.GasPrice(),
			tx.Hash(),
		); err != nil {
			return err
		}
		idx++
	}
	return nil
}

func (relayer *Relayer) checkAccount(account *Account) (bool, error) {
	task, err := relayer.recorder.LatestTask(relayer.tubeID, account.Address)
	switch errors.Cause(err) {
	default:
		return false, err
	case gorm.ErrRecordNotFound:
		// do nothing
	case nil:
		if task.Nonce >= account.PendingNonce {
			// TODO: need to clean up this task
			return false, nil
		}
		if task.Nonce > account.Nonce {
			// do nothing
			return false, nil
		}
		if task.Nonce == account.Nonce {
			if task.UpdatedAt.Add(10 * time.Minute).Before(time.Now()) {
				tx, err := relayer.relay(account, task, true)
				if err != nil {
					return false, err
				}
				if tx == nil {
					return true, nil
				}
				if err := relayer.recorder.UpdateRelayTask(
					task.TaskID(),
					account.Address,
					tx.Nonce(),
					tx.GasPrice(),
					tx.Hash(),
				); err != nil {
					return false, err
				}
			}
			return false, nil
		}
	}
	return true, nil
}

func (relayer *Relayer) reloadNonce(account *Account) error {
	pendingNonce, err := relayer.tube.PendingNonceAt(account.Address)
	if err != nil {
		return errors.Wrapf(err, "failed to get pending nonce of %s", account.Address.Hex())
	}
	account.PendingNonce = pendingNonce
	nonce, err := relayer.tube.NonceAt(account.Address)
	if err != nil {
		return errors.Wrapf(err, "failed to get nonce of %s", account.Address.Hex())
	}
	account.Nonce = nonce

	return nil
}

func (relayer *Relayer) relay(account *Account, task *types.RelayTask, speedup bool) (*ethtypes.Transaction, error) {
	settled, err := relayer.tube.IsSettled(task.TaskID())
	if err != nil {
		return nil, err
	}
	if settled {
		return nil, nil
	}
	signatures := [][]byte{}
	for _, testimony := range task.Testimonies {
		signature, err := hex.DecodeString(testimony.Signature)
		if err != nil {
			return nil, err
		}
		signatures = append(signatures, signature)
	}
	gasPrice, err := relayer.tube.SuggestGasPrice()
	if err != nil {
		return nil, errors.Wrapf(err, "failed to get suggested gas price")
	}
	if gasPrice.Cmp(relayer.gasPriceUpperBound) > 0 {
		return nil, errors.Wrapf(errGasPriceTooHigh, "suggested gas price %d > limit %d", gasPrice, relayer.gasPriceUpperBound)
	}
	balance, err := relayer.tube.BalanceAt(account.Address)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to get balance of operator account")
	}
	if new(big.Int).Mul(new(big.Int).SetUint64(relayer.gasLimit), gasPrice).Cmp(balance) > 0 {
		return nil, errors.Errorf("insufficient balance for gas fee")
	}
	chainID, err := relayer.tube.ChainID()
	if err != nil {
		return nil, err
	}
	var opts *bind.TransactOpts
	if speedup {
		if task.GasPrice.BigInt().Cmp(gasPrice) >= 0 {
			log.Printf("suggested gas price %d is not higher than original gas price %d\n", gasPrice, task.GasPrice)
			return nil, nil
		}
		opts, err = account.transactionOpts(chainID, task.Nonce, relayer.gasLimit, gasPrice)
	} else {
		opts, err = account.transactionOpts(chainID, account.PendingNonce, relayer.gasLimit, gasPrice)
	}
	if err != nil {
		return nil, err
	}

	log.Printf(
		"relay task %s: %d from (%d, %s, %s) to (%d, %s, %s)\n",
		task.ID,
		task.Deposit.Amount.BigInt(),
		task.Deposit.SourceTubeID,
		task.Deposit.SourceToken,
		task.Deposit.Sender,
		task.Deposit.DestinationTubeID,
		task.Deposit.DestinationToken,
		task.Deposit.Recipient,
	)

	return relayer.tube.Submit(
		opts,
		task.Deposit.SourceTubeID,
		task.Deposit.SourceTubeNonce,
		common.HexToAddress(task.Deposit.DestinationToken),
		common.HexToAddress(task.Deposit.Recipient),
		task.Deposit.Amount.BigInt(),
		signatures,
	)
}
