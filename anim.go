package casper

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"time"
)

func (r *animation) Play(interval time.Duration) {
	for _, f := range r.frames {
		clearConsole()

		fmt.Print(f.data)

		time.Sleep(interval)
	}
}

func clearConsole() {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "linux":
		cmd = exec.Command("clear")
	default:
		cmd = exec.Command("cmd", "/c", "cls")
	}

	cmd.Stdout = os.Stdout
	_ = cmd.Run()
}