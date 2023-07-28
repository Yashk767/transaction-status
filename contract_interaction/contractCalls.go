package contract_interaction

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"razor/bindings"
	"razor/structs"
	"razor/utils"
)

func SetValue(txnData structs.TransactionOptions, value *big.Int) (common.Hash, error) {
	fmt.Println("Sending Set transaction...")
	txnData.ContractAddress = utils.ContractAddress
	txnData.MethodName = "set"
	txnData.ABI = bindings.GetSetContractMetaData.ABI
	txnData.Parameters = []interface{}{value}
	txnOpts := utils.GetTxnOpts(txnData)

	getSetContract := GetGetSetContract(txnData.Client)
	txn, err := getSetContract.Set(txnOpts, value)
	if err != nil {
		return common.Hash{0x00}, err
	}
	fmt.Println("Set Txn Hash: ", txn.Hash())
	return txn.Hash(), nil
}

func GetValue(client *ethclient.Client) *big.Int {
	getSetContract := GetGetSetContract(client)
	callOpts := utils.GetOptions()
	value, err := getSetContract.Get(&callOpts)
	utils.CheckError("Error in getting value: ", err)
	return value
}
