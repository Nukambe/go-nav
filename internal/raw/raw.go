package raw

import (
	"fmt"
	"os"
	"syscall"
	"unsafe"
)

// EnableRawMode
// configures the terminal to operate in "raw mode," where input is read byte-by-byte without waiting for the "Enter" key.
func EnableRawMode() (*syscall.Termios, error) {
	/*
		syscall.Termios manages terminal attributes.
		Disabling canonical mode (ICANON) allows reading input without pressing Enter.
		Disabling echo (ECHO) prevents the terminal from displaying input.
	*/

	// Get terminal settings
	var oldState syscall.Termios
	// #nosec G103: Use of syscall is intentional and reviewed
	if _, _, err := syscall.Syscall6(syscall.SYS_IOCTL, uintptr(os.Stdin.Fd()), uintptr(syscall.TCGETS), uintptr(unsafe.Pointer(&oldState)), 0, 0, 0); err != 0 {
		return nil, err
	}

	// make a copy to modify
	newState := oldState
	newState.Lflag &^= syscall.ICANON | syscall.ECHO // Disable canonical mode and echoing
	newState.Cc[syscall.VMIN] = 1                    // Minimum number of bytes for read
	newState.Cc[syscall.VTIME] = 0                   // No timeout

	// apply new settings
	// #nosec G103: Use of syscall is intentional and reviewed
	if _, _, err := syscall.Syscall6(syscall.SYS_IOCTL, uintptr(os.Stdin.Fd()), uintptr(syscall.TCSETS), uintptr(unsafe.Pointer(&newState)), 0, 0, 0); err != 0 {
		return nil, err
	}

	return &oldState, nil
}

func restoreRawMode(state *syscall.Termios) {
	// Restore the old terminal settings
	// #nosec G103: Use of syscall is intentional and reviewed
	if _, _, err := syscall.Syscall6(syscall.SYS_IOCTL, uintptr(os.Stdin.Fd()), uintptr(syscall.TCSETS), uintptr(unsafe.Pointer(state)), 0, 0, 0); err != 0 {
		fmt.Printf("error recovering from raw mode: %v", err)
	}
}

func HandleInterrupt(state *syscall.Termios, c chan os.Signal) {
	<-c
	HandleExit(state, 0)
}

func HandleExit(state *syscall.Termios, code int) {
	resetCursor()
	restoreRawMode(state)
	os.Exit(code)
}

func hideCursor() {
	fmt.Print("\033[?25l") // hide cursor
	fmt.Print("\033[1;1H") // move cursor one line up
}

func resetCursor() {
	fmt.Print("\033[?25h")     // show cursor
	fmt.Print("\033[H\033[2J") // Reset to top-left and clear screen
}

func TernaryString(b bool, t string, f string) string {
	if b {
		return t
	}
	return f
}
