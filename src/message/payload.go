package message

import "github.com/nem0z/gecho/message/payloads"

type Payload interface {
	ToByte() ([]byte, error)
}

func NewPayload(command string, data []byte) *Payload {
	var payload Payload

	switch command {
	case "echo":
		payload = messages.NewEcho(data)
	}

	return &payload
}
