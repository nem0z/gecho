package handlers

import (
	"bytes"
	"fmt"
	"time"

	"github.com/nem0z/gecho/message"
	"github.com/nem0z/gecho/message/payloads"
	"github.com/nem0z/gecho/peer"
)

func Pong(peer *peer.Peer) peer.Handler {
	return func(msg *message.Message) {
		pong := (*msg.GetPayload()).(*payloads.Pong)
		fmt.Println("Received pong :", pong.Nonce)

		if bytes.Equal(peer.LastNonce, pong.Nonce) {
			peer.LastPong = time.Now()
		}
	}
}
