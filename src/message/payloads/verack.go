package payloads

type Verack struct {
}

func NewVerack(data []byte) *Verack {
	return &Verack{}
}

func ParseVerack(data []byte) *Verack {
	return NewVerack(data)
}

func (verack *Verack) ToByte() ([]byte, error) {
	return []byte{}, nil
}
