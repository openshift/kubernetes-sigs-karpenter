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

package gopackages

import (
	"fmt"

	"github.com/dave/dst/decorator/resolver"
	"golang.org/x/tools/go/packages"
)

func New(dir string) *RestorerResolver {
	return &RestorerResolver{Dir: dir}
}

func WithConfig(dir string, config packages.Config) *RestorerResolver {
	return &RestorerResolver{Config: config, Dir: dir}
}

func WithHints(dir string, hints map[string]string) *RestorerResolver {
	return &RestorerResolver{Dir: dir, Hints: hints}
}

type RestorerResolver struct {
	Dir    string
	Config packages.Config

	// Hints (package path -> name) is first checked before asking the packages package
	Hints map[string]string
}

func (r *RestorerResolver) ResolvePackage(path string) (string, error) {

	if name, ok := r.Hints[path]; ok {
		return name, nil
	}

	if r.Dir != "" {
		r.Config.Dir = r.Dir
	}
	r.Config.Mode = packages.LoadTypes
	r.Config.Tests = false

	pkgs, err := packages.Load(&r.Config, "pattern="+path)
	if err != nil {
		return "", err
	}

	if len(pkgs) > 1 {
		return "", fmt.Errorf("%d packages found for %s, %s", len(pkgs), path, r.Config.Dir)
	}
	if len(pkgs) == 0 {
		return "", resolver.ErrPackageNotFound
	}

	p := pkgs[0]

	if len(p.Errors) > 0 {
		return "", p.Errors[0]
	}

	return p.Name, nil
}
