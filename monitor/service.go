package monitor

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/crosschained-io/crosschained-service/registry"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
)

type (
	Service struct {
		db                      *gorm.DB
		chainClient             *ethclient.Client
		interval                time.Duration
		assetRegistry           *registry.AssetRegistry
		assetUpperBoundRegistry *registry.AssetUpperBoundRegistry
		validatorRegistry       *registry.ValidatorRegistry
		testimonyDAO            *registry.TestimonyDAO
		monitors                []Monitor
		terminate               chan struct{}
	}
)

func NewServiceForValidator(
	cfgs []MonitorConfig,
	chainAPI string,
	assetRegistryCfg registry.AssetRegistryConfig,
	assetUpperBoundRegistryCfg registry.AssetUpperBoundRegistryConfig,
	validatorRegistryCfg registry.ValidatorRegistryConfig,
	databaseURL string,
	interval time.Duration,
) (*Service, error) {
	if len(cfgs) == 0 {
		return nil, errors.New("no monitor configs")
	}
	chainClient, err := ethclient.DialContext(context.Background(), chainAPI)
	if err != nil {
		return nil, err
	}
	db, err := gorm.Open(postgres.Open(databaseURL), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		return nil, err
	}
	recorder, err := NewRecorderForValidator(db)
	if err != nil {
		return nil, err
	}

	assetRegistry, err := registry.NewAssetRegistry(assetRegistryCfg, chainClient, db)
	if err != nil {
		return nil, errors.Wrap(err, "failed to generate asset registry")
	}
	assetUpperBoundRegistry, err := registry.NewAssetUpperBoundRegistry(assetUpperBoundRegistryCfg, chainClient, db)
	if err != nil {
		return nil, errors.Wrap(err, "failed to generate asset upper bound registry")
	}
	validatorRegistry, err := registry.NewValidatorRegistry(validatorRegistryCfg, chainClient, db)
	if err != nil {
		return nil, errors.Wrap(err, "failed to generate validator registry")
	}

	service := &Service{
		db:                      db,
		chainClient:             chainClient,
		testimonyDAO:            registry.NewTestimonyDAO(chainClient, validatorRegistry.Validator),
		assetRegistry:           assetRegistry,
		assetUpperBoundRegistry: assetUpperBoundRegistry,
		validatorRegistry:       validatorRegistry,
		interval:                interval,
		terminate:               make(chan struct{}),
		monitors:                []Monitor{},
	}
	for _, cfg := range cfgs {
		fmt.Println("creating validator monitor for tube", cfg.ContractAddr)
		monitor, err := NewValidatorMonitor(
			cfg.ChainAPI,
			common.HexToAddress(cfg.Validator),
			common.HexToAddress(cfg.ContractAddr),
			cfg.StartHeight,
			uint64(cfg.HeightToConfirm),
			uint64(cfg.BatchSize),
			assetRegistry,
			assetUpperBoundRegistry,
			recorder,
		)
		if err != nil {
			return nil, err
		}
		service.monitors = append(service.monitors, monitor)
	}

	return service, nil
}

