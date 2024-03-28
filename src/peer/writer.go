package peer

import (
	"fmt"

	"github.com/nem0z/gecho/message"
)

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
