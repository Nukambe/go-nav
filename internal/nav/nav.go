package nav

import (
	"fmt"
	"os"
	"strings"
)

type Directory struct {
	Pwd         string
	Directories []string
	Files       []string
	Target      int
}

func (dir *Directory) GetDirectory() {
	items, err := os.ReadDir(dir.Pwd)
	if err != nil {
		fmt.Println("could not read current directory:", err)
		return
	}

	for _, item := range items {
		if item.IsDir() {
			dir.Directories = append(dir.Directories, item.Name())
		} else {
			dir.Files = append(dir.Files, item.Name())
		}
	}
}

func (dir *Directory) GetPreview() string {
	next := Directory{Pwd: fmt.Sprintf("%s/%s", dir.Pwd, dir.Directories[dir.Target])}
	next.GetDirectory()

	if len(next.Directories) < 1 {
		return "no subdirectories..."
	}
	return "/" + strings.Join(next.Directories, ", /")
}
