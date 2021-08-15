package cli

import (
	"fmt"
	"log"

	"github.com/docongminh/dapps/blockchain/core"
)

func (cli *CLI) createBlockchain(address, nodeID string) {
	if !core.ValidateAddress(address) {
		log.Panic("ERROR: Address is not valid")
	}
	bc := core.CreateBlockchain(address, nodeID)
	defer bc.DB.Close()

	UTXOSet := core.UTXOSet{bc}
	UTXOSet.Reindex()

	fmt.Println("Done!")
}
