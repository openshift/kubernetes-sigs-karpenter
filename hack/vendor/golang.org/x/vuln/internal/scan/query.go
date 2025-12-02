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

// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package scan

import (
	"context"
	"fmt"
	"regexp"

	"golang.org/x/vuln/internal/client"
	"golang.org/x/vuln/internal/govulncheck"
	isem "golang.org/x/vuln/internal/semver"
)

// runQuery reports vulnerabilities that apply to the queries in the config.
func runQuery(ctx context.Context, handler govulncheck.Handler, cfg *config, c *client.Client) error {
	reqs := make([]*client.ModuleRequest, len(cfg.patterns))
	for i, query := range cfg.patterns {
		mod, ver, err := parseModuleQuery(query)
		if err != nil {
			return err
		}
		if err := handler.Progress(queryProgressMessage(mod, ver)); err != nil {
			return err
		}
		reqs[i] = &client.ModuleRequest{
			Path: mod, Version: ver,
		}
	}

	resps, err := c.ByModules(ctx, reqs)
	if err != nil {
		return err
	}

	ids := make(map[string]bool)
	for _, resp := range resps {
		for _, entry := range resp.Entries {
			if _, ok := ids[entry.ID]; !ok {
				err := handler.OSV(entry)
				if err != nil {
					return err
				}
				ids[entry.ID] = true
			}
		}
	}

	return nil
}

func queryProgressMessage(module, version string) *govulncheck.Progress {
	return &govulncheck.Progress{
		Message: fmt.Sprintf("Looking up vulnerabilities in %s at %s...", module, version),
	}
}

var modQueryRegex = regexp.MustCompile(`(.+)@(.+)`)

func parseModuleQuery(pattern string) (_ string, _ string, err error) {
	matches := modQueryRegex.FindStringSubmatch(pattern)
	// matches should be [module@version, module, version]
	if len(matches) != 3 {
		return "", "", fmt.Errorf("invalid query %s: must be of the form module@version", pattern)
	}
	mod, ver := matches[1], matches[2]
	if !isem.Valid(ver) {
		return "", "", fmt.Errorf("version %s is not valid semver", ver)
	}

	return mod, ver, nil
}
