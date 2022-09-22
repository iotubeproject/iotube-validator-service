package main

import (
	"fmt"
	"net/http"

	"github.com/crosschained-io/crosschained-service/registry"
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"
)

var httpCmd = &cobra.Command{
	Use:   "http",
	Short: "fetch the signature via http",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 1 {
			response, err := http.Get(args[0])
			if err != nil {
				return err
			}
			cmd.Printf("Response: %s\n", response.Status)
			return nil
		}
		signature, err := registry.FetchSignatureViaHttp(args[0], common.HexToHash(args[1]))
		if err != nil {
			return err
		}
		cmd.Printf("Signature: %x\n", signature)
		return nil
	},
}

var grpcCmd = &cobra.Command{
	Use:  "grpc",
	Args: cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		signature, err := registry.FetchSignatureViaGRPC(args[0], common.HexToHash(args[1]))
		if err != nil {
			return err
		}
		cmd.Printf("Signature: %x\n", signature)
		return nil
	},
}

func NewCommand() *cobra.Command {
	var cmd = &cobra.Command{}
	cmd.AddCommand(httpCmd)
	cmd.AddCommand(grpcCmd)
	return cmd
}

func main() {
	rootCmd := NewCommand()
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
