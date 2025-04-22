package main

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"encoding/json"
	token "ethclient-study/contract"
	"ethclient-study/eth"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/ethclient/simulated"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/miguelmota/go-ethereum-hdwallet"
	"golang.org/x/crypto/sha3"
	"log"
	"math"
	"math/big"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

const mainnet string = ""
const sepolia string = ""
const mainnetWS string = ""
const sepoliaWS string = ""
const privateKeySepoliaAccount1 = ""

func main() {
	//testCloudflareETH()
	//testAddress()
	//testGetBalance()
	//testCreateWallet()
	//testCreateKeyStore()
	//testImportKeyStore()
	//testCreateWalletByHDWallet()
	//testCheckAddress()
	//testCheckTransaction()
	//testTransferETH()
	//testSubscribeNewHead()
	//testRawTransaction()
	//testDeploySmartContract()
	//testLoadContract()
	//testContractWrite()
	//testContractRead()
	//testEventSubscript()
	//testContractWrite()
	//testEventRead()
	//testReadERC20()
	//testEventRead0xprotocol()
	//testSign()
	testClientSimulated()
}

func testClientSimulated() {
	// 1.创建私钥与账户地址
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		fmt.Println("1")
		log.Fatal(err)
	}
	fromAddress := crypto.PubkeyToAddress(privateKey.PublicKey)
	// 2.设置账户初始余额
	alloc := types.GenesisAlloc{
		fromAddress: {
			Balance: big.NewInt(1e18),
		},
	}
	// 3.设置模拟客户端
	sim := simulated.NewBackend(alloc, simulated.WithBlockGasLimit(8000000))

	defer func(sim *simulated.Backend) {
		err := sim.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(sim)

	// 4.创建交易签名
	chainID := big.NewInt(1337) // 模拟链ID
	nonce, err := sim.Client().PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		fmt.Println("Error getting nonce", err)
		log.Fatal(err)
	}

	// 6.生成接收地址
	toAddress := common.HexToAddress("0x1111111111111111111111111111111111111111")

	// 7.构建交易
	value := big.NewInt(1e17)
	gasPrice := big.NewInt(1)
	gasLimit := uint64(21000)
	txData := &types.LegacyTx{
		Nonce:    nonce,
		GasPrice: gasPrice,
		Gas:      gasLimit,
		To:       &toAddress,
		Value:    value,
		Data:     nil,
	}
	tx := types.NewTx(txData)

	// 8.签名交易
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		fmt.Println("2")
		log.Fatal(err)
	}

	// 9.广播交易
	err = sim.Client().SendTransaction(context.Background(), signedTx)
	if err != nil {
		fmt.Println("3")
		log.Fatal(err)
	}
	// 10.手动提交区块
	sim.Commit()

	// 11.查询交易回执
	//receipt, err := sim.Client().TransactionReceipt(context.Background(), signedTx.Hash())
	receipt, err := bind.WaitMined(context.Background(), sim.Client(), signedTx)
	if err != nil {
		fmt.Println("4")
		log.Fatal(err)
	}

	fmt.Println("Transaction successful:", receipt.Status == 1)
	fmt.Println("Gas used:", receipt.GasUsed)
	fmt.Println("Tx Hash", signedTx.Hash().Hex())
}

// 使用私对数据加密, 公钥验证
func testSign() {

	privateKey, err := crypto.HexToECDSA(privateKeySepoliaAccount1)
	if err != nil {
		log.Fatal(err)
	}

	// 加密内容hash
	data := []byte("hello world")
	hash := crypto.Keccak256Hash(data)
	fmt.Println(hash.Hex())
	// 签名
	signature, err := crypto.Sign(hash.Bytes(), privateKey)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(hexutil.Encode(signature))

	// 验证
	// 私钥生成的公钥
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}
	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	// 将 ecdsa.PublicKey 转换为原始字节表示（65 字节，前缀 0x04 + x + y）。

	// 从签名中恢复公钥字节
	sigPublicKey, err := crypto.Ecrecover(hash.Bytes(), signature)
	if err != nil {
		log.Fatal(err)
	}
	// 私钥+代加密数据->签名->crypto.Ecrecover签名公钥
	fmt.Println(hexutil.Encode(sigPublicKey))
	// 私钥->公钥
	fmt.Println(hexutil.Encode(publicKeyBytes))
	matches := bytes.Equal(sigPublicKey, publicKeyBytes)
	fmt.Println(matches)

	// 另一种方法
	// 签名恢复出ecdsa.PublicKey
	sigPublicKeyECDSA, err := crypto.SigToPub(hash.Bytes(), signature)
	if err != nil {
		log.Fatal(err)
	}
	sigPublicKeyBytes := crypto.FromECDSAPub(sigPublicKeyECDSA)
	matches = bytes.Equal(sigPublicKeyBytes, publicKeyBytes)
	fmt.Println(matches)

	// 第三种方法
	// VerifySignature的签名格式是64字节,只包含r和s,因此去掉v(恢复ID)
	signatureNoRecoverID := signature[:len(signature)-1]
	verified := crypto.VerifySignature(publicKeyBytes, hash.Bytes(), signatureNoRecoverID)
	fmt.Println(verified)

}

