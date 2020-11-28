package classic

import (
	"os"

	"github.com/sirupsen/logrus"
)

// CreateJSONLogger returns a new logger configured for JSON output
func CreateJSONLogger() *logrus.Logger {
	return &logrus.Logger{
		Out:       os.Stderr,
		Formatter: new(logrus.JSONFormatter),
		Level:     logrus.InfoLevel,
	}
}

// CreateTextLogger returns a new logger configured for text output
func CreateTextLogger() *logrus.Logger {
	return &logrus.Logger{
		Out: os.Stderr,
		Formatter: &logrus.TextFormatter{
			DisableColors: true,
			FullTimestamp: true,
		},
		Level: logrus.InfoLevel,
	}
}
