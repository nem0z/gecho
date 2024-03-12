package messages

type Echo struct {
	data []byte
}

func NewEcho(data []byte) *Echo {
	return &Echo{data: data}
}

func (echo *Echo) ToByte() ([]byte, error) {
	return echo.data, nil
}
