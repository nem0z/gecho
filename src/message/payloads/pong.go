package payloads

type Pong struct {
	Nonce []byte
}

func NewPong(nonce []byte) *Pong {
	return &Pong{Nonce: nonce}
}

func ParsePong(data []byte) *Pong {
	return NewPong(data)
}

func (pong *Pong) ToByte() ([]byte, error) {
	return pong.Nonce, nil
}
