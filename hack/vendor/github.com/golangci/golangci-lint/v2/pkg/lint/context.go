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

package lint

import (
	"context"
	"fmt"

	"github.com/golangci/golangci-lint/v2/internal/cache"
	"github.com/golangci/golangci-lint/v2/pkg/config"
	"github.com/golangci/golangci-lint/v2/pkg/exitcodes"
	"github.com/golangci/golangci-lint/v2/pkg/goanalysis/load"
	"github.com/golangci/golangci-lint/v2/pkg/lint/linter"
	"github.com/golangci/golangci-lint/v2/pkg/logutils"
)

type ContextBuilder struct {
	cfg *config.Config

	pkgLoader *PackageLoader

	pkgCache *cache.Cache

	loadGuard *load.Guard
}

func NewContextBuilder(cfg *config.Config, pkgLoader *PackageLoader,
	pkgCache *cache.Cache, loadGuard *load.Guard,
) *ContextBuilder {
	return &ContextBuilder{
		cfg:       cfg,
		pkgLoader: pkgLoader,
		pkgCache:  pkgCache,
		loadGuard: loadGuard,
	}
}

func (cl *ContextBuilder) Build(ctx context.Context, log logutils.Log, linters []*linter.Config) (*linter.Context, error) {
	pkgs, deduplicatedPkgs, err := cl.pkgLoader.Load(ctx, linters)
	if err != nil {
		return nil, fmt.Errorf("failed to load packages: %w", err)
	}

	if len(deduplicatedPkgs) == 0 {
		return nil, fmt.Errorf("%w: running `go mod tidy` may solve the problem", exitcodes.ErrNoGoFiles)
	}

	ret := &linter.Context{
		Packages: deduplicatedPkgs,

		// At least `unused` linters works properly only on original (not deduplicated) packages,
		// see https://github.com/golangci/golangci-lint/pull/585.
		OriginalPackages: pkgs,

		Cfg:       cl.cfg,
		Log:       log,
		PkgCache:  cl.pkgCache,
		LoadGuard: cl.loadGuard,
	}

	return ret, nil
}