func testEventRead0xprotocol() {
	type LogFill struct {
		Maker                  common.Address
		Taker                  common.Address
		FeeRecipient           common.Address
		MakerToken             common.Address
		TakerToken             common.Address
		FilledMakerTokenAmount *big.Int
		FilledTakerTokenAmount *big.Int
		PaidMakerFee           *big.Int
		PaidTakerFee           *big.Int
		Tokens                 [32]byte
		OrderHash              [32]byte
	}

	type LogCancel struct {
		Maker                     common.Address
		FeeRecipient              common.Address
		MakerToken                common.Address
		TakerToken                common.Address
		CancelledMakerTokenAmount *big.Int
		CancelledTakerTokenAmount *big.Int
		Tokens                    [32]byte
		OrderHash                 [32]byte
	}

	type LogError struct {
		ErrorID   uint8
		OrderHash [32]byte
	}

	client, err := ethclient.Dial(mainnet)
	if err != nil {
		log.Fatal(err)
	}

	contractAddress := common.HexToAddress("")
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(0x3F70A1),
		ToBlock:   big.NewInt(0x3F8557),
		Addresses: []common.Address{contractAddress},
	}

	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}

	contractAbi, err := abi.JSON(strings.NewReader(token.ExchangeMetaData.ABI))
	if err != nil {
		log.Fatal(err)
	}

	logFillSig := []byte("LogFill(address,address,address,address,address,uint,uint,uint,uint,bytes32,bytes32)")
	logFillEvent := crypto.Keccak256Hash(logFillSig)
	logCancelSig := []byte("LogCancel(address,address,address,address,uint,uint,bytes32,bytes32)")
	logCancelEvent := crypto.Keccak256Hash(logCancelSig)
	logErrorSig := []byte("LogError(uint8,bytes32)")
	logErrorEvent := crypto.Keccak256Hash(logErrorSig)

	for _, vLog := range logs {
		fmt.Printf("Log Block Number: %d\n", vLog.BlockNumber)
		fmt.Printf("Log Index: %d\n", vLog.Index)
		switch vLog.Topics[0].Hex() {
		case logFillEvent.Hex():
			fmt.Println("Log Fill Event")
			var fillEvent LogFill
			err := contractAbi.UnpackIntoInterface(&fillEvent, "LogFill", vLog.Data)
			if err != nil {
				log.Fatal(err)
			}
			fillEvent.Maker = common.HexToAddress(vLog.Topics[1].Hex())
			fillEvent.FeeRecipient = common.HexToAddress(vLog.Topics[2].Hex())
			fillEvent.Tokens = common.HexToHash(vLog.Topics[3].Hex())

			fmt.Printf("Maker: %s\n", fillEvent.Maker.Hex())
			fmt.Printf("Taker: %s\n", fillEvent.Taker.Hex())
			fmt.Printf("FeeRecipient: %s\n", fillEvent.FeeRecipient.Hex())
			fmt.Printf("MakerToken: %s\n", fillEvent.MakerToken.Hex())
			fmt.Printf("TakerToken: %s\n", fillEvent.TakerToken.Hex())
			fmt.Printf("FilledMakerTokenAmount: %s\n", fillEvent.FilledMakerTokenAmount.String())
			fmt.Printf("FilledTakerTokenAmount: %s\n", fillEvent.FilledTakerTokenAmount.String())
			fmt.Printf("PaidMakerFee: %s\n", fillEvent.PaidMakerFee.String())
			fmt.Printf("PaidTakerFee: %s\n", fillEvent.PaidTakerFee.String())
			fmt.Printf("Tokens:%s\n", hexutil.Encode(fillEvent.Tokens[:]))
			fmt.Printf("OrderHash:%s\n", hexutil.Encode(fillEvent.OrderHash[:]))
		case logCancelEvent.Hex():
			fmt.Println("Log Cancel Event")
			var cancelEvent LogCancel
			err := contractAbi.UnpackIntoInterface(&cancelEvent, "LogCancel", vLog.Data)
			if err != nil {
				log.Fatal(err)
			}
			cancelEvent.Maker = common.HexToAddress(vLog.Topics[1].Hex())
			cancelEvent.FeeRecipient = common.HexToAddress(vLog.Topics[2].Hex())
			cancelEvent.Tokens = common.HexToHash(vLog.Topics[3].Hex())

			fmt.Printf("Maker: %s\n", cancelEvent.Maker.Hex())
			fmt.Printf("FeeRecipient:%s\n", cancelEvent.FeeRecipient.Hex())
			fmt.Printf("MakerToken: %s\n", cancelEvent.MakerToken.Hex())
			fmt.Printf("TakerToken: %s\n", cancelEvent.TakerToken.Hex())
			fmt.Printf("CancelledMakerTokenAmount: %s\n", cancelEvent.CancelledMakerTokenAmount.String())
			fmt.Printf("CancelledTakerTokenAmount: %s\n", cancelEvent.CancelledTakerTokenAmount.String())
			fmt.Printf("Tokens:%s\n", hexutil.Encode(cancelEvent.Tokens[:]))
			fmt.Printf("OrderHash:%s\n", hexutil.Encode(cancelEvent.OrderHash[:]))
		case logErrorEvent.Hex():
			fmt.Println("Log Fill Event")
			errorID, err := strconv.ParseInt(vLog.Topics[1].Hex(), 16, 64)
			if err != nil {
				log.Fatal(err)
			}
			errorEvent := &LogError{
				ErrorID:   uint8(errorID),
				OrderHash: vLog.Topics[2],
			}

			fmt.Printf("ErrirID: %d\n", errorEvent.ErrorID)
			fmt.Printf("OrderHash: %s\n", hexutil.Encode(errorEvent.OrderHash[:]))

		}
	}
}

