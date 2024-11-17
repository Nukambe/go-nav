package commands

import (
	"fmt"
	"github.com/Nukambe/go-nav/internal/nav"
)

var rightCommand = command{
	name:        "right",
	description: "move to the selected subdirectory",
	callback: func(dir *nav.Directory) {
		if dir.GetPreview() != nav.NO_SUBDIRECTORIES {
			dir.Pwd = fmt.Sprintf("%s/%s", dir.Pwd, dir.Directories[dir.Target])
			dir.GetDirectory()
		}
	},
}
