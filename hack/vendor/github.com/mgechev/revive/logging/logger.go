/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Package logging provides a logger and related methods.
package logging

import (
	"io"
	"log/slog"
	"os"
)

const logFile = "revive.log"

var logger *slog.Logger

// GetLogger retrieves an instance of an application logger which outputs
// to a file if the debug flag is enabled.
func GetLogger() (*slog.Logger, error) {
	if logger != nil {
		return logger, nil
	}

	debugModeEnabled := os.Getenv("DEBUG") != ""
	if !debugModeEnabled {
		// by default, suppress all logging output
		return slog.New(slog.NewTextHandler(io.Discard, nil)), nil // TODO: change to slog.New(slog.DiscardHandler) when we switch to Go 1.24
	}

	fileWriter, err := os.Create(logFile)
	if err != nil {
		return nil, err
	}

	logger = slog.New(slog.NewTextHandler(io.MultiWriter(os.Stderr, fileWriter), nil))

	logger.Info("Logger initialized", "logFile", logFile)

	return logger, nil
}
