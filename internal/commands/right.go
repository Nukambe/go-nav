package commands

import (
	"fmt"
	"github.com/Nukambe/go-nav/internal/nav"
	"github.com/Nukambe/go-nav/internal/raw"
)

var rightCommand = command{
	name:        "right",
	description: "move to the selected subdirectory",
	callback: func(dir *nav.Directory) {
		if dir.GetPreview() != nav.NO_SUBDIRECTORIES {
			dir.Pwd = fmt.Sprintf("%s/%s", raw.TernaryString(dir.Pwd != "/", dir.Pwd, ""), dir.Directories[dir.Target])
			dir.GetDirectory()
		}
	},
}
