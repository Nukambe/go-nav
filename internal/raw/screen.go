package raw

import (
	"fmt"
	"github.com/Nukambe/go-nav/internal/nav"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"syscall"
	"unsafe"
)

func ClearScreen() {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		if err := cmd.Run(); err != nil {
			return
		}
	} else {
		// \033 = ESC, [H = Move cursor to home (1,1), [2J = Clear screen, [3J = Clear scroll-back
		fmt.Print("\033[H\033[2J\033[3J")
	}
}

func DrawScreen(dir *nav.Directory) {
	width, _, err := getScreenSize()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Working Directory: %s/\n", dir.Pwd) // print working directory
	fmt.Println(strings.Repeat("-", width))         // print line across window

	fmt.Println("Directories:")
	for _, directory := range dir.Directories { // print all directories
		fmt.Printf("	/%s\n", directory)
	}

	fmt.Println("Files:")
	for _, file := range dir.Files {
		fmt.Printf("	%s\n", file)
	}
}

// getScreenSize
// gets the width and height of the terminal (Linux, macOS)
func getScreenSize() (int, int, error) {
	var ws struct {
		Row    uint16
		Col    uint16
		Xpixel uint16
		Ypixel uint16
	}
	/*
		syscall.SYS_IOCTL: Used to interact with device drivers.
		syscall.TIOCGWINSZ: requests the terminal's size.
		This method works only on Unix-like systems (Linux, macOS).
	*/

	_, _, err := syscall.Syscall(syscall.SYS_IOCTL, uintptr(syscall.Stdin), uintptr(syscall.TIOCGWINSZ), uintptr(unsafe.Pointer(&ws)))
	if err != 0 {
		return 0, 0, fmt.Errorf("error getting terminal size: %v", err)
	}

	return int(ws.Col), int(ws.Row), nil
}
