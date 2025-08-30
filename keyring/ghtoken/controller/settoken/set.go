package settoken

import (
	"fmt"
	"io"
	"log/slog"
	"strings"
)

func (c *Controller) Set(logger *slog.Logger) error {
	text, err := c.get(logger)
	if err != nil {
		return fmt.Errorf("get a GitHub access Token: %w", err)
	}
	if err := c.tokenManager.Set(strings.TrimSpace(string(text))); err != nil {
		return fmt.Errorf("set a GitHub access Token to the secret store: %w", err)
	}
	return nil
}

const (
	PanicLevel int = iota
	FatalLevel
	ErrorLevel
	WarnLevel
	InfoLevel
	DebugLevel
	TraceLevel
)

func (c *Controller) get(logger *slog.Logger) ([]byte, error) {
	if c.param.IsStdin {
		s, err := io.ReadAll(c.param.Stdin)
		if err != nil {
			return nil, fmt.Errorf("read a GitHub access token from stdin: %w", err)
		}
		logger.Debug("read a GitHub access token from stdin")
		return s, nil
	}
	text, err := c.term.ReadPassword()
	if err != nil {
		return nil, fmt.Errorf("read a GitHub access Token from terminal: %w", err)
	}
	return text, nil
}
