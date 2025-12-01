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

// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package scan

import (
	"context"
	"fmt"

	"golang.org/x/tools/go/packages"
	"golang.org/x/vuln/internal/client"
	"golang.org/x/vuln/internal/derrors"
	"golang.org/x/vuln/internal/govulncheck"
	"golang.org/x/vuln/internal/vulncheck"
)

// runSource reports vulnerabilities that affect the analyzed packages.
//
// Vulnerabilities can be called (affecting the package, because a vulnerable
// symbol is actually exercised) or just imported by the package
// (likely having a non-affecting outcome).
func runSource(ctx context.Context, handler govulncheck.Handler, cfg *config, client *client.Client, dir string) (err error) {
	defer derrors.Wrap(&err, "govulncheck")

	if cfg.ScanLevel.WantPackages() && len(cfg.patterns) == 0 {
		return nil // don't throw an error here
	}
	if !gomodExists(dir) {
		return errNoGoMod
	}
	graph := vulncheck.NewPackageGraph(cfg.GoVersion)
	pkgConfig := &packages.Config{
		Dir:   dir,
		Tests: cfg.test,
		Env:   cfg.env,
	}
	if err := graph.LoadPackagesAndMods(pkgConfig, cfg.tags, cfg.patterns, cfg.ScanLevel == govulncheck.ScanLevelSymbol); err != nil {
		if isGoVersionMismatchError(err) {
			return fmt.Errorf("%v\n\n%v", errGoVersionMismatch, err)
		}
		return fmt.Errorf("loading packages: %w", err)
	}

	if cfg.ScanLevel.WantPackages() && len(graph.TopPkgs()) == 0 {
		return nil // early exit
	}
	return vulncheck.Source(ctx, handler, &cfg.Config, client, graph)
}