func testReadERC20() {
	type LogTransfer struct {
		From   common.Address
		To     common.Address
		Tokens *big.Int
	}
	type LogApproval struct {
		TokenOwner common.Address
		Spender    common.Address
		Tokens     *big.Int
	}

	client, err := ethclient.Dial(mainnet)
	if err != nil {
		log.Fatal(err)
	}
	contractAddress := common.HexToAddress("")
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(0x112A888),
		ToBlock:   big.NewInt(0x1175AC1),
		Addresses: []common.Address{contractAddress},
	}

	// 用FilterLogs过滤日志
	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("logs:", len(logs))

	// 解析abi
	contractAbi, err := abi.JSON(strings.NewReader(token.Myerc20tokenMetaData.ABI))
	if err != nil {
		log.Fatal(err)
	}

	// 为了按某种日志类型过滤, 需要弄清楚每个时间日志函数签名的keccak256哈希值
	// 时间日志函数签名哈希始终是topic[0]
	logTransferSig := []byte("Transfer(address,address,uint256)")
	logApprovalSig := []byte("Approval(address,address,uint256)")
	logTransferSigHash := crypto.Keccak256Hash(logTransferSig)
	logApprovalSigHash := crypto.Keccak256Hash(logApprovalSig)

	// 遍历所有日志,设置switch语句按事件日志类型过滤
	for _, vLog := range logs {

		fmt.Printf("Log Block Number: %d\n", vLog.BlockNumber)
		fmt.Printf("Log Index: %d\n", vLog.Index)
		fmt.Printf(vLog.Topics[0].Hex())
		switch vLog.Topics[0].Hex() {
		case logTransferSigHash.Hex():
			fmt.Println("Log Transaction")
			var transferEvent LogTransfer
			err := contractAbi.UnpackIntoInterface(&transferEvent, "Transfer", vLog.Data)
			if err != nil {
				log.Fatal(err)
			}
			transferEvent.From = common.HexToAddress(vLog.Topics[1].Hex())
			transferEvent.To = common.HexToAddress(vLog.Topics[2].Hex())

			fmt.Printf("TransferEvent From: %s, To: %s\n", transferEvent.From, transferEvent.To)
			fmt.Printf("TransferEvent Tokens: %s\n", transferEvent.Tokens.String())

		case logApprovalSigHash.Hex():
			fmt.Println("Log Transaction")
			var approvalEvent LogApproval
			err := contractAbi.UnpackIntoInterface(&approvalEvent, "Approval", vLog.Data)
			if err != nil {
				log.Fatal(err)
			}
			approvalEvent.TokenOwner = common.HexToAddress(vLog.Topics[1].Hex())
			approvalEvent.Spender = common.HexToAddress(vLog.Topics[2].Hex())
			fmt.Printf("ApprovalEvent Token Owner: %s\n", approvalEvent.TokenOwner.Hex())
			fmt.Printf("ApprovalEvent Spender: %s\n", approvalEvent.Spender.Hex())
			fmt.Printf("ApprovalEvent Tokens: %s\n", approvalEvent.Tokens.String())

		}
	}

}
func testEventRead() {
	client, err := ethclient.Dial(sepolia)
	if err != nil {
		log.Fatal(err)
	}
	contractAddress := common.HexToAddress("")
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(1337),
		ToBlock:   big.NewInt(1337),
		Addresses: []common.Address{
			contractAddress,
		},
	}

	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}

	contractAbi, err := abi.JSON(strings.NewReader(string(token.StroeMetaData.ABI)))
	if err != nil {
		log.Fatal(err)
	}

	var topics [4]string

	for _, vLog := range logs {
		event := struct {
			Key   [32]byte
			Value [32]byte
		}{}
		err := contractAbi.UnpackIntoInterface(&event, "ItemSet", vLog.Data)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(event.Key[:]))   // foo
		fmt.Println(string(event.Value[:])) // bar
		fmt.Println(vLog.BlockHash.Hex())
		fmt.Println(vLog.BlockNumber)
		fmt.Println(vLog.TxHash.Hex())
		// 主题(索引)
		for i := range vLog.Topics {
			topics[i] = vLog.Topics[i].Hex()
		}
	}

	eventSignature := []byte("ItemSet(byte32, bytes32)")
	hash := crypto.Keccak256Hash(eventSignature)
	fmt.Println(hash.Hex())
}

