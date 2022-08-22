package tubeclient

import (
	"context"
	"fmt"
	"math/big"
	"strings"

	"github.com/crosschained-io/crosschained-service/contract"
	"github.com/pkg/errors"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type (
	EthereumClient struct {
		client               *ethclient.Client
		query                ethereum.FilterQuery
		tubeABI              abi.ABI
		tubeContractAddr     common.Address
		routerABI            abi.ABI
		routerContractAddr   common.Address
		verifierABI          abi.ABI
		verifierContractAddr common.Address
		tubeContract         *contract.ERC20Tube
	}
	Event struct {
		Parsed interface{}
		Raw    types.Log
	}
)

func NewEthereumClient(
	client *ethclient.Client,
	tubeContractAddr string,
	tubeRouterContractAddr string,
) (*EthereumClient, error) {
	tubeAddr := common.HexToAddress(tubeContractAddr)
	query := ethereum.FilterQuery{
		Addresses: []common.Address{tubeAddr},
	}
	routerAddr := common.Address{}
	if tubeRouterContractAddr != "" {
		routerAddr = common.HexToAddress(tubeRouterContractAddr)
		query.Addresses = append(query.Addresses, routerAddr)
	}

	tubeContract, err := contract.NewERC20Tube(tubeAddr, client)
	if err != nil {
		return nil, err
	}
	verifierContractAddr, err := tubeContract.Verifier(nil)
	if err != nil {
		return nil, err
	}
	fmt.Println("verifier:", verifierContractAddr)
	query.Addresses = append(query.Addresses, verifierContractAddr)
	tubeABI, err := abi.JSON(strings.NewReader(contract.ERC20TubeABI))
	if err != nil {
		return nil, err
	}
	routerABI, err := abi.JSON(strings.NewReader(contract.ERC20TubeRouterABI))
	if err != nil {
		return nil, err
	}
	verifierABI, err := abi.JSON(strings.NewReader(contract.VerifierABI))
	if err != nil {
		return nil, err
	}

	return &EthereumClient{
		client:               client,
		query:                query,
		tubeABI:              tubeABI,
		tubeContractAddr:     tubeAddr,
		routerABI:            routerABI,
		routerContractAddr:   routerAddr,
		verifierABI:          verifierABI,
		verifierContractAddr: verifierContractAddr,
		tubeContract:         tubeContract,
	}, nil
}

func (et *EthereumClient) ID() (*big.Int, error) {
	return et.tubeContract.TubeID(nil)
}

func (et *EthereumClient) ChainID() (*big.Int, error) {
	return et.client.ChainID(context.Background())
}

func (et *EthereumClient) PendingNonceAt(addr common.Address) (uint64, error) {
	return et.client.PendingNonceAt(context.Background(), addr)
}

func (et *EthereumClient) NonceAt(addr common.Address) (uint64, error) {
	return et.client.NonceAt(context.Background(), addr, nil)
}

func (et *EthereumClient) SuggestGasPrice() (*big.Int, error) {
	return et.client.SuggestGasPrice(context.Background())
}

func (et *EthereumClient) BalanceAt(addr common.Address) (*big.Int, error) {
	return et.client.BalanceAt(context.Background(), addr, nil)
}

func (et *EthereumClient) TubeContractAddress() common.Address {
	return et.tubeContractAddr
}

func (et *EthereumClient) IsSettled(id common.Hash) (bool, error) {
	return et.tubeContract.IsSettled(nil, id)
}

func (et *EthereumClient) TaskID(_srcTubeID *big.Int, _nonce *big.Int, _token common.Address, _recipient common.Address, _amount *big.Int) (common.Hash, error) {
	return et.tubeContract.GenKey(nil, _srcTubeID, _nonce, _token, _recipient, _amount)
}

func (et *EthereumClient) Submit(
	opts *bind.TransactOpts,
	tubeID uint64,
	idx uint64,
	token common.Address,
	recipient common.Address,
	amount *big.Int,
	signatures [][]byte,
) (*types.Transaction, error) {
	s := []byte{}
	for _, signature := range signatures {
		s = append(s, signature...)
	}

	return et.tubeContract.Withdraw(opts, new(big.Int).SetUint64(tubeID), new(big.Int).SetUint64(idx), token, recipient, amount, s)
}

func (et *EthereumClient) FetchLogs(startHeight, endHeight uint64) (uint64, []*Event, error) {
	tip, err := et.client.BlockNumber(context.Background())
	if err != nil {
		return 0, nil, errors.Wrap(err, "failed to get tip height")
	}
	if endHeight > tip {
		endHeight = tip
	}
	if endHeight < startHeight {
		return tip, nil, nil
	}
	et.query.FromBlock = new(big.Int).SetUint64(startHeight)
	et.query.ToBlock = new(big.Int).SetUint64(endHeight)
	logs, err := et.client.FilterLogs(context.Background(), et.query)
	if err != nil {
		return 0, nil, err
	}
	events := make([]*Event, 0, len(logs))
	var event *Event
	for _, log := range logs {
		switch log.Address {
		case et.tubeContractAddr:
			event, err = et.parseTubeEvent(log)
		case et.routerContractAddr:
			event, err = et.parseRouterEvent(log)
		case et.verifierContractAddr:
			event, err = et.parseVerifierEvent(log)
		default:
			return 0, nil, errors.Errorf("unexpected contract address %s", log.Address)
		}
		if err != nil {
			return 0, nil, err
		}
		if event != nil {
			events = append(events, event)
			event = nil
		}
	}

	return tip, events, nil
}

func (et *EthereumClient) parseTubeEvent(log types.Log) (*Event, error) {
	event, err := et.tubeABI.EventByID(log.Topics[0])
	if err != nil {
		return nil, err
	}
	switch event.Name {
	case "Receipt":
		parsed, err := et.tubeContract.ParseReceipt(log)
		if err != nil {
			return nil, err
		}
		return &Event{
			Parsed: parsed,
			Raw:    log,
		}, nil
	case "Settled":
		parsed, err := et.tubeContract.ParseSettled(log)
		if err != nil {
			return nil, err
		}
		return &Event{
			Parsed: parsed,
			Raw:    log,
		}, nil
	}

	return nil, nil
}

func (et *EthereumClient) parseRouterEvent(log types.Log) (*Event, error) {
	event, err := et.routerABI.EventByID(log.Topics[0])
	if err != nil {
		return nil, err
	}
	if event.Name == "RelayFeeReceipt" {
		receipt := new(contract.ERC20TubeRouterRelayFeeReceipt)
		err := UnpackLog(receipt, et.routerABI, event.Name, log)
		if err != nil {
			return nil, err
		}
		receipt.Raw = log

		return &Event{
			Parsed: receipt,
			Raw:    log,
		}, err
	}

	return nil, nil
}

func (et *EthereumClient) parseVerifierEvent(log types.Log) (*Event, error) {
	event, err := et.verifierABI.EventByID(log.Topics[0])
	if err != nil {
		return nil, err
	}
	switch event.Name {
	case "ValidatorAdded":
		va := new(contract.VerifierValidatorAdded)
		err := UnpackLog(va, et.verifierABI, event.Name, log)
		if err != nil {
			return nil, err
		}
		va.Raw = log

		return &Event{
			Parsed: va,
			Raw:    log,
		}, err
	case "ValidatorRemoved":
		vr := new(contract.VerifierValidatorRemoved)
		err := UnpackLog(vr, et.verifierABI, event.Name, log)
		if err != nil {
			return nil, err
		}
		vr.Raw = log

		return &Event{
			Parsed: vr,
			Raw:    log,
		}, err
	}

	return nil, nil
}

func UnpackLog(out interface{}, a abi.ABI, event string, log types.Log) error {
	if log.Topics[0] != a.Events[event].ID {
		return errors.New("event signature mismatch")
	}
	if len(log.Data) > 0 {
		if err := a.UnpackIntoInterface(out, event, log.Data); err != nil {
			return err
		}
	}
	var indexed abi.Arguments
	for _, arg := range a.Events[event].Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}

	return abi.ParseTopics(out, indexed, log.Topics[1:])
}
