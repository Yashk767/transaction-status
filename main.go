package main

import (
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"razor/contract_interaction"
	"razor/structs"
	"razor/utils"
	"sync"
	"time"
)

func main() {
	address := "0xbd3e0a1d11163934df10501c9e1a18fbaa9ecaf4"
	password := "Test@123"
	provider := "https://staging-v3.skalenodes.com/v1/staging-aware-chief-gianfar"

	client, err := ethclient.Dial(provider)
	utils.CheckError("Error in dialing: ", err)

	valueBefore := contract_interaction.GetValue(client)
	fmt.Println("Value Before setting: ", valueBefore)
	txnData := structs.TransactionOptions{
		Client:         client,
		AccountAddress: address,
		Password:       password,
	}

	value := big.NewInt(1).Add(valueBefore, big.NewInt(1))

	setTxn, err := contract_interaction.SetValue(txnData, value)
	utils.CheckError("Set Transaction error: ", err)

	WaitForBlockCompletion(client, setTxn.String())

	valueAfter := contract_interaction.GetValue(client)
	fmt.Println("Value After setting: ", valueAfter)

}

func WaitForBlockCompletion(client *ethclient.Client, hashToRead string) {
	timeout := 30
	value := contract_interaction.GetValue(client)

	var wg sync.WaitGroup
	wg.Add(2)
	defer wg.Wait()
	go func() {
		for start := time.Now(); time.Since(start) < time.Duration(timeout)*time.Second; {
			fmt.Println("Checking if transaction is mined....")
			transactionStatus := utils.CheckTransactionReceipt(client, hashToRead)
			if transactionStatus == 0 {
				fmt.Println("TIME TO GET TRANSACTION STATUS: ", time.Since(start))
				err := errors.New("transaction mining unsuccessful")
				utils.CheckError("Error", err)
				wg.Done()
				return
			} else if transactionStatus == 1 {
				fmt.Println("TIME TO GET TRANSACTION STATUS: ", time.Since(start))
				fmt.Println("Transaction mined successfully")
				wg.Done()
				return
			}
			time.Sleep(3 * time.Second)
		}
	}()
	go func() {
		for start := time.Now(); time.Since(start) < time.Duration(timeout)*time.Second; {
			updatedValue := contract_interaction.GetValue(client)
			fmt.Println("Checking state change, value ", updatedValue)
			if value.Cmp(updatedValue) != 0 {
				fmt.Println("Value has been updated, state has changed")
				fmt.Println("TIME TO GET STATE CHANGE: ", time.Since(start))
				wg.Done()
				return
			}
			time.Sleep(3 * time.Second)
		}
	}()
}
