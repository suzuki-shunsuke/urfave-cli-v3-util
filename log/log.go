package log

import (
	"errors"
	"fmt"

	"github.com/sirupsen/logrus"
)

func New(program, version string) *logrus.Entry {
	logger := logrus.New()
	logger.Formatter = &logrus.TextFormatter{
		DisableQuote: true,
	}
	return logger.WithFields(logrus.Fields{
		"program_version": version,
		"program":         program,
	})
}

func SetLevel(level string, logE *logrus.Entry) error {
	if level == "" {
		return nil
	}
	lvl, err := logrus.ParseLevel(level)
	if err != nil {
		return fmt.Errorf("parse log level: %w", err)
	}
	logE.Logger.Level = lvl
	return nil
}

func SetColor(color string, logE *logrus.Entry) error {
	switch color {
	case "", "auto":
		return nil
	case "always":
		logrus.SetFormatter(&logrus.TextFormatter{
			ForceColors: true,
		})
		return nil
	case "never":
		logrus.SetFormatter(&logrus.TextFormatter{
			DisableColors: true,
		})
		return nil
	default:
		return errors.New("invalid log_color")
	}
}
