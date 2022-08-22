package registry

type (
	AssetInfo struct {
		ID      uint64 `json:"id" yaml:"id"`
		Address string `json:"address" yaml:"address"`
		TubeID  uint64 `json:"tubeID" yaml:"tubeID"`
	}

	AssetRegistryConfig struct {
		ContractAddr string      `json:"contractAddr" yaml:"contractAddr"`
		StartHeight  uint64      `json:"startHeight" yaml:"startHeight"`
		BatchSize    uint64      `json:"batchSize" yaml:"batchSize"`
		Assets       []AssetInfo `json:"assets" yaml:"assets"`
		// 0: local first
		// 1: contract first
		// 2: local only
		// 3: contract only
		Mode int `json:"mode" yaml:"mode"`
	}

	ValidatorInfo struct {
		Genre uint64 `json:"genre" yaml:"genre"`
		URI   string `json:"uri" yaml:"uri"`
	}

	ValidatorRegistryConfig struct {
		ContractAddr string                   `json:"contractAddr" yaml:"contractAddr"`
		StartHeight  uint64                   `json:"startHeight" yaml:"startHeight"`
		BatchSize    uint64                   `json:"batchSize" yaml:"batchSize"`
		Validators   map[string]ValidatorInfo `json:"validators" yaml:"validators"`
		// 0: local first
		// 1: contract first
		// 2: local only
		// 3: contract only
		Mode int `json:"mode" yaml:"mode"`
	}

	AssetUpperBoundRegistryConfig struct {
		ContractAddr string            `json:"contractAddr" yaml:"contractAddr"`
		StartHeight  uint64            `json:"startHeight" yaml:"startHeight"`
		BatchSize    uint64            `json:"batchSize" yaml:"batchSize"`
		UpperBounds  map[uint64]string `json:"upperBounds" yaml:"upperBounds"`
		// 0: local first
		// 1: contract first
		// 2: local only
		// 3: contract only
		Mode int `json:"mode" yaml:"mode"`
	}
)
