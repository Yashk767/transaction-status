package structs

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

type TransactionOptions struct {
	Client          *ethclient.Client
	Password        string
	EtherValue      *big.Int
	Amount          *big.Int
	AccountAddress  string
	ChainId         *big.Int
	ContractAddress string
	MethodName      string
	Parameters      []interface{}
	ABI             string
}
