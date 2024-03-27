package server

import (
	"github.com/nem0z/gecho/peer"
	"github.com/nem0z/gecho/peer/handlers"
)

type HandlerFactory func(peer *peer.Peer) peer.Handler

func (s *server) RegiterDefaultHandlers() {
	s.Register("echo", handlers.Echo)
	s.Register("ping", handlers.Ping)
	s.Register("pong", handlers.Pong)
}
