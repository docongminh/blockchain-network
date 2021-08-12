package cli

import (
	"fmt"
	"log"

	"github.com/docongminh/dapps/blockchain/core"
	"github.com/docongminh/dapps/blockchain/utils"
)

func (cli *CLI) getBalance(address string) {
	if !core.ValidateAddress(address) {
		log.Panic("ERROR: Address is not valid")
	}
	bc := core.NewBlockchain()
	UTXOSet := core.UTXOSet{bc}
	defer bc.DB.Close()

	balance := 0
	pubKeyHash := utils.Base58Decode([]byte(address))
	pubKeyHash = pubKeyHash[1 : len(pubKeyHash)-4]
	UTXOs := UTXOSet.FindUTXO(pubKeyHash)

	for _, out := range UTXOs {
		balance += out.Value
	}

	fmt.Printf("Balance of '%s': %d\n", address, balance)
}
