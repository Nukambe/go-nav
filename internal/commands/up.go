package commands

import "github.com/Nukambe/go-nav/internal/nav"

var upCommand = command{
	name:        "up",
	description: "move the cursor up one line",
	callback: func(dir *nav.Directory) {
		if dir.Target > 0 {
			dir.Target--
		}
	},
}
