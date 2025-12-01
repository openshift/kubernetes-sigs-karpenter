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

package viper

import (
	"context"
	"log/slog"
)

// WithLogger sets a custom logger.
func WithLogger(l *slog.Logger) Option {
	return optionFunc(func(v *Viper) {
		v.logger = l
	})
}

type discardHandler struct{}

func (n *discardHandler) Enabled(_ context.Context, _ slog.Level) bool {
	return false
}

func (n *discardHandler) Handle(_ context.Context, _ slog.Record) error {
	return nil
}

func (n *discardHandler) WithAttrs(_ []slog.Attr) slog.Handler {
	return n
}

func (n *discardHandler) WithGroup(_ string) slog.Handler {
	return n
}
