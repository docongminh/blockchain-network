package server

import (
	"bytes"
	"encoding/gob"
	"log"

	"github.com/docongminh/core"
)

type Verzion struct {
	Version    int
	BestHeight int
	AddrFrom   string
}

func SendVersion(addr string, bc *core.Blockchain) {
	bestHeight := bc.GetBestHeight()
	payload := gobEncode(Verzion{nodeVersion, bestHeight, nodeAddress})

	request := append(commandToBytes("version"), payload...)

	SendData(addr, request)
}

func handleVersion(request []byte, bc *core.Blockchain) {
	var buff bytes.Buffer
	var payload Verzion

	buff.Write(request[commandLength:])
	dec := gob.NewDecoder(&buff)
	err := dec.Decode(&payload)
	if err != nil {
		log.Panic(err)
	}

	myBestHeight := bc.GetBestHeight()
	foreignerBestHeight := payload.BestHeight

	if myBestHeight < foreignerBestHeight {
		SendGetBlocks(payload.AddrFrom)
	} else if myBestHeight > foreignerBestHeight {
		SendVersion(payload.AddrFrom, bc)
	}

	// sendAddr(payload.AddrFrom)
	if !nodeIsKnown(payload.AddrFrom) {
		KnownNodes = append(KnownNodes, payload.AddrFrom)
	}
}
