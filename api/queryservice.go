package api

import (
	"context"
	"encoding/hex"
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"

	grpc "google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/pkg/errors"
)

type (
	APIService struct {
		UnimplementedValidatorServiceServer
		httpPort      int
		port          int
		validatorAddr common.Address
		validateAddr  bool
		recorder      *Recorder
		grpcServer    *grpc.Server
		httpServer    *http.Server
	}
)

func NewAPIService(
	port int,
	httpPort int,
	databaseURL string,
	validatorAddr common.Address,
) (*APIService, error) {
	recorder, err := NewRecorder(databaseURL)
	if err != nil {
		return nil, err
	}
	zeroAddr := common.Address{}

	return &APIService{
		port:          port,
		httpPort:      httpPort,
		recorder:      recorder,
		validatorAddr: validatorAddr,
		validateAddr:  validatorAddr.Hex() != zeroAddr.Hex(),
	}, nil
}

func (s *APIService) Start(ctx context.Context) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", s.port))
	if err != nil {
		return errors.Wrap(err, "failed to create listener")
	}
	s.grpcServer = grpc.NewServer()
	RegisterValidatorServiceServer(s.grpcServer, s)
	reflection.Register(s.grpcServer)
	go func() {
		if err := s.grpcServer.Serve(lis); err != nil {
			log.Fatal("failed to start grpc server", err)
		}
	}()
	gwmux := runtime.NewServeMux()
	if err := RegisterValidatorServiceHandlerServer(ctx, gwmux, s); err != nil {
		return err
	}
	if err := gwmux.HandlePath("GET", "/health", func(w http.ResponseWriter, _ *http.Request, _ map[string]string) {
		log.Println("health check")
		w.Write([]byte("active"))
	}); err != nil {
		log.Fatal("failed to add health check", err)
	}
	s.httpServer = &http.Server{
		Addr:    fmt.Sprintf(":%d", s.httpPort),
		Handler: gwmux,
	}

	return s.httpServer.ListenAndServe()
}

func (s *APIService) Stop(ctx context.Context) error {
	if err := s.httpServer.Shutdown(ctx); err != nil {
		return err
	}
	s.grpcServer.Stop()

	return nil
}

func (s *APIService) QueryByID(ctx context.Context, request *QueryRequest) (*Testimony, error) {
	log.Printf("Receive query request for %s\n", request.Id)
	id := common.HexToHash(strings.TrimPrefix(request.Id, "0x"))
	signature, err := s.recorder.Signature(id)
	if err != nil {
		return nil, err
	}
	pubkey, err := crypto.SigToPub(id[:], signature)
	if err != nil {
		return nil, err
	}
	validator := crypto.PubkeyToAddress(*pubkey)
	if s.validateAddr && validator != s.validatorAddr {
		return nil, errors.Errorf("invalid validator")
	}
	return &Testimony{Validator: validator.Hex(), Signature: hex.EncodeToString(signature)}, nil
}
