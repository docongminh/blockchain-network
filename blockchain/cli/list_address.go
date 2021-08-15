package cli

import (
	"fmt"
	"log"

	"github.com/docongminh/dapps/blockchain/core"
)

func (cli *CLI) listAddresses(nodeID string) {
	wallets, err := core.NewWallets(nodeID)
	if err != nil {
		log.Panic(err)
	}
	addresses := wallets.GetAddresses()

	for _, address := range addresses {
		fmt.Println(address)
	}
}
