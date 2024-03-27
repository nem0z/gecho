package payloads

type Ping struct {
	Nonce []byte
}

func NewPing(nonce []byte) *Ping {
	return &Ping{Nonce: nonce}
}

func ParsePing(data []byte) *Ping {
	return &Ping{Nonce: data}
}

func (ping *Ping) ToByte() ([]byte, error) {
	return ping.Nonce, nil
}
