package config

import (
	"os"

	"github.com/sirupsen/logrus"
)

func NamedLogger(name string) logrus.Logger {
	return logrus.Logger{
		Out:       os.Stderr,
		Formatter: &logrus.TextFormatter{},
		Hooks:     make(logrus.LevelHooks),
		Level:     logrus.DebugLevel,
	}
}
