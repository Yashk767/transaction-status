package utils

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"math/big"
	"os"
	"path"
	"razor/structs"
	"strings"
)

func GetOptions() bind.CallOpts {
	block, _ := new(big.Int).SetString("", 10)
	return bind.CallOpts{
		Pending:     false,
		BlockNumber: block,
		Context:     context.Background(),
	}
}

func GetTxnOpts(transactionData structs.TransactionOptions) *bind.TransactOpts {
	log.Debug("Getting transaction options...")
	defaultPath, err := os.Getwd()
	CheckError("Error in fetching parent path: ", err)
	keystorePath := path.Join(defaultPath, "test_accounts")
	privateKey, err := GetPrivateKey(transactionData.AccountAddress, transactionData.Password, keystorePath)
	if privateKey == nil || err != nil {
		CheckError("Error in fetching private key: ", errors.New(transactionData.AccountAddress+" not present in razor-go"))
	}
	nonce, err := transactionData.Client.NonceAt(context.Background(), common.HexToAddress(transactionData.AccountAddress), nil)
	CheckError("Error in fetching nonce: ", err)

	txnOpts, err := bind.NewKeyedTransactorWithChainID(privateKey, ChainId)
	CheckError("Error in getting transactor: ", err)
	txnOpts.Nonce = big.NewInt(int64(nonce))
	gasPrice, err := transactionData.Client.SuggestGasPrice(context.Background())
	CheckError("Error in getting suggested gas price: ", err)
	txnOpts.GasPrice = gasPrice
	txnOpts.Value = transactionData.EtherValue
	gasLimit, err := GetGasLimit(transactionData, txnOpts)
	CheckError("Error in getting gasLimit: ", err)
	txnOpts.GasLimit = gasLimit
	//txnOpts.GasLimit = 50000000
	return txnOpts
}

func GetPrivateKey(address string, password string, keystorePath string) (*ecdsa.PrivateKey, error) {
	ks := keystore.NewKeyStore(keystorePath, keystore.StandardScryptN, keystore.StandardScryptP)
	allAccounts := ks.Accounts()
	for _, account := range allAccounts {
		if strings.EqualFold(account.Address.Hex(), address) {
			return GetPrivateKeyFromKeystore(account.URL.Path, password)
		}
	}
	return nil, errors.New("no keystore file found")
}

func GetPrivateKeyFromKeystore(keystorePath string, password string) (*ecdsa.PrivateKey, error) {
	jsonBytes, err := os.ReadFile(keystorePath)
	CheckError("Error in reading keystore: ", err)

	key, err := keystore.DecryptKey(jsonBytes, password)
	CheckError("Error in fetching private key: ", err)

	return key.PrivateKey, nil
}

func CheckTransactionReceipt(client *ethclient.Client, _txHash string) int {
	txHash := common.HexToHash(_txHash)
	tx, err := client.TransactionReceipt(context.Background(), txHash)
	if err != nil {
		return -1
	}
	return int(tx.Status)
}

func CheckError(msg string, err error) {
	if err != nil {
		panic(msg + err.Error())
	}
}

func GetGasLimit(transactionData structs.TransactionOptions, txnOpts *bind.TransactOpts) (uint64, error) {
	if transactionData.MethodName == "" {
		return 0, nil
	}
	parsedData, err := abi.JSON(strings.NewReader(transactionData.ABI))
	if err != nil {
		log.Error("Error in parsing abi: ", err)
		return 0, err
	}
	inputData, err := parsedData.Pack(transactionData.MethodName, transactionData.Parameters...)
	if err != nil {
		log.Error("Error in calculating inputData: ", err)
		return 0, err
	}
	contractAddress := common.HexToAddress(transactionData.ContractAddress)
	msg := ethereum.CallMsg{
		From:     common.HexToAddress(transactionData.AccountAddress),
		To:       &contractAddress,
		GasPrice: txnOpts.GasPrice,
		Value:    txnOpts.Value,
		Data:     inputData,
	}
	gasLimit, err := transactionData.Client.EstimateGas(context.Background(), msg)
	if err != nil {
		log.Error("error in estimating gas: ", err)
		return 0, err
	}
	fmt.Println("Estimated Gas: ", gasLimit)
	return gasLimit, nil
}
