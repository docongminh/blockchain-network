package core

import (
	"bytes"

	"github.com/docongminh/dapps/blockchain/utils"
)

// representation a transaction input, it's reference to previous output
type TXInput struct {
	Txid      []byte
	Vout      int
	Signature []byte
	PubKey    []byte
}

// Build struct representation a transaction output
// Outputs are indivisible -> when reference to new trans will spent whole value
// If greater amount, will generate and send back to sender (same real world)
type TXOutput struct {
	Value      int
	PubKeyHash []byte
}

// check  whether the address initiated the transaction
func (in *TXInput) UsesKey(pubKeyHash []byte) bool {
	lockingHash := HashPubKey(in.PubKey)

	return bytes.Compare(lockingHash, pubKeyHash) == 0
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