// 订阅
func testEventSubscript() {
	client, err := ethclient.Dial(sepoliaWS)
	if err != nil {
		log.Fatal(err)
	}

	// 创建筛选查询
	contractAddress := common.HexToAddress("")
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}
	// 通过go channel接收事件
	logs := make(chan types.Log)

	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Subscribed")
	go func() {
		for {
			select {
			case err := <-sub.Err():
				log.Fatal(err)
			case vLog := <-logs:
				fmt.Println("接收到订阅")
				fmt.Println(vLog)
				return
			}
		}
	}()

}

func testContractRead() {
	client, err := ethclient.Dial(sepolia)
	if err != nil {
		log.Fatal(err)
	}
	contractAddress := common.HexToAddress("")
	bytecode, err := client.CodeAt(context.Background(), contractAddress, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(hex.EncodeToString(bytecode))
}

func testContractWrite() {
	client, err := ethclient.Dial(sepolia)
	if err != nil {
		log.Fatal(err)
	}
	privateKey, err := crypto.HexToECDSA(privateKeySepoliaAccount1)
	if err != nil {
		log.Fatal(err)
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatal(err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(300000)
	auth.GasPrice = gasPrice

	// 合约地址
	address := common.HexToAddress("")
	// 获取合约
	instance, err := token.NewStroe(address, client)
	if err != nil {
		log.Fatal(err)
	}
	key := [32]byte{}
	value := [32]byte{}
	copy(key[:], []byte("foo4"))
	copy(value[:], []byte("bar4"))

	tx, err := instance.SetItem(auth, key, value)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s\n", tx.Hash().Hex())

	result, err := instance.Items(nil, key)
	if err != nil {
		log.Fatal(err)
	}
	// cleanValue := bytes.Trim(result[:], "\x00")
	fmt.Printf(string(result[:]))

}

func testLoadContract() {
	client, err := ethclient.Dial(sepolia)
	if err != nil {
		log.Fatal(err)
	}
	address := common.HexToAddress("")
	instance, err := token.NewStroe(address, client)
	if err != nil {
		log.Fatal(err)
	}

	version, _ := instance.Version(nil)
	fmt.Println("Version: ", version)
}

func testDeploySmartContract() {
	client, err := ethclient.Dial(sepolia)
	if err != nil {
		fmt.Println("1")
		log.Fatal(err)
	}
	privateKey, err := crypto.HexToECDSA(privateKeySepoliaAccount1)
	if err != nil {
		fmt.Println("2")
		log.Fatal(err)
	}
	// 从私钥派生需要发送的账户的公共地址
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		fmt.Println("3")
		log.Fatal("error casting public key to ECDSA")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	// 获取随机数 nonce
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		fmt.Println("4")
		log.Fatal(err)
	}

	// gasPrice := big.NewInt(10000000000)
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		fmt.Println("5")
		log.Fatal(err)
	}

	// 获取当前以太坊主网ChainID
	chainId, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	if err != nil {
		log.Fatal(err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(300000)
	auth.GasPrice = gasPrice

	input := "1.0"
	address, tx, instance, err := token.DeployStroe(auth, client, input)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(address.Hex())
	fmt.Println(tx.Hash().Hex())
	fmt.Println(instance)

}

// 原始交易
func testRawTransaction() {
	client, err := ethclient.Dial(sepolia)
	if err != nil {
		fmt.Println("1")
		log.Fatal(err)
	}

	// 加载私钥
	privateKey, err := crypto.HexToECDSA(privateKeySepoliaAccount1)
	if err != nil {
		fmt.Println("2")
		log.Fatal(err)
	}

	// 从私钥派生需要发送的账户的公共地址
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		fmt.Println("3")
		log.Fatal("error casting public key to ECDSA")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	// 获取随机数 nonce
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		fmt.Println("4")
		log.Fatal(err)
	}

	// 设置需要转移的ETH数量
	value := big.NewInt(5000000000000000)
	// 设置gasLimit
	gasLimit := uint64(21000)
	// 设置gasPrice
	// gasPrice := big.NewInt(10000000000)
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		fmt.Println("5")
		log.Fatal(err)
	}
	// 发送给谁
	toAddress := common.HexToAddress("")
	// 生成事务
	txData := &types.LegacyTx{
		Nonce:    nonce,
		GasPrice: gasPrice,
		Gas:      gasLimit,
		To:       &toAddress,
		Value:    value,
		Data:     nil,
	}
	tx := types.NewTx(txData)

	// 用发件人私钥对事务签名
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		fmt.Println("6")
		log.Fatal(err)
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		fmt.Println("7")
		log.Fatal(err)
	}

	// 所有交易在发送到网络前，必须被编码成 RLP 格式，链上节点才能识别它为合法的交易格式。
	rawTxBytes, err := rlp.EncodeToBytes(signedTx)
	if err != nil {
		log.Fatal(err)
	}
	rawTxHex := hex.EncodeToString(rawTxBytes)
	fmt.Println("rawTxHex:", rawTxHex)

	// 使用curl 或 Postman广播
	//curl -X POST https://sepolia.infura.io/v3/YOUR_PROJECT_ID \
	//-H "Content-Type: application/json" \
	//-d '{"jsonrpc":"2.0","method":"eth_sendRawTransaction","params":["0xf86..."],"id":1}'

	tx = new(types.Transaction)
	rlp.DecodeBytes(rawTxBytes, &tx)

	err = client.SendTransaction(context.Background(), tx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("tx sent:%S", tx.Hash().Hex())
}

// 订阅新区块
func testSubscribeNewHead() {
	// 连接客户端
	client, err := ethclient.Dial(mainnetWS)
	if err != nil {
		fmt.Println("1")
		log.Fatal(err)
	}

	// 创建新通道用于接收新区块头
	headers := make(chan *types.Header)
	// 调用SubscribeNewHead
	sub, err := client.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		log.Fatal(err)
	}
	// 订阅
	count := 0
	// 超时通道
	timeout := time.After(60 * time.Second)
	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case header := <-headers:
			if header == nil {
				log.Fatal("nil header")
				continue
			}
			fmt.Println("接收到一个新区块")
			fmt.Println(header.Hash().Hex())
			// 获取完整区块
			block, err := client.BlockByHash(context.Background(), header.Hash())
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(block.Number())
			fmt.Println(block.Hash().Hex())
			fmt.Println(block.Time())
			fmt.Println(block.Nonce())
			fmt.Println(len(block.Transactions()))
			count++
			if count >= 2 {
				fmt.Println("已接收两个新区块，退出程序。")
				return
			}
		case <-timeout:
			fmt.Println("超时60s")
			return
		}

	}
}

