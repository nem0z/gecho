package peer

import "github.com/nem0z/gecho/message"

type Handler func(msg *message.Message)

func (peer *Peer) Register(command string, handler Handler) {
	peer.handlers[command] = handler
}
