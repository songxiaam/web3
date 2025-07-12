package utils

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"strings"
)

// 参数是前端传入的
func VerifySignature(message, signature, address string) (bool, error) {
	prefixedMessage := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(message), message)
	//message的Hash
	messageHash := crypto.Keccak256Hash([]byte(prefixedMessage))
	// 解码签名
	sig := common.FromHex(signature)
	if len(sig) != 65 {
		return false, errors.New("invalid signature length")
	}
	// 恢复公钥
	if sig[64] >= 27 {
		sig[64] -= 27
	}
	pubKey, err := crypto.SigToPub(messageHash.Bytes(), sig)
	if err != nil {
		return false, err
	}
	// 从公钥获取地址
	recoveredAddress := crypto.PubkeyToAddress(*pubKey)
	// 比较地址
	return strings.EqualFold(address, recoveredAddress.String()), nil
}

func Trade(to string) (uint64, error) {
	client, err := ethclient.Dial("https://mainnet.infura.io/v3/你的API_KEY")
	if err != nil {
		return 0, err
	}
	defer client.Close()

	// 私钥
	privateKeyHex := "你的私钥16进制（不要带0x）"
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return 0, err
	}

	// 公钥
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return 0, err
	}

	// 钱包地址
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	// 获取nonce
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return 0, err
	}

	// 构造交易
	value := big.NewInt(10000000000000000)
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return 0, err
	}
	gasLimit := uint64(21000)
	toAddress := common.HexToAddress(to)

	txData := types.LegacyTx{
		Nonce:    nonce,      // nonce of sender account
		GasPrice: gasPrice,   // wei per gas
		Gas:      gasLimit,   // gas limit
		To:       &toAddress, // nil means contract creation
		Value:    value,      // wei amount
		Data:     nil,        // contract invocation input data
	}

	tx := types.NewTx(&txData)

	// 签名交易
	chainID := big.NewInt(1) // 模拟链ID
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		return 0, err
	}
	// 广播交易
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return 0, err
	}

	receipt, err := bind.WaitMined(context.Background(), client, signedTx)
	if err != nil {
		return 0, err
	}
	fmt.Println("Transaction successful:", receipt.Status == 1)
	fmt.Println("Gas used:", receipt.GasUsed)
	fmt.Println("Tx Hash", signedTx.Hash().Hex())
	return 0, nil
}