// 转移代币,需要调用合约方法
func testTransferDaibi() {
	// 连接客户端
	client, err := ethclient.Dial(sepolia)
	if err != nil {
		fmt.Println("1")
		log.Fatal(err)
	}
	// 加载私钥
	privateKey, err := crypto.HexToECDSA(privateKeySepoliaAccount1)
	if err != nil {
		fmt.Println("2")
		log.Fatal(err)
	}

	// 从私钥派生需要发送的账户的公共地址
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		fmt.Println("3")
		log.Fatal("error casting public key to ECDSA")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	// 获取随机数 nonce
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		fmt.Println("4")
		log.Fatal(err)
	}

	// 设置需要转移的ETH数量, 代币传输不需要传输ETH
	value := big.NewInt(0)
	// 设置gasLimit, 使用预估gasLimit方法
	//gasLimit := uint64(21000)

	// 设置gasPrice
	// gasPrice := big.NewInt(10000000000)
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		fmt.Println("5")
		log.Fatal(err)
	}
	// 发送给谁
	toAddress := common.HexToAddress("")

	// 构建事务Data
	// 代币合约地址
	tokenAddress := common.HexToAddress("")
	transferFnSignature := []byte("transfer(address,uint256)")
	// 生成签名函数的Keccak256哈希
	hash := sha3.NewLegacyKeccak256()
	hash.Write(transferFnSignature)
	// 使用前4个字节获取方法ID
	methodId := hash.Sum(nil)[:4]
	fmt.Println(hexutil.Encode(methodId))
	// 左填充32字节
	paddedAddress := common.LeftPadBytes(toAddress.Bytes(), 32)
	fmt.Println(hexutil.Encode(paddedAddress))
	// 确定发送多少代币
	amount := new(big.Int)
	amount.SetString("1000000000000000000000", 10) //1000tokens
	// 代币量左填充到32字节
	paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)
	fmt.Println(hexutil.Encode(paddedAmount))
	// 方法ID, 填充后的地址, 转账量, 拼接到将成为我们数据字段的字节片
	var data []byte
	data = append(data, methodId...)
	data = append(data, paddedAddress...)
	data = append(data, paddedAmount...)
	// 预估gasLimit
	gasLimit, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
		To:   &toAddress,
		Data: data,
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(gasLimit)

	// 构建交易事务类型, To: 合约地址
	txData := &types.LegacyTx{
		Nonce:    nonce,
		GasPrice: gasPrice,
		Gas:      gasLimit,
		To:       &tokenAddress,
		Value:    value,
		Data:     data,
	}
	// 构建交易
	tx := types.NewTx(txData)

	// 发件人私钥签名
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}
	// 广播交易
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("tx sent: %s\n", signedTx.Hash().Hex())
}

