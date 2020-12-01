// Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0
// This product includes software developed at Datadog (https://www.datadoghq.com/). Copyright 2020 Datadog, Inc.

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
