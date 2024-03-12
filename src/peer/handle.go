package peer

import (
	"log"
)

func (peer Peer) Handle() {
	log.Println("Sarting to handle messagaes with map :", peer.handlers)
	for {
		msg, err := ReadMessage(peer)
		if err != nil {
			log.Println("Error handling new message :", err)
		}

		log.Printf("Handled new message on peer %v : %v\n", peer.RemoteAddr(), msg)

		handler, ok := peer.handlers[msg.GetCommand()]
		if !ok {
			log.Println("Received message with unknown command", msg.GetCommand())
			continue
		}

		go handler(msg)
	}
}
