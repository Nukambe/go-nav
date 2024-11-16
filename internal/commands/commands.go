package commands

import (
	"fmt"
	"os"
)

func ReadCommand() (string, error) {
	// buffer to read input
	buffer := make([]byte, 3)

	// read input
	n, err := os.Stdin.Read(buffer)
	if err != nil {
		return "", fmt.Errorf("error reading input: %w", err)
	}

	// quit on 'q'
	if n == 1 && buffer[0] == 'q' {
		return "quit", nil
	}

	var cmd string
	// Handle escape sequences for arrow keys
	if n == 3 && buffer[0] == 27 && buffer[1] == 91 { // ESC + '['
		switch buffer[2] {
		case 65:
			cmd = fmt.Sprintln("up")
		case 66:
			cmd = fmt.Sprintln("down")
		case 67:
			cmd = fmt.Sprintln("right")
		case 68:
			cmd = fmt.Sprintln("left")
		default:
			cmd = fmt.Sprintln("???")
		}
	} else {
		cmd = fmt.Sprintln("other key")
	}
	return cmd, nil
}
