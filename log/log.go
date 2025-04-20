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

func Set(logE *logrus.Entry, level, color string) error {
	if err := setLevel(logE, level); err != nil {
		return err
	}
	if err := setColor(logE, color); err != nil {
		return err
	}
	return nil
}

func setLevel(logE *logrus.Entry, level string) error {
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

func setColor(logE *logrus.Entry, color string) error {
	switch color {
	case "", "auto":
		return nil
	case "always":
		logE.Logger.SetFormatter(&logrus.TextFormatter{
			ForceColors: true,
		})
		return nil
	case "never":
		logE.Logger.SetFormatter(&logrus.TextFormatter{
			DisableColors: true,
		})
		return nil
	default:
		return errors.New("invalid log_color")
	}
}
