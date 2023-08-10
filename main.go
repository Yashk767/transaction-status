package main

import (
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"math/big"
	"os"
	"path"
	"razor/contract_interaction"
	"razor/structs"
	"razor/utils"
	"time"
)

func main() {
	address := "0xbd3e0a1d11163934df10501c9e1a18fbaa9ecaf4"
	password := "Test@123"
	provider := "https://staging-v3.skalenodes.com/v1/staging-aware-chief-gianfar"

	client, err := ethclient.Dial(provider)
	utils.CheckError("Error in dialing: ", err)

	txnData := structs.TransactionOptions{
		Client:         client,
		AccountAddress: address,
		Password:       password,
	}

	account := utils.Account{
		Address:  address,
		Password: password,
	}
	epoch := uint32(5638336)

	log.Debug("Getting transaction options...")
	defaultPath, err := os.Getwd()
	utils.CheckError("Error in fetching parent path: ", err)
	keystorePath := path.Join(defaultPath, "test_accounts")

	signature, _, err := utils.CalculateSecret(account, epoch, keystorePath, utils.ChainId)
	utils.CheckError("Error in calculating secret: ", err)

	fmt.Println("SIGNATURE: ", signature)

	commitData := utils.CommitData{
		AssignedCollections:    map[int]bool{0: true},
		SeqAllottedCollections: []*big.Int{big.NewInt(0), big.NewInt(0), big.NewInt(0)},
		Leaves:                 []*big.Int{big.NewInt(184431)},
	}

	merkleTree, err := utils.CreateMerkle(commitData.Leaves)
	utils.CheckError("Error in getting merkle tree: ", err)

	log.Debug("Generating tree reveal data...")
	treeRevealData := utils.GenerateTreeRevealData(merkleTree, commitData)

	fmt.Println("TREE REVEAL DATA: ", treeRevealData)
	fmt.Println("SIGNATURE: ", signature)

	revealTxn, err := contract_interaction.Reveal(txnData, epoch, treeRevealData, signature)
	utils.CheckError("Reveal Txn error: ", err)

	WaitForBlockCompletion(client, revealTxn.String())

}

func WaitForBlockCompletion(client *ethclient.Client, hashToRead string) {
	timeout := 30
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
}
