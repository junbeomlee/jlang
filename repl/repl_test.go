package repl

import (
	"os"
	"testing"
)

func TestStart(t *testing.T) {
	Start(os.Stdin, os.Stdout)
}
