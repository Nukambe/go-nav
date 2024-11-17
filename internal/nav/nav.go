package nav

import (
	"fmt"
	"os"
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
