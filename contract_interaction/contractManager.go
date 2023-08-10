package contract_interaction

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"razor/bindings"
	"razor/utils"
)

func GetRevealMockContract(client *ethclient.Client) *bindings.RevealMock {
	revealMock, err := bindings.NewRevealMock(common.HexToAddress(utils.RevealMockAddress), client)
	if err != nil {
		panic(err)
	}
	return revealMock
}