func NewServiceForRelayer(
	cfgs []MonitorConfig,
	chainAPI string,
	assetRegistryCfg registry.AssetRegistryConfig,
	validatorRegistryCfg registry.ValidatorRegistryConfig,
	databaseURL string,
	explorerDatabaseURL string,
	explorerTableName string,
	interval time.Duration,
) (*Service, error) {
	if len(cfgs) == 0 {
		return nil, errors.New("no monitor configs")
	}
	chainClient, err := ethclient.DialContext(context.Background(), chainAPI)
	if err != nil {
		return nil, err
	}
	db, err := gorm.Open(postgres.Open(databaseURL), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		return nil, err
	}
	var explorerDB *ExplorerDB
	if explorerDatabaseURL != "" && explorerTableName != "" {
		explorerDB, err = NewExplorerDB(explorerDatabaseURL, explorerTableName)
		if err != nil {
			return nil, err
		}
	}
	recorder, err := NewRecorderForRelayer(db, explorerDB)
	if err != nil {
		return nil, err
	}
	assetRegistry, err := registry.NewAssetRegistry(assetRegistryCfg, chainClient, db)
	if err != nil {
		return nil, err
	}
	validatorRegistry, err := registry.NewValidatorRegistry(validatorRegistryCfg, chainClient, db)
	if err != nil {
		return nil, err
	}
	service := &Service{
		db:                db,
		chainClient:       chainClient,
		testimonyDAO:      registry.NewTestimonyDAO(chainClient, validatorRegistry.Validator),
		assetRegistry:     assetRegistry,
		validatorRegistry: validatorRegistry,
		interval:          interval,
		terminate:         make(chan struct{}),
		monitors:          []Monitor{},
	}
	for _, cfg := range cfgs {
		var monitor *RelayerMonitor
		if cfg.RouterContractAddr == "" {
			monitor, err = NewRelayerMonitorForAllTasks(
				cfg.ChainAPI,
				common.HexToAddress(cfg.ContractAddr),
				cfg.StartHeight,
				uint64(cfg.HeightToConfirm),
				uint64(cfg.BatchSize),
				assetRegistry,
				recorder,
				service.testimonyDAO,
			)
		} else {
			monitor, err = NewRelayerMonitorForRouter(
				cfg.ChainAPI,
				common.HexToAddress(cfg.ContractAddr),
				common.HexToAddress(cfg.RouterContractAddr),
				cfg.StartHeight,
				uint64(cfg.HeightToConfirm),
				uint64(cfg.BatchSize),
				assetRegistry,
				recorder,
				service.testimonyDAO,
			)
		}
		if err != nil {
			return nil, err
		}
		service.monitors = append(service.monitors, monitor)
	}

	return service, nil
}

func (service *Service) Start(ctx context.Context) error {
	if err := service.assetRegistry.Start(ctx); err != nil {
		return err
	}
	if service.assetUpperBoundRegistry != nil {
		if err := service.assetUpperBoundRegistry.Start(ctx); err != nil {
			return err
		}
	}
	if service.validatorRegistry != nil {
		if err := service.validatorRegistry.Start(ctx); err != nil {
			return err
		}
	}
	for _, monitor := range service.monitors {
		if err := monitor.Start(ctx); err != nil {
			return err
		}
	}
	go func() {
		for {
			select {
			case <-service.terminate:
				return
			default:
				if err := service.assetRegistry.FetchEvents(); err != nil {
					log.Printf("failed to fetch events of asset registry, %+v\n", err)
					continue
				}
				if service.assetUpperBoundRegistry != nil {
					if err := service.assetUpperBoundRegistry.FetchEvents(); err != nil {
						log.Printf("failed to fetch events of asset upper bound registry, %+v\n", err)
						continue
					}
				}
				if service.validatorRegistry != nil {
					if err := service.validatorRegistry.FetchEvents(); err != nil {
						log.Printf("failed to fetch events of validator registry, %+v\n", err)
						continue
					}
				}
				for _, monitor := range service.monitors {
					if err := monitor.Run(); err != nil {
						log.Printf("failed to monitor tube %d, %+v\n", monitor.TubeID(), err)
					}
					// TODO: better error handling
				}
			}
			time.Sleep(service.interval)
		}
	}()
	return nil
}

func (service *Service) Stop(ctx context.Context) error {
	service.terminate <- struct{}{}
	time.Sleep(service.interval * 2)
	for _, monitor := range service.monitors {
		if err := monitor.Stop(ctx); err != nil {
			return err
		}
	}
	if service.validatorRegistry != nil {
		if err := service.validatorRegistry.Stop(ctx); err != nil {
			return err
		}
	}
	if err := service.assetRegistry.Stop(ctx); err != nil {
		return err
	}
	service.chainClient.Close()

	db, err := service.db.DB()
	if err != nil {
		return err
	}
	return db.Close()
}
