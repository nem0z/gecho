package peer

import (
	"fmt"
	"io"
	"net"
	"time"

	"github.com/nem0z/gecho/message"
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

func (peer *Peer) Register(command string, handler Handler) {
	peer.handlers[command] = handler
}

func (p *Peer) Send(msg *message.Message) error {
	m, err := msg.Marshall()
	if err != nil {
		return err
	}

	n, err := p.Write(m)
	if err != nil {
		return err
	}

	if n < len(m) {
		return fmt.Errorf("not enough bytes written")
	}

	return nil
}

func (p *Peer) ReadAll() ([]byte, error) {
	const bufSize = 1024
	msg := []byte{}
	buf := make([]byte, bufSize)

	for {
		n, err := p.Read(buf)

		if err != nil && err != io.EOF {
			return nil, err
		}

		msg = append(msg, buf[:n]...)
		if n < bufSize || err != nil {
			break
		}
	}

	return msg, nil
}
