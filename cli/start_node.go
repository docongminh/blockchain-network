package cli

import (
	"fmt"
	"log"

	"github.com/docongminh/dapps/blockchain/core"
	"github.com/docongminh/dapps/blockchain/server"
)

func (cli *CLI) startNode(nodeID, minerAddress string) {
	fmt.Printf("Starting node %s\n", nodeID)
	if len(minerAddress) > 0 {
		if core.ValidateAddress(minerAddress) {
			fmt.Println("Mining is on. Address to receive rewards: ", minerAddress)
		} else {
			log.Panic("Wrong miner address!")
		}
	}
	server.StartServer(nodeID, minerAddress)
}
