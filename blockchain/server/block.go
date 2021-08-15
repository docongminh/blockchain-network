package server

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"

	"github.com/docongminh/dapps/blockchain/core"
)

type Block struct {
	AddrFrom string
	Block    []byte
}

type GetBlocks struct {
	AddrFrom string
}

func SendBlock(addr string, b *core.Block) {
	data := Block{nodeAddress, b.SerializerBlock()}
	payload := gobEncode(data)
	request := append(commandToBytes("block"), payload...)

	SendData(addr, request)
}
func RequestBlocks() {
	for _, node := range KnownNodes {
		SendGetBlocks(node)
	}
}
func HandleBlock(request []byte, bc *core.Blockchain) {
	var buff bytes.Buffer
	var payload Block

	buff.Write(request[commandLength:])
	dec := gob.NewDecoder(&buff)
	err := dec.Decode(&payload)
	if err != nil {
		log.Panic(err)
	}

	blockData := payload.Block
	block := core.DeserializerBlock(blockData)

	fmt.Println("Recevied a new block!")
	bc.AddBlock(block)

	fmt.Printf("Added block %x\n", block.Hash)

	if len(blocksInTransit) > 0 {
		blockHash := blocksInTransit[0]
		SendGetData(payload.AddrFrom, "block", blockHash)

		blocksInTransit = blocksInTransit[1:]
	} else {
		UTXOSet := core.UTXOSet{bc}
		UTXOSet.Reindex()
	}
}
func HandleGetBlocks(request []byte, bc *core.Blockchain) {
	var buff bytes.Buffer
	var payload GetBlocks

	buff.Write(request[commandLength:])
	dec := gob.NewDecoder(&buff)
	err := dec.Decode(&payload)
	if err != nil {
		log.Panic(err)
	}

	blocks := bc.GetBlockHashes()
	SendInv(payload.AddrFrom, "block", blocks)
}
func SendGetBlocks(address string) {
	payload := gobEncode(GetBlocks{nodeAddress})
	request := append(commandToBytes("getblocks"), payload...)

	SendData(address, request)
}
