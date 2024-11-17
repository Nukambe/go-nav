package commands

import "github.com/Nukambe/go-nav/internal/nav"

var downCommand = command{
	name:        "down",
	description: "move the cursor down one line",
	callback: func(dir *nav.Directory) {
		if dir.Target < len(dir.Directories)-1 {
			dir.Target++
		}
	},
}
