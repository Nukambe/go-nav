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
	state, err := raw.EnableRawMode()
	if err != nil {
		fmt.Println("Failed to enable raw mode:", err)
		raw.HandleExit(state, 1)
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	go raw.HandleInterrupt(state, sigChan)

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
		cmd, ok := commands.ReadCommand()
		if ok != nil {
			fmt.Printf("error reading input: %s", cmd)
			raw.HandleExit(state, 1)
		}
		fmt.Println("Command:", cmd)
		if cmd == "quit" {
			raw.HandleExit(state, 0)
		}
	}
}
