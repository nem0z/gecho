package peer

import (
	"fmt"
	"io"

	"github.com/nem0z/gecho/message"
)

const bufferSize = 1024

func readHeader(peer Peer) (*message.Header, error) {
	header := make([]byte, message.HeaderLength)
	_, err := peer.Read(header)
	if err != nil {
		return nil, err
	}

	return message.ParseHeader(header)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func ReadMessage(peer Peer) (*message.Message, error) {
	header, err := readHeader(peer)
	if err != nil {
		return nil, err
	}

	payload := []byte{}
	buf := make([]byte, min(bufferSize, header.Len()))

	for {
		n, err := peer.Read(buf)

		if err != nil && err != io.EOF {
			return nil, err
		}

		payload = append(payload, buf[:n]...)
		if n < bufferSize || err != nil {
			break
		}
	}

	if len(payload) != header.Len() {
		return nil, fmt.Errorf("expected %v, read %v", header.Len(), len(payload))
	}

	return message.New(header, message.ParsePayload(header.GetCommand(), payload)), nil
}
