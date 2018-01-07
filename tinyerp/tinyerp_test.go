package tinyerp

import (
	"os"
	"os/exec"
	"testing"
)

func TestMustEnv(t *testing.T) {
	// If env variable "TEST" is not set then fatal it
	var value string
	key := "TEST"

	if value = os.Getenv(key); value != "" {
		os.Unsetenv(key)
		defer func() { os.Setenv(key, value) }()
	}

	if os.Getenv("TEST_MUST_ENV") == "1" {
		mustEnv(key)
		return
	}

	cmd := exec.Command(os.Args[0], "-test.run=TestMustEnv")
	cmd.Env = append(os.Environ(), "TEST_MUST_ENV=1")
	err := cmd.Run()
	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		return
	}
	t.Errorf("mustEnv() failed to fatal if env %q is not present", key)
}
