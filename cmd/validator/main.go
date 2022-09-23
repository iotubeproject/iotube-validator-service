package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"math/big"
	"os"
	"time"

	"github.com/crosschained-io/crosschained-service/api"
	"github.com/crosschained-io/crosschained-service/monitor"
	"github.com/crosschained-io/crosschained-service/registry"
	"github.com/crosschained-io/crosschained-service/validator"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"
	"go.uber.org/config"
)

type (
	PacingConfig struct {
		Duration     time.Duration `json:"duration" yaml:"duration"`
		Amount       string        `json:"amount" yaml:"amount"`
		QuotaPerHour string        `json:"quotaPerHour" yaml:"quotaPerHour"`
	}

	// APIConfig defines the configuration of the api query service
	APIConfig struct {
		DatabaseURL string `json:"databaseURL" yaml:"databaseURL"`
		GrpcPort    int    `json:"grpcPort" yaml:"grpcPort"`
		HttpPort    int    `json:"httpPort" yaml:"httpPort"`
		Validator   string `json:"validator" yaml:"validator"`
	}

	// SignerConfig defines the configuration of the signer service
	SignerConfig struct {
		Interval    time.Duration           `json:"interval" yaml:"interval"`
		DatabaseURL string                  `json:"databaseURL" yaml:"databaseURL"`
		PrivateKey  string                  `json:"privateKey" yaml:"privateKey"`
		Tubes       map[uint64]PacingConfig `json:"tubes" yaml:"tubes"`
	}

	// MonitorConfig defines the configuration of the monitor service
	MonitorConfig struct {
		AssetRegistry           registry.AssetRegistryConfig           `json:"assetRegistry" yaml:"assetRegistry"`
		AssetUpperBoundRegistry registry.AssetUpperBoundRegistryConfig `json:"assetUppserBoundRegistry" yaml:"assetUpperBoundRegistry"`
		ValidatorRegistry       registry.ValidatorRegistryConfig       `json:"validatorRegistry" yaml:"validatorRegistry"`
		Interval                time.Duration                          `json:"interval" yaml:"interval"`
		IoTeXChainAPI           string                                 `json:"iotexChainAPI" yaml:"iotexChainAPI"`
		DatabaseURL             string                                 `json:"databaseURL" yaml:"databaseURL"`
		Monitors                []monitor.MonitorConfig                `json:"monitors" yaml:"monitors"`
	}

	StarterStopper interface {
		Start(context.Context) error
		Stop(context.Context) error
	}
)

var (
	signerConfigFile  = flag.String("signer", "", "path of signer config file")
	monitorConfigFile = flag.String("monitor", "", "path of monitor config file")
	apiConfigFile     = flag.String("api", "", "path of api config file")

	defaultConfig = APIConfig{
		GrpcPort: 9000,
		HttpPort: 9001,
	}
)

func init() {
	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, "Usage:", os.Args[0], "-signer <filename> -monitor <filename> -api <filename>")
		flag.PrintDefaults()
	}
}

func createAPIService() (*api.APIService, error) {
	if *apiConfigFile == "" {
		return nil, nil
	}
	yaml, err := config.NewYAML(
		config.Static(defaultConfig),
		config.File(*apiConfigFile),
	)
	if err != nil {
		return nil, err
	}
	var cfg APIConfig
	if err := yaml.Get(config.Root).Populate(&cfg); err != nil {
		return nil, err
	}
	if dburl, ok := os.LookupEnv("API_DB_URL"); ok {
		cfg.DatabaseURL = dburl
	}
	validatorAddr := common.HexToAddress(cfg.Validator)
	log.Println("validator:", validatorAddr)

	return api.NewAPIService(
		cfg.GrpcPort,
		cfg.HttpPort,
		cfg.DatabaseURL,
		validatorAddr,
	)
}

