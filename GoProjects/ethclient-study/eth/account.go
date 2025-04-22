package eth

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math"
	"math/big"
)

func ConvertToAddress(addressString string) common.Address {
	address := common.HexToAddress(addressString)
	return address
}

func TestBalance(addressString string) {
	client, err := ethclient.Dial("https://sepolia.infura.io/v3/28d8f32dde1d480f91f2e3cecf91ca3e")
	//client, err := ethclient.Dial("https://cloudflare-eth.com")
	if err != nil {
		log.Fatal(err)
	}
	account := ConvertToAddress(addressString)
	blockNumber := big.NewInt(5532993)
	balanceAt, err := client.BalanceAt(context.Background(), account, blockNumber)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("balance", balanceAt)

	fbalance := new(big.Float)
	fbalance.SetString(balanceAt.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
	fmt.Println("ethValue", ethValue)

	pendingBalance, err := client.PendingBalanceAt(context.Background(), account)
	fmt.Println("pendingBalance", pendingBalance)
}
