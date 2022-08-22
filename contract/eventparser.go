package contract

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type (
	Event struct {
		Parsed interface{}
		Height uint64
	}
	EventParser struct {
		abi         abi.ABI
		skipTopics  map[common.Hash]struct{}
		eventByName func(string, types.Log) interface{}
	}
)

func (parser *EventParser) Parse(log types.Log) (*Event, error) {
	if _, ok := parser.skipTopics[log.Topics[0]]; ok {
		return nil, nil
	}
	event, err := parser.abi.EventByID(log.Topics[0])
	if err != nil {
		return nil, err
	}
	if log.Topics[0] != event.ID {
		panic("event signature mismatch")
	}
	out := parser.eventByName(event.Name, log)
	if out == nil {
		return nil, nil
	}
	if len(log.Data) > 0 {
		if err := parser.abi.UnpackIntoInterface(out, event.Name, log.Data); err != nil {
			return nil, err
		}
	}
	var indexed abi.Arguments
	for _, arg := range event.Inputs {
		if arg.Indexed {
			indexed = append(indexed, arg)
		}
	}
	if err := abi.ParseTopics(out, indexed, log.Topics[1:]); err != nil {
		return nil, err
	}

	return &Event{Parsed: out, Height: log.BlockNumber}, nil
}

func NewVerifierEventParser() *EventParser {
	verifierABI, err := abi.JSON(strings.NewReader(VerifierABI))
	if err != nil {
		panic(err)
	}

	return &EventParser{
		abi: verifierABI,
		skipTopics: map[common.Hash]struct{}{
			common.HexToHash("0x9a174d577c7665e63298da21ff932b957bb0dc51d46920b0cc1e32f22437d8fd"): {},
		},
		eventByName: func(name string, log types.Log) interface{} {
			switch name {
			case "ValidatorAdded":
				return &VerifierValidatorAdded{Raw: log}
			case "ValidatorRemoved":
				return &VerifierValidatorRemoved{Raw: log}
			default:
				return nil
			}
		},
	}
}

func NewERC20TubeEventParser() *EventParser {
	tubeABI, err := abi.JSON(strings.NewReader(ERC20TubeABI))
	if err != nil {
		panic(err)
	}

	return &EventParser{
		abi:        tubeABI,
		skipTopics: make(map[common.Hash]struct{}),
		eventByName: func(name string, log types.Log) interface{} {
			switch name {
			case "Receipt":
				return &ERC20TubeReceipt{Raw: log}
			case "Settled":
				return &ERC20TubeSettled{Raw: log}
			default:
				return nil
			}
		},
	}
}

func NewERC20RouterEventParser() *EventParser {
	routerABI, err := abi.JSON(strings.NewReader(ERC20TubeRouterABI))
	if err != nil {
		panic(err)
	}

	return &EventParser{
		abi: routerABI,
		eventByName: func(name string, log types.Log) interface{} {
			if name != "RelayFeeReceipt" {
				return nil
			}
			return &ERC20TubeRouterRelayFeeReceipt{Raw: log}
		},
	}
}