func createMonitorService() (*monitor.Service, error) {
	if *monitorConfigFile == "" {
		return nil, nil
	}
	yaml, err := config.NewYAML(config.File(*monitorConfigFile))
	if err != nil {
		return nil, err
	}
	var cfg MonitorConfig
	if err := yaml.Get(config.Root).Populate(&cfg); err != nil {
		return nil, err
	}
	if dburl, ok := os.LookupEnv("MONITOR_DB_URL"); ok {
		cfg.DatabaseURL = dburl
	}
	if cfg.Interval == 0 {
		cfg.Interval = time.Minute
	}
	return monitor.NewServiceForValidator(
		cfg.Monitors,
		cfg.IoTeXChainAPI,
		cfg.AssetRegistry,
		cfg.AssetUpperBoundRegistry,
		cfg.ValidatorRegistry,
		cfg.DatabaseURL,
		cfg.Interval,
	)
}

func createSignerService() (*validator.Service, error) {
	if *signerConfigFile == "" {
		return nil, nil
	}
	yaml, err := config.NewYAML(
		config.File(*signerConfigFile),
	)
	if err != nil {
		return nil, err
	}
	var signerCfg SignerConfig
	if err := yaml.Get(config.Root).Populate(&signerCfg); err != nil {
		return nil, err
	}
	if pk, ok := os.LookupEnv("VALIDATOR_PRIVATE_KEY"); ok {
		signerCfg.PrivateKey = pk
	}
	if dburl, ok := os.LookupEnv("VALIDATOR_DB_URL"); ok {
		signerCfg.DatabaseURL = dburl
	}
	if signerCfg.Interval == 0 {
		signerCfg.Interval = time.Minute
	}
	privateKey, err := crypto.HexToECDSA(signerCfg.PrivateKey)
	if err != nil {
		return nil, errors.Wrap(err, "failed to decode private key\n")
	}
	log.Println("validator address", crypto.PubkeyToAddress(privateKey.PublicKey))
	ids := []uint64{}
	quotas := []*big.Int{}
	amounts := []*big.Int{}
	durations := []time.Duration{}
	for id, tube := range signerCfg.Tubes {
		ids = append(ids, id)
		quota, ok := new(big.Int).SetString(tube.QuotaPerHour, 10)
		if !ok {
			return nil, errors.Errorf("failed to parse quota '%s' for tube %d", tube.QuotaPerHour, id)
		}
		quotas = append(quotas, quota)
		amount, ok := new(big.Int).SetString(tube.Amount, 10)
		if !ok {
			return nil, errors.Errorf("failed to parse amount '%s' for tube %d", tube.Amount, id)
		}
		amounts = append(amounts, amount)
		durations = append(durations, tube.Duration)
	}

	return validator.NewService(privateKey, ids, quotas, durations, amounts, signerCfg.DatabaseURL, signerCfg.Interval), nil
}

func main() {
	flag.Parse()
	ctx := context.Background()
	services := []StarterStopper{}
	defer func() {
		for i := len(services) - 1; i >= 0; i-- {
			if err := services[i].Stop(ctx); err != nil {
				log.Fatal("failed to stop api service:", err)
			}
		}
	}()
	apiService, err := createAPIService()
	if err != nil {
		log.Fatal("failed to create api service:", err)
	}
	if apiService != nil {
		if err := apiService.Start(context.Background()); err != nil {
			log.Fatal("failed to start api service:", err)
		}
		services = append(services, apiService)
	}
	monitorService, err := createMonitorService()
	if err != nil {
		log.Fatal("failed to create monitor service:", err)
	}
	if monitorService != nil {
		if err := monitorService.Start(context.Background()); err != nil {
			log.Fatal("failed to start monitor service:", err)
		}
		services = append(services, monitorService)
	}
	signerService, err := createSignerService()
	if err != nil {
		log.Fatal("failed to create signer service:", err)
	}
	if signerService != nil {
		if err := signerService.Start(context.Background()); err != nil {
			log.Fatal("failed to start validator service:", err)
		}
		services = append(services, signerService)
	}
	log.Println("Serving...")
	select {}
}
