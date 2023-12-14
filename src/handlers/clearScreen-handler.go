package handlers

import (
	"os"
	"os/exec"
	"runtime"
)

// ! Function to handle screen clearing/reloading
func ClearScreen() {
	var cmd *exec.Cmd
	// runtime.GOOS: programmatically dermine what the current operating system is.
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	// Before command is ran, need to store the OS standard output inside of the cmd object.
	cmd.Stdout = os.Stdout
	// ! Run Command
	cmd.Run()
}