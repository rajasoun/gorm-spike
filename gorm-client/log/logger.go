package log

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/sirupsen/logrus"
)

// NewZeroLogger returns a new console logger for zerolog
func NewZeroLogger() zerolog.Logger {
	consoleWriter := zerolog.ConsoleWriter{
		Out:          os.Stdout,
		PartsExclude: []string{"time", "level"},
	}
	return log.Output(consoleWriter).With().Logger()
}

// NewlogrusLogger returns a new console logger for logrus without timestamp and level
func NewlogrusLogger() *logrus.Logger {
	logger := logrus.New()
	logger.SetOutput(os.Stdout)
	logger.SetFormatter(&logrus.TextFormatter{
		DisableTimestamp: true,
	})
	return logger
}
