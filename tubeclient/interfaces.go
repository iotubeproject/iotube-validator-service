package tubeclient

import (
	"github.com/ethereum/go-ethereum/core/types"
)

type (
	LogSubscriber interface {
		ProcessLogs(uint64, uint64, uint64, uint64, ...types.Log) error
	}

	// NextHeight returns the next height to fetch for tube
	NextHeight func(uint64) (uint64, error)
)
