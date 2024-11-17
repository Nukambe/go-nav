package main

import (
	"fmt"
	"github.com/Nukambe/go-nav/internal/commands"
	"github.com/Nukambe/go-nav/internal/nav"
	"github.com/Nukambe/go-nav/internal/raw"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// enable raw mode
	state, err := raw.EnableRawMode()
	if err != nil {
		fmt.Println("Failed to enable raw mode:", err)
		raw.HandleExit(state, 1)
	}

	// listen for signal interruption
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	go raw.HandleInterrupt(state, sigChan)

	// register commands
	cmds := commands.InitCommands()

	// init directory
	navDir := nav.Directory{}
	if wd, err := os.Getwd(); err != nil {
		fmt.Println("could not get current directory:", err)
		raw.HandleExit(state, 1)
	} else {
		navDir.Pwd = wd
		navDir.GetDirectory()
	}

	for {
		raw.ClearScreen()
		raw.DrawScreen(&navDir)
		cmds.ReadCommand(state, &navDir)
	}
}
