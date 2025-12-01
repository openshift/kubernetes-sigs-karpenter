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

// Package goenv A set of functions to get information from `go env`.
package goenv

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"os/exec"
	"strings"
)

// GetAll gets information from "go env".
func GetAll(ctx context.Context) (map[string]string, error) {
	v, err := Get(ctx)
	if err != nil {
		return nil, err
	}

	return v, nil
}

// GetOne gets information from "go env" for one environment variable.
func GetOne(ctx context.Context, name string) (string, error) {
	v, err := Get(ctx, name)
	if err != nil {
		return "", err
	}

	return v[name], nil
}

// Get gets information from "go env" for one or several environment variables.
func Get(ctx context.Context, name ...string) (map[string]string, error) {
	args := append([]string{"env", "-json"}, name...)
	cmd := exec.CommandContext(ctx, "go", args...) //nolint:gosec // The env var names must be checked by the user.

	out, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("command %q: %w: %s", strings.Join(cmd.Args, " "), err, string(out))
	}

	v := map[string]string{}

	err = json.NewDecoder(bytes.NewBuffer(out)).Decode(&v)
	if err != nil {
		return nil, err
	}

	return v, nil
}
