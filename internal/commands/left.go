package commands

import (
	"github.com/Nukambe/go-nav/internal/nav"
	"strings"
)

var leftCommand = command{
	name:        "left",
	description: "move to the parent directory",
	callback: func(dir *nav.Directory) {
		if len(dir.Pwd) > 1 {
			i := strings.LastIndex(dir.Pwd, "/")
			if i > 0 {
				dir.Pwd = dir.Pwd[:i]
				dir.GetDirectory()
			}
			if i == 0 {
				dir.Pwd = "/"
				dir.GetDirectory()
			}
		}
	},
}
