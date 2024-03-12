package message

import "fmt"

const CommandLength = 12

var commandList = map[string]uint8{
	"hello": 0,
}

func GetCommandKey(command string) (uint8, error) {
	if key, ok := commandList[command]; ok {
		return key, nil
	}

	return 0, fmt.Errorf("command not found")
}

func IsValidCommand(command string) bool {
	_, ok := commandList[command]
	return ok
}

func ParseCommand(command []byte) (string, error) {
	if len(command) != CommandLength {
		return "", fmt.Errorf("invalid command length")
	}

	len := 0
	for ; len < CommandLength && command[len] != 0; len++ {
	}

	return string(command[:len]), nil
}

func FormatCommand(command string) ([]byte, error) {
	diff := CommandLength - len(command)
	if diff < 0 {
		return nil, fmt.Errorf("invalid command length")
	}

	nullBytes := make([]byte, diff)
	return append([]byte(command), nullBytes...), nil
}
