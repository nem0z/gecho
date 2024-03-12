package peer

import "github.com/nem0z/gecho/message"

type Handler func(msg *message.Message)
