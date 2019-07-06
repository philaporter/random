package app

import (
	"os"
	"os/exec"
	"testing"
)

func TestShutdownHandler(t *testing.T) {
	if os.Getenv("SHUTDOWN") == "1" {
		shutdownHandler()
		return
	}
	cmd := exec.Command(os.Args[0], "-test.run=TestShutdownHandler")
	cmd.Env = append(os.Environ(), "SHUTDOWN=1")
	err := cmd.Run()
	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		return
	}
	t.Fatalf("process ran with err %v, want exit status 1", err)
}