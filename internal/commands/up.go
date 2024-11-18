package commands

import "github.com/Nukambe/go-nav/internal/nav"

var upCommand = command{
	name:        "up",
	description: "move the cursor up one line",
	callback: func(dir *nav.Directory) {
		if dir.Target > 0 { // can decrement?
			dir.Target--

			offset := 0
			if dir.IsUpArrow() {
				offset = 1
			}
			if dir.Target < dir.Window.Start+offset {
				dir.Window.Start--
			}
		}
	},
}
