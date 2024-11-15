package service

import (
	"Lotso_Airdrop_Server/contracts/airdrop"
	"Lotso_Airdrop_Server/model/base"
	"Lotso_Airdrop_Server/utils/flags"
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/labstack/gommon/log"
	"math/big"
)

func ClaimAirdrop(airdropCount uint64, contractAddress common.Address, proofs [][32]byte, privateKey *ecdsa.PrivateKey) (response *base.Response) {
	amount := scaleTokens(new(big.Int).SetUint64(airdropCount), flags.Decimals)

	client, err := ethclient.Dial(flags.ApiUrl)
	if err != nil {
		log.Error(err)
		return
	}
	defer client.Close()

	ctx := context.Background()

	gasPrice, err := client.SuggestGasPrice(ctx)
	if err != nil {
		log.Error(err)
		return
	}

	transactOpts, _ := bind.NewKeyedTransactorWithChainID(privateKey, new(big.Int).SetInt64(flags.ChainID))
	transactOpts.Value = big.NewInt(0) // in wei
	transactOpts.GasPrice = gasPrice
	transactOpts.GasLimit = 2000000

	airdropContract, err := airdrop.NewContracts(contractAddress, client)
	if err != nil {
		log.Error(err)
		return
	}

	transaction, err := airdropContract.Claim(transactOpts, amount, proofs)
	if err != nil {
		log.Error(err)
		response = base.NewErrorResponse(err, base.ClaimAirdropsFailed)
		return
	}
	fmt.Println(transaction.Hash().Hex())
	receipt, err := client.TransactionReceipt(ctx, transaction.Hash())

	if err != nil {
		log.Error(err)
		response = base.NewErrorResponse(err, base.ClaimAirdropsFailed)
		return
	}

	if receipt.Status == 0 {
		err = fmt.Errorf("transaction status failed, hash: %s", transaction.Hash().Hex())
		response = base.NewErrorResponse(err, base.ClaimAirdropsFailed)
		return
	}

	response = base.NewDataResponse(transaction.Hash())
	return
}
