package payloads

type Echo struct {
	data []byte
}

func NewEcho(data []byte) *Echo {
	return &Echo{data: data}
}

func ParseEcho(data []byte) *Echo {
	return NewEcho(data)
}

func (echo *Echo) ToByte() ([]byte, error) {
	return echo.data, nil
}
