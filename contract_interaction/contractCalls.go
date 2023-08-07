package contract_interaction

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"razor/bindings"
	"razor/structs"
	"razor/utils"
)

func Reveal(txnData structs.TransactionOptions, epoch uint32, tree bindings.RevealMerkleTree, signature []byte) (common.Hash, error) {
	fmt.Println("Sending Reveal transaction...")
	txnData.ContractAddress = utils.RevealMockAddress
	txnData.MethodName = "reveal"
	txnData.ABI = bindings.RevealMockMetaData.ABI
	txnData.Parameters = []interface{}{epoch, tree, signature}
	txnOpts := utils.GetTxnOpts(txnData)

	revealMockContract := GetRevealMockContract(txnData.Client)
	txn, err := revealMockContract.Reveal(txnOpts, epoch, tree, signature)
	if err != nil {
		return common.Hash{0x00}, err
	}
	fmt.Println("Reveal txn hash: ", txn.Hash())
	return txn.Hash(), nil
}
