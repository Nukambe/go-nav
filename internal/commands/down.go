package commands

import "github.com/Nukambe/go-nav/internal/nav"

var downCommand = command{
	name:        "down",
	description: "move the cursor down one line",
	callback: func(dir *nav.Directory) {
		if dir.Target < len(dir.Directories)-1 { // can increment?
			dir.Target++

			offset := 0
			if dir.IsDownArrow() { // is there a down arrow printed?
				offset = -1 // decrease scroll limit
			}
			if dir.Target >= dir.Window.Start+dir.Window.Height+offset {
				dir.Window.Start++
			}
		}
	},
}
