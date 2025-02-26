package tests

import (
	"os/exec"
	"strings"
	"testing"
)

func TestFileSearchByName(t *testing.T) {
	cmd := exec.Command("go", "run", "main.go", "README.md")
	cmd.Dir = "../"
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("Error running file search by name: %v", err)
	}

	expectedOutput := "README.md"
	if !strings.Contains(string(output), expectedOutput) {
		t.Errorf("Expected file '%s' not found in output. Got: %s", expectedOutput, output)
	}
}
