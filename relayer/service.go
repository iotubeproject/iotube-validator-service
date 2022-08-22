package relayer

import (
	"context"
	"crypto/ecdsa"
	"log"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pkg/errors"
)

var errGasPriceTooHigh = errors.New("Gas price is too high")

type (
	Service struct {
		recorder  *Recorder
		interval  time.Duration
		relayers  []*Relayer
		terminate chan struct{}
	}
)

func NewService(
	cfgs []Config,
	databaseURL string,
	interval time.Duration,
) (*Service, error) {
	recorder := NewRecorder(databaseURL)
	relayers := make([]*Relayer, 0, len(cfgs))
	for _, cfg := range cfgs {
		privateKeys := []*ecdsa.PrivateKey{}
		for _, pk := range strings.Split(cfg.PrivateKeys, ",") {
			pk = strings.TrimSpace(pk)
			if len(pk) == 0 {
				continue
			}
			privateKey, err := crypto.HexToECDSA(pk)
			if err != nil {
				return nil, err
			}
			privateKeys = append(privateKeys, privateKey)
		}
		relayer, err := NewRelayer(cfg, recorder, privateKeys)
		if err != nil {
			return nil, err
		}
		relayers = append(relayers, relayer)
	}

	return &Service{
		recorder:  recorder,
		interval:  interval,
		terminate: make(chan struct{}),
		relayers:  relayers,
	}, nil
}

func (service *Service) Start(ctx context.Context) error {
	if err := service.recorder.Start(ctx); err != nil {
		return err
	}
	for _, relayer := range service.relayers {
		if err := relayer.Start(ctx); err != nil {
			return err
		}
	}
	go func() {
		for {
			select {
			case <-service.terminate:
				return
			default:
				for _, relayer := range service.relayers {
					if err := relayer.Run(); err != nil {
						log.Printf("failed to relay transfer to tube %d, %+v\n", relayer.TubeID(), err)
					}
				}
			}
			time.Sleep(service.interval)
		}
	}()

	return nil
}

func (service *Service) Stop(ctx context.Context) error {
	close(service.terminate)
	time.Sleep(service.interval * 2)
	for _, relayer := range service.relayers {
		if err := relayer.Stop(ctx); err != nil {
			return err
		}
	}
	return service.recorder.Stop(ctx)
}
