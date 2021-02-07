package gitnotes

import (
	"os/exec"
	"testing"
)

func TestGitExists(t *testing.T) {
	_, err := exec.LookPath("git")
	if err != nil {
		t.Fatalf("Could not find git.exe in path")
	}
}
