package service

import (
	"Lotso_Airdrop_Server/contracts"
	"Lotso_Airdrop_Server/database/mysql"
	"Lotso_Airdrop_Server/model"
	"Lotso_Airdrop_Server/utils/flags"
	"context"
	"crypto/ecdsa"
	"errors"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/labstack/gommon/log"
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

func LotsoDistributeAirdrops(addresses *[]model.AirdropItem) (hash common.Hash, err error) {
	client, err := ethclient.Dial(flags.ApiUrl)
	if err != nil {
		return
	}
	defer client.Close()

	privateKey, err := crypto.HexToECDSA(flags.PrivateKey)
	if err != nil {
		return
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		err = errors.New("error casting public key to ECDSA")
		return
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return
	}

	calculateGasLimit := func(addressCount int) (gasLimit uint64) {
		gasLimit = 25000*uint64(addressCount) + 35000 // 23012 32376
		return
	}

	transactOpts, _ := bind.NewKeyedTransactorWithChainID(privateKey, new(big.Int).SetInt64(flags.ChainID))
	transactOpts.Nonce = big.NewInt(int64(nonce))
	transactOpts.Value = big.NewInt(0)                         // in wei
	transactOpts.GasLimit = calculateGasLimit(len(*addresses)) // in units
	transactOpts.GasPrice = gasPrice

	//创建合约对象
	lotsoAirdrop, err := contracts.NewContracts(flags.AirdropContractAddress, client)
	if err != nil {
		return
	}

	recipients := make([]common.Address, len(*addresses))
	for i := range *addresses {
		recipients[i] = common.HexToAddress((*addresses)[i].Address)
	}

	amounts := make([]*big.Int, len(*addresses))
	for i := range amounts {
		amounts[i] = calAirdropAmount((*addresses)[i].AirdropCount - (*addresses)[i].HasAirdroppedCount)
	}

	tx, err := lotsoAirdrop.SetAirdrop(transactOpts, recipients, amounts)
	if err != nil {
		return
	}

	hash = tx.Hash()
	log.Info("SetAirdrop tx sent, hash: ", hash)
	return
}

func LotsoClaimAirdrops(privateKey *ecdsa.PrivateKey) (hash common.Hash, err error) {
	client, err := ethclient.Dial(flags.ApiUrl)
	if err != nil {
		return
	}
	defer client.Close()

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		err = errors.New("error casting public key to ECDSA")
		return
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return
	}

	transactOpts, _ := bind.NewKeyedTransactorWithChainID(privateKey, new(big.Int).SetInt64(flags.ChainID))
	transactOpts.Nonce = big.NewInt(int64(nonce))
	transactOpts.Value = big.NewInt(0) // in wei
	transactOpts.GasLimit = 100000     // in units
	transactOpts.GasPrice = gasPrice

	//创建合约对象
	lotsoAirdrop, err := contracts.NewContracts(flags.AirdropContractAddress, client)
	if err != nil {
		return
	}

	tx, err := lotsoAirdrop.ClaimAirdrop(transactOpts)
	if err != nil {
		return
	}

	hash = tx.Hash()
	log.Info("ClaimAirdrop tx sent, hash: ", hash)

	return
}

func LotsoRecipientsCount() (recipientsCount *big.Int, err error) {
	client, err := ethclient.Dial(flags.ApiUrl)
	if err != nil {
		return
	}
	defer client.Close()

	//创建合约对象
	lotsoAirdrop, err := contracts.NewContracts(flags.AirdropContractAddress, client)
	if err != nil {
		return
	}

	recipientsCount, err = lotsoAirdrop.RecipientsCount(nil)
	return
}

// amount = AirdropCount * 10**Decimals
func calAirdropAmount(airdropCount uint64) (amount *big.Int) {
	amount = new(big.Int).Mul(new(big.Int).SetUint64(airdropCount), new(big.Int).Exp(new(big.Int).SetInt64(10), new(big.Int).SetUint64(flags.Decimals), nil))
	return
}

func IsBuyer(address common.Address) (ok bool, err error) {
	return mysql.TopicExist(flags.PairAddress, address.Hex()[2:])
}
