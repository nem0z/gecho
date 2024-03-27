package message

import (
	"bytes"

	"github.com/nem0z/gecho/utils"
)

const HeaderLength = 24

type Message struct {
	header  *Header
	payload Payload
}

func Format(command string, data []byte) *Message {
	header := &Header{
		command:  command,
		length:   len(payloadData),
		checksum: utils.Checksum(payloadData),
}

func New(header *Header, data []byte) *Message {
	payload := NewPayload(header.command, data)
	return &Message{header, *payload}
}

func (msg *Message) GetCommand() string {
	return msg.header.command
}

func (msg *Message) IsValid() bool {
	command := IsValidCommand(msg.header.command)
	if !command {
		return false
	}

	payload, err := msg.payload.ToByte()
	if err != nil {
		return false
	}

	if msg.header.length != len(payload) {
		return false
	}

	return bytes.Equal(utils.Checksum(payload), msg.header.checksum)
}

func (msg *Message) Marshall() ([]byte, error) {
	header, err := msg.header.Marshall()
	if err != nil {
		return nil, err
	}

	payload, err := msg.payload.ToByte()
	if err != nil {
		return nil, err
	}

	return append(header, payload...), nil
}
