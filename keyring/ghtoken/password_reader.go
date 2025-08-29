package ghtoken

import (
	"fmt"
	"io"
	"syscall"

	"github.com/charmbracelet/x/term"
)

type PasswordReader struct {
	stdout io.Writer
}

func NewPasswordReader(stdout io.Writer) *PasswordReader {
	return &PasswordReader{
		stdout: stdout,
	}
}

func (p *PasswordReader) ReadPassword() ([]byte, error) {
	fmt.Fprint(p.stdout, "Enter a GitHub access token: ")
	b, err := term.ReadPassword(uintptr(syscall.Stdin))
	fmt.Fprintln(p.stdout, "")
	if err != nil {
		return nil, fmt.Errorf("read a GitHub access token from terminal: %w", err)
	}
	return b, nil
}
