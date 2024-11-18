package nav

import "fmt"

func (dir *Directory) IsDownArrow() bool {
	return dir.Window.Start+dir.Window.Height < len(dir.Directories)
}

func (dir *Directory) IsUpArrow() bool {
	return dir.Window.Start > 0
}

func (dir *Directory) End() int {
	end := dir.Window.Start + dir.Window.Height
	if end > len(dir.Directories) {
		end = len(dir.Directories)
	}
	return end
}

func (dir *Directory) GetDirectoryText(i int) string {
	if i == dir.Window.Start && dir.IsUpArrow() {
		return " ↑\n"
	}
	if i == dir.End()-1 && dir.IsDownArrow() {
		return " ↓"
	}

	folder := dir.Directories[i]
	target := " "
	preview := ""
	end := "\n"
	if i == dir.Target {
		target = ">"
		preview = fmt.Sprintf(" → \u001B[90m%s\u001B[0m", dir.GetPreview())
	}
	if i == dir.End()-1 {
		end = ""
	}

	return fmt.Sprintf(" %s %s%s%s", target, folder, preview, end)
}
