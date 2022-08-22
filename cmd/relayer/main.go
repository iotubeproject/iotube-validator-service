package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/crosschained-io/crosschained-service/monitor"
	"github.com/crosschained-io/crosschained-service/registry"
	"github.com/crosschained-io/crosschained-service/relayer"

	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/config"
)

type (
	// MonitorConfig defines the configuration of the relay service
	MonitorConfig struct {
		AssetRegistry       registry.AssetRegistryConfig     `json:"assetRegistry" yaml:"assetRegistry"`
		ValidatorRegistry   registry.ValidatorRegistryConfig `json:"validatorRegistry" yaml:"validatorRegistry"`
		Interval            time.Duration                    `json:"interval" yaml:"interval"`
		IoTeXChainAPI       string                           `json:"iotexChainAPI" yaml:"iotexChainAPI"`
		DatabaseURL         string                           `json:"databaseURL" yaml:"databaseURL"`
		ExplorerDatabaseURL string                           `json:"explorerDatabaseURL" yaml:"explorerDatabaseURL"`
		ExplorerTableName   string                           `json:"explorerTableName" yaml:"explorerTableName"`
		Monitors            []monitor.MonitorConfig          `json:"monitors" yaml:"monitors"`
	}
	// RelayConfig defines the configuration of the relay service
	RelayConfig struct {
		Interval      time.Duration    `json:"interval" yaml:"interval"`
		DatabaseURL   string           `json:"databaseURL" yaml:"databaseURL"`
		RelayServices []relayer.Config `json:"relayServices" yaml:"relayServices"`
	}

	StarterStopper interface {
		Start(context.Context) error
		Stop(context.Context) error
	}
)

var monitorConfigFile = flag.String("monitor", "", "path of monitor config file")

var relayConfigFile = flag.String("relay", "", "path of relay config file")

func init() {
	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, "Usage:", os.Args[0], "-relay <filename> -monitor <filename>")
		flag.PrintDefaults()
	}
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
	return monitor.NewServiceForRelayer(
		cfg.Monitors,
		cfg.IoTeXChainAPI,
		cfg.AssetRegistry,
		cfg.ValidatorRegistry,
		cfg.DatabaseURL,
		cfg.ExplorerDatabaseURL,
		cfg.ExplorerTableName,
		cfg.Interval,
	)
}

func createRelayService() (*relayer.Service, error) {
	if *relayConfigFile == "" {
		return nil, nil
	}
	yaml, err := config.NewYAML(config.File(*relayConfigFile))
	if err != nil {
		return nil, err
	}
	var cfg RelayConfig
	if err := yaml.Get(config.Root).Populate(&cfg); err != nil {
		return nil, err
	}
	if pks, ok := os.LookupEnv("RELAYER_PRIVATE_KEYS"); ok {
		for i := range cfg.RelayServices {
			cfg.RelayServices[i].PrivateKeys += "," + pks
		}
	}
	if dburl, ok := os.LookupEnv("RELAYER_DB_URL"); ok {
		cfg.DatabaseURL = dburl
	}
	if cfg.Interval == 0 {
		cfg.Interval = time.Minute
	}
	return relayer.NewService(
		cfg.RelayServices,
		cfg.DatabaseURL,
		cfg.Interval,
	)
}

func main() {
	flag.Parse()
	ctx := context.Background()
	services := []StarterStopper{}
	defer func() {
		for i := len(services) - 1; i >= 0; i-- {
			if err := services[i].Stop(ctx); err != nil {
				log.Fatal(err)
			}
		}
	}()
	relayService, err := createRelayService()
	if err != nil {
		log.Fatal("failed to create relay service", err)
	}
	if relayService != nil {
		if err := relayService.Start(ctx); err != nil {
			log.Fatal(err)
		}
		services = append(services, relayService)
	}
	monitorService, err := createMonitorService()
	if err != nil {
		log.Fatal("failed to create relay service", err)
	}
	if monitorService != nil {
		if err := monitorService.Start(ctx); err != nil {
			log.Fatal(err)
		}
		services = append(services, monitorService)
	}

	log.Println("Serving...")
	select {}
}
