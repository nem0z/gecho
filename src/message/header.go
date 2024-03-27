package message

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type Header struct {
	command  string
	length   int
	checksum []byte
}

func (h *Header) Len() int {
	return h.length
}

func (h *Header) GetCommand() string {
	return h.command
}

func (h *Header) Marshall() ([]byte, error) {
	command, err := FormatCommand(h.command)
	if err != nil {
		return nil, err
	}

	length := make([]byte, 8)
	binary.BigEndian.PutUint64(length, uint64(h.length))

	return bytes.Join([][]byte{command, length, h.checksum[:]}, nil), nil
}

func ParseHeader(data []byte) (*Header, error) {
	if len(data) != HeaderLength {
		return nil, fmt.Errorf("data's length doesn't match with header's length")
	}

	command, err := ParseCommand(data[:12])
	if err != nil {
		return nil, err
	}

	length := binary.BigEndian.Uint64(data[12:20])

	return &Header{
		command:  command,
		length:   int(length),
		checksum: data[20:24],
	}, nil
}
