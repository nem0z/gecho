package handlers

import (
	"fmt"
	"log"

	"github.com/nem0z/gecho/message"
	"github.com/nem0z/gecho/message/payloads"
	"github.com/nem0z/gecho/peer"
)

func Ping(peer *peer.Peer) peer.Handler {
	return func(msg *message.Message) {
		fmt.Println(msg.GetCommand(), msg.IsValid())
		ping := (*msg.GetPayload()).(*payloads.Ping)

		pong := payloads.NewPong(ping.Nonce)
		pongMsg, err := message.Format("pong", pong)
		if err != nil {
			log.Println("Ping handler :", err)
			return
		}

		fmt.Println("Sending pong :", ping.Nonce)
		err = peer.Send(pongMsg)
		if err != nil {
			log.Println("Ping handler :", err)
		}
	}
}