func testTransferETH() {
	client, err := ethclient.Dial(sepolia)
	if err != nil {
		fmt.Println("1")
		log.Fatal(err)
	}

	// 加载私钥
	privateKey, err := crypto.HexToECDSA(privateKeySepoliaAccount1)
	if err != nil {
		fmt.Println("2")
		log.Fatal(err)
	}

	// 从私钥派生需要发送的账户的公共地址
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		fmt.Println("3")
		log.Fatal("error casting public key to ECDSA")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	// 获取随机数 nonce
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		fmt.Println("4")
		log.Fatal(err)
	}

	// 设置需要转移的ETH数量
	value := big.NewInt(5000000000000000)
	// 设置gasLimit
	gasLimit := uint64(21000)
	// 设置gasPrice
	// gasPrice := big.NewInt(10000000000)
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		fmt.Println("5")
		log.Fatal(err)
	}
	// 发送给谁
	toAddress := common.HexToAddress("0xf4AC31d5f2Fa732C039DC2BF5F93599082Cda371")

	// 生成事务
	txData := &types.LegacyTx{
		Nonce:    nonce,
		GasPrice: gasPrice,
		Gas:      gasLimit,
		To:       &toAddress,
		Value:    value,
		Data:     nil,
	}
	tx := types.NewTx(txData)

	// 用发件人私钥对事务签名
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		fmt.Println("6")
		log.Fatal(err)
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		fmt.Println("7")
		log.Fatal(err)
	}

	// 广播已签名事务
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		fmt.Println("8")
		log.Fatal(err)
	}
	fmt.Printf("tx sent: %s\n", signedTx.Hash().Hex())

}

