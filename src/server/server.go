package server

import (
	"fmt"
	"log"
	"net"
	"sync"

	"github.com/nem0z/gecho/peer"
)

type server struct {
	listener net.Listener
	peers    []*peer.Peer
	handlers map[string]HandlerFactory
	mu       sync.Mutex
}

func New(port int) (*server, error) {
	ln, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		return nil, err
	}

	serv := &server{listener: ln, peers: []*peer.Peer{}, handlers: map[string]HandlerFactory{}}
	serv.RegiterDefaultHandlers()
	go serv.accept()

	return serv, nil
}

func (s *server) Register(key string, handler HandlerFactory) {
	s.handlers[key] = handler
}

func (s *server) RegisterPeerHandlers(peer *peer.Peer) {
	for command, handler := range s.handlers {
		peer.Register(command, handler(peer))
	}
}

func (s *server) accept() {
	log.Println("Waiting for incoming connections")
	for {
		conn, err := s.listener.Accept()
		if err != nil {
			log.Println("Error accepting new conn :", err)
			continue
		}

		log.Printf("Handled new connection %v\n", conn.RemoteAddr())
		p := peer.NewFromConn(conn)
		s.RegisterPeerHandlers(p)
		s.addPeer(p)
	}
}

func (s *server) addPeer(peer *peer.Peer) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.peers = append(s.peers, peer)
}

func (s *server) Close() {
	for _, peer := range s.peers {
		peer.Close()
	}
	s.listener.Close()
}
