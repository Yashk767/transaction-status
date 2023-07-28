package contract_interaction

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"razor/bindings"
	"razor/utils"
)

func GetGetSetContract(client *ethclient.Client) *bindings.GetSetContract {
	getSetContract, err := bindings.NewGetSetContract(common.HexToAddress(utils.ContractAddress), client)
	if err != nil {
		panic(err)
	}
	return getSetContract
}
