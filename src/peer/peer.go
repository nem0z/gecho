package peer

import (
	"net"
	"time"
)

type Peer struct {
	net.Conn
	handlers  map[string]Handler
	LastNonce []byte
	LastPing  time.Time
	LastPong  time.Time
}

func New(addr string) (*Peer, error) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return nil, err
	}

	p := &Peer{Conn: conn, handlers: map[string]Handler{}}
	go p.Handle()

	return p, p.Ping()
}

func NewFromConn(conn net.Conn) *Peer {
	p := &Peer{Conn: conn, handlers: map[string]Handler{}}
	go p.Handle()
	p.Ping()
	return p
}
