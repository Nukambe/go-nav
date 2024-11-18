package commands

import (
	"fmt"
	"github.com/Nukambe/go-nav/internal/nav"
	"github.com/Nukambe/go-nav/internal/raw"
)

var enterCommand = command{
	name:        "enter",
	description: "change working directory to the target subdirectory and exit go-nav",
	callback: func(dir *nav.Directory) {
		if err := dir.OpenTargetDirectory(); err != nil {
			fmt.Println(err)
			return
		}
		raw.HandleExit(dir.Terminal, 0)
	},
}
