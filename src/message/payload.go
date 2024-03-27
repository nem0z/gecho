package message

import "github.com/nem0z/gecho/message/payloads"

type Payload interface {
	ToByte() ([]byte, error)
}

func ParsePayload(command string, data []byte) *Payload {
	var payload Payload

	switch command {
	case "echo":
		payload = payloads.ParseEcho(data)
	case "ping":
		payload = payloads.ParsePing(data)
	case "pong":
		payload = payloads.ParsePong(data)
	}

	return &payload
}
