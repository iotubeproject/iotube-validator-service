package registry

import (
	"context"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"net/http"

	"github.com/crosschained-io/crosschained-service/api"
	"github.com/crosschained-io/crosschained-service/contract"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	ErrNotRegistered    = errors.New("not registered")
	ErrNotFound         = errors.New("not found")
	ErrUnknownGenre     = errors.New("unknown genre")
	ErrInvalidSignature = errors.New("invalid signature")
)

type (
	FetchValidatorFile func(common.Address) (uint64, string, error)
	TestimonyDAO       struct {
		client             *ethclient.Client
		fetchValidatorFile FetchValidatorFile
	}
)

func NewTestimonyDAO(client *ethclient.Client, fetchValidatorFile FetchValidatorFile) *TestimonyDAO {
	return &TestimonyDAO{
		client:             client,
		fetchValidatorFile: fetchValidatorFile,
	}
}

func sigToAddress(id []byte, signature []byte) (common.Address, error) {
	pubkey, err := crypto.SigToPub(id, signature)
	if err != nil {
		return common.Address{}, err
	}
	return crypto.PubkeyToAddress(*pubkey), nil
}

func checkSignature(id []byte, signature []byte, validator common.Address) error {
	addr, err := sigToAddress(id, signature)
	if err != nil {
		return err
	}
	if addr != validator {
		return ErrInvalidSignature
	}
	return nil
}

func (dao *TestimonyDAO) Testimony(validator common.Address, id common.Hash) ([]byte, error) {
	genre, uri, err := dao.fetchValidatorFile(validator)
	if err != nil {
		return nil, err
	}
	var signature []byte
	switch genre {
	default:
		return nil, ErrUnknownGenre
	case NotRegistered:
		return nil, ErrNotRegistered
	case CONTRACT:
		bs, err := hex.DecodeString(uri)
		if err != nil {
			return nil, err
		}
		filter, err := contract.NewTestimonyDAOFilterer(common.BytesToAddress(bs), dao.client)
		if err != nil {
			return nil, err
		}
		iter, err := filter.FilterTestimony(&bind.FilterOpts{}, []common.Address{validator}, [][32]byte{id})
		if err != nil {
			return nil, err
		}
		if !iter.Next() {
			return nil, ErrNotFound
		}
		signature = iter.Event.Testimony
	case GRPC_API:
		conn, err := grpc.Dial(uri, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			return nil, err
		}

		defer conn.Close()
		remote := api.NewValidatorServiceClient(conn)
		response, err := remote.QueryByID(context.Background(), &api.QueryRequest{Id: id.Bytes()})
		if err != nil {
			return nil, err
		}

		signature = response.GetSignature()
	case HTTP:
		response, err := http.Get(uri + "/" + base64.URLEncoding.EncodeToString(id.Bytes()))
		if err != nil {
			return nil, err
		}
		if response.StatusCode != 200 {
			return nil, errors.Errorf("failed to query signature, with response status %d", response.StatusCode)
		}
		defer response.Body.Close()
		var value struct {
			Validator string `json:"validator"`
			Signature string `json:"signature"`
		}
		if err := json.NewDecoder(response.Body).Decode(&value); err != nil {
			return nil, err
		}
		signature, err = base64.StdEncoding.DecodeString(value.Signature)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to decode signature %s", value.Signature)
		}
	}
	if err := checkSignature(id.Bytes(), signature, validator); err != nil {
		return nil, err
	}

	return signature, nil
}
