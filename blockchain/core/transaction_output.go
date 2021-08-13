package core

import (
	"bytes"

	"github.com/docongminh/dapps/blockchain/utils"
)

// Build struct representation a transaction output
type TXOutput struct {
	Value      int
	PubKeyHash []byte
}

// signs the output
func (out *TXOutput) Lock(address []byte) {
	pubKeyHash := utils.Base58Decode(address)
	pubKeyHash = pubKeyHash[1 : len(pubKeyHash)-4]
	out.PubKeyHash = pubKeyHash
}

// checks if the output can be used by the owner of the publey
func (out *TXOutput) IsLockedWithKey(pubKeyHash []byte) bool {
	return bytes.Compare(out.PubKeyHash, pubKeyHash) == 0
}

// create new TXOutput
func NewTXOutput(value int, address string) *TXOutput {
	txo := &TXOutput{value, nil}
	txo.Lock([]byte(address))

	return txo
}
