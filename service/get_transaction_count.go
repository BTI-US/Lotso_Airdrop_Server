package service

import (
	"Lotso_Airdrop_Server/utils/flags"
	"errors"
	"github.com/ethereum/go-ethereum/rpc"
	"math/big"
	"strings"
)

// EthGetTransactionCount Returns the number of transactions sent from an address.
// https://ethereum.org/en/developers/docs/apis/json-rpc/#eth_gettransactioncount
func EthGetTransactionCount(address string) (transactionCount uint64, err error) {
	client, err := rpc.DialHTTP(flags.ApiUrl)
	if err != nil {
		return
	}
	defer client.Close()

	var result string

	if err = client.Call(&result, "eth_getTransactionCount", address, flags.CutoffBlock); err != nil {
		return
	}

	transactionCountHex := strings.TrimPrefix(result, "0x")

	var transactionCountBigInt big.Int
	_, success := transactionCountBigInt.SetString(transactionCountHex, 16)
	if !success {
		err = errors.New("failed to convert hex string to big.Int")
		return
	}
	transactionCount = transactionCountBigInt.Uint64()
	return
}
