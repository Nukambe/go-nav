package commands

import (
	"fmt"
	"github.com/Nukambe/go-nav/internal/nav"
	"github.com/Nukambe/go-nav/internal/raw"
	"os"
	"syscall"
)

type command struct {
	name        string
	description string
	callback    func(directory *nav.Directory)
}

type Commands map[string]command

func InitCommands() Commands {
	cmds := Commands{}
	cmds.registerCommand(quitCommand)
	cmds.registerCommand(upCommand)
	cmds.registerCommand(downCommand)
	cmds.registerCommand(rightCommand)
	cmds.registerCommand(leftCommand)
	cmds.registerCommand(enterCommand)
	return cmds
}

func (cmds Commands) registerCommand(cmd command) {
	cmds[cmd.name] = cmd
}

func (cmds Commands) ReadCommand(state *syscall.Termios, dir *nav.Directory) {
	// buffer to read input
	buffer := make([]byte, 3)

	// read input
	n, err := os.Stdin.Read(buffer)
	if err != nil {
		fmt.Printf("error reading input: %v", err)
		return
	}

	// quit on 'q'
	if n == 1 && buffer[0] == 'q' {
		raw.HandleExit(state, 0)
		// return
	}

	// Enter key (ASCII 10)
	if n == 1 && buffer[0] == 10 {
		cmds["enter"].callback(dir)
		return
	}

	// Handle escape sequences for arrow keys
	if n == 3 && buffer[0] == 27 && buffer[1] == 91 { // ESC + '['
		switch buffer[2] {
		case 65:
			cmds["up"].callback(dir)
			return
		case 66:
			cmds["down"].callback(dir)
			return
		case 67:
			cmds["right"].callback(dir)
			return
		case 68:
			cmds["left"].callback(dir)
			return
		}
	} else {
		fmt.Println("other key")
		return
	}
}
