package utils

import (
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"strings"
)

// 参数是前端传入的
func VerifySignature(message, signature, address string) (bool, error) {
	prefixedMessage := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(signature), signature)
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
