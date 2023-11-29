package main

import (
	"context"
	"fmt"

	"github.com/0xPolygon/cdk-validium-node/encoding"
	"github.com/0xPolygon/cdk-validium-node/log"
	ERC20 "github.com/0xPolygon/cdk-validium-node/test/contracts/bin/ERC20"
	WETH "github.com/0xPolygon/cdk-validium-node/test/contracts/bin/WETH"
	"github.com/0xPolygon/cdk-validium-node/test/contracts/bin/uniswap/v2/core/UniswapV2Factory"
	"github.com/0xPolygon/cdk-validium-node/test/contracts/bin/uniswap/v2/core/UniswapV2Pair"
	"github.com/0xPolygon/cdk-validium-node/test/contracts/bin/uniswap/v2/interface/UniswapInterfaceMulticall"
	"github.com/0xPolygon/cdk-validium-node/test/contracts/bin/uniswap/v2/periphery/UniswapV2Router02"
	"github.com/0xPolygon/cdk-validium-node/test/operations"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	// if you want to test using goerli network
	// replace this by your goerli infura url
	//networkURL = "http://localhost:8123"
	networkURL = "http://localhost:8545"
	// replace this by your account private key
	//pk = "0xdfd01798f92667dbf91df722434e8fbe96af0211d4d1b82bbbbc8f1def7a814f"
	pk = operations.DefaultSequencerPrivateKey
)

func main() {
	ctx := context.Background()
	log.Infof("connecting to %v", networkURL)
	client, err := ethclient.Dial(networkURL)
	uniswap.ChkErr(err)
	log.Infof("connected")
	chainID, err := client.ChainID(ctx)
	uniswap.ChkErr(err)
	log.Infof("chainID: %v", chainID)
	auth := uniswap.GetAuth(ctx, client, pk)
	fmt.Println()
	deployments := uniswap.DeployContractsAndAddLiquidity(client, auth)
	for i := 0; i < 5; i++ {
		uniswap.SwapTokens(client, auth, deployments)
	}

}
