package nav

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"syscall"
)

const NO_SUBDIRECTORIES = "no subdirectories..."

type Directory struct {
	Pwd         string
	Directories []string
	Files       []string
	Target      int
	Window      Window
	Terminal    *syscall.Termios
}
type Window struct {
	Height int
	Width  int
	Start  int
}

func (dir *Directory) GetDirectory() {
	items, err := os.ReadDir(dir.Pwd)
	if err != nil {
		fmt.Println("could not read current directory:", err)
		return
	}

	dir.Directories = []string{}
	dir.Files = []string{}

	for _, item := range items {
		if item.IsDir() {
			dir.Directories = append(dir.Directories, item.Name())
		} else {
			dir.Files = append(dir.Files, item.Name())
		}
	}

	dir.Target = 0
	dir.Window.Start = 0
}

func (dir *Directory) GetPreview() string {
	next := Directory{Pwd: fmt.Sprintf("%s/%s", dir.Pwd, dir.Directories[dir.Target])}
	next.GetDirectory()

	if len(next.Directories) < 1 {
		return NO_SUBDIRECTORIES
	}

	preview := strings.Join(next.Directories, ", /")
	targetPadding := 6 // " ,>, , ,->, " = 6
	extraPadding := 2
	maxLen := dir.Window.Width - len(dir.Directories[dir.Target]) - targetPadding - extraPadding
	if maxLen < 0 {
		return "invalid width"
	}
	if len(preview) > maxLen {
		preview = preview[:maxLen-3] + "..."
	}
	return "/" + preview
}

func (dir *Directory) OpenTargetDirectory() error {
	slash := "/"
	if dir.Pwd == "/" {
		slash = ""
	}
	wd := fmt.Sprintf("%s%s%s", dir.Pwd, slash, dir.Directories[dir.Target])
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		// On Windows, use cmd.exe to open a new terminal and run commands
		cmd = exec.Command("cmd.exe", "/c", "start", "cmd.exe", "/K", fmt.Sprintf("cd %s", wd))
	case "darwin":
		// On macOS, use the `open` command with Terminal
		cmd = exec.Command("open", "-a", "Terminal", wd)
	case "linux":
		if isWSL() {
			// For WSL, open a new WSL terminal, change directory, and keep it open
			cmd = exec.Command("cmd.exe", "/c", "start", "wsl.exe", "bash", "-c", fmt.Sprintf("cd %s && exec bash", wd))
		} else {
			// For native Linux, open a new terminal (gnome-terminal, xterm, etc.)
			cmd = exec.Command("gnome-terminal", "--", "bash", "-c", fmt.Sprintf("cd %s; exec bash", wd))
		}
	default:
		// Unsupported platform
		return fmt.Errorf("unsupported platform: %s", runtime.GOOS)
	}
	if err := cmd.Start(); err != nil {
		return fmt.Errorf("error opening new terminal: %w", err)
	}
	return nil
}

func isWSL() bool {
	// Check /proc/version
	if data, err := os.ReadFile("/proc/version"); err == nil && strings.Contains(strings.ToLower(string(data)), "microsoft") {
		return true
	}

	// Check environment variables
	if _, exists := os.LookupEnv("WSL_INTEROP"); exists {
		return true
	}
	if _, exists := os.LookupEnv("WSL_DISTRO_NAME"); exists {
		return true
	}

	// Check uname
	if out, err := exec.Command("uname", "-r").Output(); err == nil && strings.Contains(strings.ToLower(string(out)), "microsoft") {
		return true
	}

	return false
}
