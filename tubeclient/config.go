package tubeclient

import "time"

// Config is the config of one tube
type Config struct {
	ContractAddr       string        `json:"contractAddr" yaml:"contractAddr"`
	RouterContractAddr string        `json:"routerContractAddr" yaml:"routerContractAddr"`
	ChainAPI           string        `json:"chainAPI" yaml:"chainAPI"`
	BatchSize          int           `json:"batchSize" yaml:"batchSize"`
	HeightToConfirm    int           `json:"heightToConfirm" yaml:"heightToConfirm"`
	Interval           time.Duration `json:"interval" yaml:"interval"`
	StartHeight        int           `json:"startHeight" yaml:"startHeight"`
}

func (tc *Config) UnmarshalYAML(unmarshal func(interface{}) error) error {
	tc.BatchSize = 100
	tc.HeightToConfirm = 10
	tc.Interval = time.Minute
	tc.StartHeight = 0

	type plain Config
	if err := unmarshal((*plain)(tc)); err != nil {
		return err
	}

	return nil
}
