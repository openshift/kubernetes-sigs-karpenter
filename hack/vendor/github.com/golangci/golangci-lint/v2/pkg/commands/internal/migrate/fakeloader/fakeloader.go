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

package fakeloader

import (
	"fmt"
	"os"

	"github.com/go-viper/mapstructure/v2"

	"github.com/golangci/golangci-lint/v2/pkg/commands/internal/migrate/parser"
	"github.com/golangci/golangci-lint/v2/pkg/config"
)

// Load is used to keep case of configuration.
// Viper serialize raw map keys in lowercase, this is a problem with the configuration of some linters.
func Load(srcPath string, old any) error {
	file, err := os.Open(srcPath)
	if err != nil {
		return fmt.Errorf("open file: %w", err)
	}

	defer func() { _ = file.Close() }()

	raw := map[string]any{}

	err = parser.Decode(file, raw)
	if err != nil {
		return err
	}

	// NOTE: this is inspired by viper internals.
	cc := &mapstructure.DecoderConfig{
		Result:           old,
		WeaklyTypedInput: true,
		DecodeHook:       config.DecodeHookFunc(),
	}

	decoder, err := mapstructure.NewDecoder(cc)
	if err != nil {
		return fmt.Errorf("constructing mapstructure decoder: %w", err)
	}

	err = decoder.Decode(raw)
	if err != nil {
		return fmt.Errorf("decoding configuration file: %w", err)
	}

	return nil
}