func testCheckTransaction() {
	client, err := ethclient.Dial(mainnet)
	if err != nil {
		log.Fatal(err)
	}
	// 返回最新的区块头
	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("header: %v\n", header)

	// 根据区块编号查询完整区块
	blockNumber := big.NewInt(5671745)
	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("block: %v\n", block)
	fmt.Println(block.Number().Uint64())
	fmt.Println(block.Time())
	fmt.Println(block.Difficulty().Uint64())
	fmt.Println(block.Hash().Hex())
	fmt.Println(block.Transactions())
	fmt.Println("-------------")

	for idx, tx := range block.Transactions() {
		fmt.Println(strconv.FormatInt(int64(idx), 10) + "----------------------")
		// 交易哈希
		fmt.Println("交易Hash")
		fmt.Println(tx.Hash().Hex())
		// 交易金额(wei)
		fmt.Println("交易金额")
		fmt.Println(tx.Value().String())
		// Gas上限
		fmt.Println("Gas上限 GasLimit")
		fmt.Println(tx.Gas())
		// Gas单价
		fmt.Println("Gas单价")
		fmt.Println(tx.GasPrice().Uint64())
		// 防止重放攻击
		fmt.Println("防止重放攻击")
		fmt.Println(tx.Nonce())
		// 交易携带的数据
		fmt.Println("交易数据")
		fmt.Println(tx.Data())
		// 接收方地址
		fmt.Println("接收方地址:" + tx.To().Hex())

		// 获取发送方地址
		// 获取当前以太坊主网ChainID
		chainId, err := client.NetworkID(context.Background())
		if err != nil {
			log.Fatal(err)
		}
		// 构建EIP55签名器
		// 使用types.Sender 从签名中恢复发送方地址
		if sender, err := types.Sender(types.NewEIP155Signer(chainId), tx); err == nil {
			fmt.Println("sender", sender.Hex())
		}

		// 获取交易状态
		// 获取交易回执receipt
		receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
		if err != nil {
			log.Fatal(err)
		}
		data, _ := json.MarshalIndent(receipt, "", "  ")
		fmt.Printf("receipt: %v\n", string(data))
		fmt.Printf("receipt.status = %v\n", receipt.Status)
	}

	// 访问区块中的所有交易,查询某比交易详细信息是, 是否在等待确认
	// 目标区块哈希值
	blockHash := common.HexToHash("")

	// 查询有多少笔交易
	count, err := client.TransactionCount(context.Background(), blockHash)

	// 遍历
	fmt.Println("区块交易哈希")
	for idx := uint(0); idx < count; idx++ {
		tx, err := client.TransactionInBlock(context.Background(), blockHash, idx)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(tx.Hash().Hex())
	}

	// 查询单笔交易的详细信息
	fmt.Println("查询单笔交易")
	// 交易哈希
	txHash := common.HexToHash("")
	// 获取交易
	tx, isPending, err := client.TransactionByHash(context.Background(), txHash)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(tx.Hash().Hex())
	fmt.Println(isPending)
}

// 使用正则检查地址是否有效
func testCheckAddress() {
	// 检查格式
	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
	addressString := ""
	addressString2 := ""

	fmt.Println(re.MatchString(addressString))
	fmt.Println(re.MatchString(addressString2))

	// 检查是否为智能合约地址
	client, err := ethclient.Dial(sepolia)
	if err != nil {
		log.Fatal(err)
	}
	scAddressString := ""
	scAddress := common.HexToAddress(scAddressString)
	bytecode, err := client.CodeAt(context.Background(), scAddress, nil)
	if err != nil {
		log.Fatal(err)
	}
	isContract := len(bytecode) > 0
	fmt.Printf("is contract: %t\n", isContract)

}

// 根据助记词生成多个钱包地址
func testCreateWalletByHDWallet() {
	wallet, err := hdwallet.NewFromMnemonic("aim ankle wolf staff pond soap emotion off turkey donkey elephant slender")
	if err != nil {
		log.Fatal(err)
	}
	path := hdwallet.MustParseDerivationPath("m/44'/60'/0'/0/0")
	account, err := wallet.Derive(path, false)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(account.Address.Hex())

	path = hdwallet.MustParseDerivationPath("m/44'/60'/0'/0/3")
	account, err = wallet.Derive(path, false)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(account.Address.Hex())
}

