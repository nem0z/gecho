package handlers

import (
	"fmt"

	"github.com/nem0z/gecho/message"
	"github.com/nem0z/gecho/peer"
)

func Echo(peer *peer.Peer) peer.Handler {
	return func(msg *message.Message) {
		fmt.Println("Received message echo :", msg)
	}
}
