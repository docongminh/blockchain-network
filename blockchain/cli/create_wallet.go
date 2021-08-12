package cli

import (
	"fmt"

	"github.com/docongminh/dapps/blockchain/core"
)

func (cli *CLI) createWallet() {
	wallets, _ := core.NewWallets()
	address := wallets.CreateWallet()
	wallets.SaveToFile()

	fmt.Printf("Your new address: %s\n", address)
}