// 从已有keystore文件导入钱包地址
func testImportKeyStore() {
	// 钱包文件路径
	file := "./wallets/UTC--2025-04-19T05-13-32.244325000Z--866252d134837862ab5d94853e7f059332398073"
	// 新建钱包文件
	ks := keystore.NewKeyStore("./tmp", keystore.StandardScryptN, keystore.StandardScryptP)
	// 读取钱包文件内容
	jsonBytes, err := os.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	// 解密这个钱包所用的密码, 必须和当初创建时一致
	passward := "123456"
	// 导入钱包账户
	account, err := ks.Import(jsonBytes, passward, passward)
	if err != nil {
		log.Fatal(err)
	}
	// 打印导入的钱包地址
	fmt.Println(account.Address.Hex())
	// 删除原始钱包文件(因为它已被导入keystore并复制了一份), 防止重复导入
	if err := os.Remove(file); err != nil {
		log.Fatal(err)
	}

}

// 创建新的以太坊钱包地址和私钥
func testCreateKeyStore() {

	// 创建密钥库, 钱包文件保存到./tmp
	// keystore.StandardScryptN, keystore.StandardScryptP 加密参数,用于保护私钥
	ks := keystore.NewKeyStore("./wallets", keystore.StandardScryptN, keystore.StandardScryptP)
	// 设置加密密码,之后需要使用这个密码解锁钱包或交易签名
	password := "123456"
	// 是用密码创建新的以太坊账户,生成对应的私钥保存在./tmp中
	account, err := ks.NewAccount(password)
	if err != nil {
		log.Fatal(err)
	}
	// 钱包地址
	fmt.Println("account:", account.Address.Hex())
}

// 生成钱包
func testCreateWallet() {

	// 1.生成私钥
	// 2.生成私钥
	// 3.生成钱包地址

	// 生成私钥结构体
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}

	// 转成字节切片,方便后续编码/存储
	privateKeyBytes := crypto.FromECDSA(privateKey)
	// 转成16进制字符窜,去掉前缀0x ==> 私钥
	fmt.Println(hexutil.Encode(privateKeyBytes)[2:])

	// 从私钥中导出公钥
	publicKey := privateKey.Public()

	// 类型断言,转成*ecdsa.PublicKey类型公钥结构体
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	// 转成字节切片
	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	// 转成16进制字符串, 去掉前缀4个字符 0x04
	fmt.Println(hexutil.Encode(publicKeyBytes)[4:])
	// 根据公钥生成以太坊地址
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	// 打印钱包地址
	fmt.Println(address)

	// 创建Keccak256哈希计算器
	hash := sha3.NewLegacyKeccak256()
	// 公钥去掉第一个字节,输入到哈希计算器(第一个字节通常是0x04,表示为压缩公钥格式,不参与计算)
	hash.Write(publicKeyBytes[1:])
	// 计算出Keccak256 取后20字节(也就是以太坊地址).并用16进制输出,结果同address
	fmt.Println(hexutil.Encode(hash.Sum(nil)[12:]))
}

func testERC20() {
	client, err := ethclient.Dial(sepolia)
	if err != nil {
		log.Fatal(err)
	}

	tokenAddress := common.HexToAddress("")
	instance, err := token.NewToken(tokenAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	callOpts := &bind.CallOpts{
		Context: context.Background(),
	}

	address := common.HexToAddress("")
	balance, err := instance.BalanceOf(callOpts, address)
	if err != nil {
		log.Fatal(err)
	}

	name, err := instance.Name(callOpts)
	if err != nil {
		log.Fatal(err)
	}

	symbol, err := instance.Symbol(callOpts)
	if err != nil {
		log.Fatal(err)
	}

	decimals, err := instance.Decimals(callOpts)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("name: %s\n", name)
	fmt.Printf("symbol: %s\n", symbol)
	fmt.Printf("decimals: %v\n", decimals)

	fmt.Printf("wei: %s\n", balance)
	fbal := new(big.Float)
	fbal.SetString(balance.String())
	value := new(big.Float).Quo(fbal, big.NewFloat(math.Pow10(int(decimals))))
	fmt.Printf("balance: %f\n", value)

}

func testGetBalance() {
	addressString := ""
	eth.TestBalance(addressString)
}

func testAddress() {
	addressString := ""
	address := eth.ConvertToAddress(addressString)
	fmt.Println(address)
	fmt.Println(address.Hex())
	fmt.Println(address.Bytes())
	hash := crypto.Keccak256Hash(address.Bytes())
	fmt.Println(hash.Hex())

}

func testLocalETH() {
	client, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to Ethereum network")
	_ = client
}

func testCloudflareETH() {
	client, err := ethclient.Dial("https://cloudflare-eth.com")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("we have a connection")
	_ = client // we'll use this in the upcoming sections
}
