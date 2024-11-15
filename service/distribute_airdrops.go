package service

import (
	"Lotso_Airdrop_Server/contracts/airdrop"
	"Lotso_Airdrop_Server/contracts/token"
	"Lotso_Airdrop_Server/database"
	"Lotso_Airdrop_Server/model/base"
	"Lotso_Airdrop_Server/utils/flags"
	"context"
	"fmt"
	smt "github.com/FantasyJony/openzeppelin-merkle-tree-go/standard_merkle_tree"
	mingfugorm "github.com/Fu-XDU/mingfu_go_common/gorm"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/labstack/gommon/log"
	"math/big"
)

type DistributeAirdropsResponse struct {
	AccountCount    uint   `json:"accountCount"`
	ContractAddress string `json:"contractAddress"`
}

func DistributeAirdrops() (response *base.Response) {
	airdropItems, err := database.GetAddressesShouldAirdrop()
	if err != nil {
		return base.NewErrorResponse(err, base.GetAirdropItemFailed)
	}

	if len(*airdropItems) == 0 {
		response = base.NewResponse()
		response.Message = "no airdrops need to be distributed"
		return
	}

	airdropValues := make([][]interface{}, len(*airdropItems))
	for i, airdropItem := range *airdropItems {
		airdropValues[i] = []interface{}{
			common.HexToAddress(airdropItem.Address.Address),
			scaleTokens(new(big.Int).SetUint64(airdropItem.AirdropCount), flags.Decimals),
		}
	}

	leafEncodings := []string{
		smt.SOL_ADDRESS,
		smt.SOL_UINT256,
	}

	tree, err := smt.CreateTree(leafEncodings)
	if err != nil {
		return
	}

	for _, v := range airdropValues {
		_, err = tree.AddLeaf(v)
		if err != nil {
			return
		}
	}

	root, err := tree.MakeTree()
	fmt.Println("MakeTree Root: ", hexutil.Encode(root))

	address, err := deployAirdropContract(root)
	if err != nil {
		fmt.Println("deployAirdropContract failed:", err)
		return
	}

	addressHex := address.Hex()[2:]
	for i := range *airdropItems {
		(*airdropItems)[i].ContractAddress = addressHex

		var proof [][]byte
		proof, err = tree.GetProof(airdropValues[i])
		if err != nil {
			return
		}

		(*airdropItems)[i].Proofs = make(mingfugorm.StringArray, len(proof))

		for j := range proof {
			(*airdropItems)[i].Proofs[j] = common.Bytes2Hex(proof[j])
		}
	}

	err = database.UpdateAirdrops(airdropItems)
	if err != nil {
		fmt.Println("Error UpdateAirdrops :", err)
		return
	}

	response = base.NewDataResponse(&DistributeAirdropsResponse{
		AccountCount:    uint(len(*airdropItems)),
		ContractAddress: "0x" + addressHex,
	})

	return
}

func deployAirdropContract(merkleRoot []byte) (address common.Address, err error) {
	client, err := ethclient.Dial(flags.ApiUrl)
	if err != nil {
		return
	}
	defer client.Close()

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return
	}

	transactOpts, _ := bind.NewKeyedTransactorWithChainID(flags.PrivateKey, new(big.Int).SetInt64(flags.ChainID))
	transactOpts.Value = big.NewInt(0) // in wei
	transactOpts.GasPrice = gasPrice

	address, transaction, _, err := airdrop.DeployContracts(transactOpts, client, [32]byte(merkleRoot), flags.TokenAddress)
	if err != nil {
		return
	}

	log.Infof("Depoly airdrop contract success, address: %s, transaction hash: %v", address.Hex(), transaction.Hash())

	// 调用Token合约来approve
	tokenContract, err := token.NewContracts(flags.TokenAddress, client)
	if err != nil {
		return
	}

	transactOpts, _ = bind.NewKeyedTransactorWithChainID(flags.PrivateKey, new(big.Int).SetInt64(flags.ChainID))
	transactOpts.Value = big.NewInt(0) // in wei
	transactOpts.GasPrice = gasPrice

	// MAX uint64
	amount, _ := new(big.Int).SetString("ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff", 16)

	transaction, err = tokenContract.Approve(transactOpts, address, amount)
	if err != nil {
		return
	}

	log.Infof("Approve airdrop contract success, address: %s, transaction hash: %v", address.Hex(), transaction.Hash())
	return
}

func scaleTokens(tokens *big.Int, decimals uint) *big.Int {
	multiplier := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(decimals)), nil)
	tokens.Mul(tokens, multiplier)
	return tokens
}
