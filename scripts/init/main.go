package main

import (
	"fmt"
	"math/big"
	"os"

	"github.com/conflux-fans/storage-cli/config"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"

	providers "github.com/openweb3/go-rpc-provider/provider_wrapper"
	"github.com/openweb3/web3go"
	"github.com/openweb3/web3go/signers"
	"github.com/openweb3/web3go/types"
)

func main() {
	initConfig()
	distributeEth()
}

func initConfig() {
	config.SetConfigFile("../../config.yaml")
	config.Init()
}

func distributeEth() {
	cfg := config.Get()

	privateKeys := append(cfg.PrivateKeys, "9a6d3ba2b0c7514b16a006ee605055d71b9edfad183aeb2d9790e9d4ccced471")
	client := web3go.MustNewClientWithOption(cfg.BlockChain.URL, web3go.ClientOption{
		SignerManager: signers.MustNewSignerManagerByPrivateKeyStrings(privateKeys),
		Option: providers.Option{
			Logger: os.Stdout,
		},
	})

	sm, _ := client.GetSignerManager()
	from := common.HexToAddress("0x0e768D12395C8ABFDEdF7b1aEB0Dd1D27d5E2A7F")
	value := big.NewInt(1).Mul(big.NewInt(100), big.NewInt(1e18))
	for _, s := range sm.List() {
		addr := s.Address()
		txHash, err := client.Eth.SendTransactionByArgs(types.TransactionArgs{
			From:  &from,
			To:    &addr,
			Value: (*hexutil.Big)(value),
		})
		if err != nil {
			panic(fmt.Sprintf("发送交易失败: %s", err))
		}
		fmt.Printf("交易已发送: %s\n", txHash)
	}
}