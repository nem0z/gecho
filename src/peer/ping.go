package peer

import (
	"time"

	"github.com/nem0z/gecho/message"
	"github.com/nem0z/gecho/message/payloads"
	"github.com/nem0z/gecho/utils"
)

func (peer *Peer) Ping() error {
	nonce, err := utils.Nonce(8)
	if err != nil {
		return err
	}

	ping := payloads.NewPing(nonce)
	msgPing, err := message.Format("ping", ping)
	if err != nil {
		return err
	}

	peer.Send(msgPing)
	peer.LastNonce = nonce
	peer.LastPing = time.Now()
	return nil
}
