package peer

import (
	"fmt"
	"io"
	"net"

	"github.com/nem0z/gecho/message"
)

type Peer struct {
	net.Conn
	handlers map[string]Handler
}

func New(addr string) (*Peer, error) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return nil, err
	}

	p := &Peer{conn, map[string]Handler{}}
	go p.Handle()

	return p, nil
}

func NewFromConn(conn net.Conn) *Peer {
	p := &Peer{conn, map[string]Handler{}}
	go p.Handle()
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
