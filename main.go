package main

import (
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"razor/bindings"
	"razor/contract_interaction"
	"razor/structs"
	"razor/utils"
	"time"
)

func main() {
	address := "0xbd3e0a1d11163934df10501c9e1a18fbaa9ecaf4"
	password := "Test@123"
	provider := "https://rpc-mumbai.maticvigil.com"

	client, err := ethclient.Dial(provider)
	utils.CheckError("Error in dialing: ", err)

	txnData := structs.TransactionOptions{
		Client:         client,
		AccountAddress: address,
		Password:       password,
	}

	epoch := uint32(5637935)
	revealMerkleTree := bindings.RevealMerkleTree{
		Values: []bindings.RevealAssignedAsset{{LeafId: 0, Value: big.NewInt(183545)}, {LeafId: 0, Value: big.NewInt(183545)}, {LeafId: 0, Value: big.NewInt(183545)}},
		Proofs: [][][32]byte{{{}}, {{}}, {{}}},
		Root:   [32]byte{154, 188, 11, 170, 156, 94, 233, 187, 24, 136, 148, 155, 128, 83, 235, 177, 203, 59, 105, 207, 72, 22, 252, 231, 135, 196, 138, 181, 81, 79, 227, 171},
	}

	signature := []byte{19, 157, 154, 240, 37, 198, 165, 111, 212, 167, 4, 153, 236, 179, 141, 78, 140, 23, 164, 89, 44, 53, 34, 224, 53, 112, 25, 193, 53, 235, 113, 246, 11, 210, 84, 30, 128, 247, 129, 73, 230, 45, 74, 66, 159, 65, 67, 91, 194, 180, 47, 8, 24, 239, 97, 246, 143, 13, 233, 11, 7, 61, 228, 239, 27}

	revealtxn, err := contract_interaction.Reveal(txnData, epoch, revealMerkleTree, signature)
	utils.CheckError("Reveal Txn error: ", err)

	WaitForBlockCompletion(client, revealtxn.String())

}

func WaitForBlockCompletion(client *ethclient.Client, hashToRead string) {
	timeout := 30
	go func() {
		for start := time.Now(); time.Since(start) < time.Duration(timeout)*time.Second; {
			fmt.Println("Checking if transaction is mined....")
			transactionStatus := utils.CheckTransactionReceipt(client, hashToRead)
			if transactionStatus == 0 {
				fmt.Println("TIME TO GET TRANSACTION STATUS: ", time.Since(start))
				err := errors.New("transaction mining unsuccessful")
				utils.CheckError("Error: ", err)
				return
			} else if transactionStatus == 1 {
				fmt.Println("TIME TO GET TRANSACTION STATUS: ", time.Since(start))
				fmt.Println("Transaction mined successfully")
				return
			}
			time.Sleep(3 * time.Second)
		}
	}()
}
