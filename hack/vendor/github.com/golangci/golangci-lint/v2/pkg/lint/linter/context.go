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

package linter

import (
	"go/ast"

	"golang.org/x/tools/go/packages"

	"github.com/golangci/golangci-lint/v2/internal/cache"
	"github.com/golangci/golangci-lint/v2/pkg/config"
	"github.com/golangci/golangci-lint/v2/pkg/goanalysis/load"
	"github.com/golangci/golangci-lint/v2/pkg/logutils"
)

type Context struct {
	// Packages are deduplicated (test and normal packages) packages
	Packages []*packages.Package

	// OriginalPackages aren't deduplicated: they contain both normal and test
	// version for each of packages
	OriginalPackages []*packages.Package

	Cfg *config.Config
	Log logutils.Log

	PkgCache  *cache.Cache
	LoadGuard *load.Guard
}

func (c *Context) Settings() *config.LintersSettings {
	return &c.Cfg.Linters.Settings
}

func (c *Context) ClearTypesInPackages() {
	for _, p := range c.Packages {
		clearTypes(p)
	}
	for _, p := range c.OriginalPackages {
		clearTypes(p)
	}
}

func clearTypes(p *packages.Package) {
	p.Types = nil
	p.TypesInfo = nil
	p.Syntax = []*ast.File{}
}
