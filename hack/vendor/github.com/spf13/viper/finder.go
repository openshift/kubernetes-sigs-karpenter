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
	"errors"

	"github.com/spf13/afero"
)

// WithFinder sets a custom [Finder].
func WithFinder(f Finder) Option {
	return optionFunc(func(v *Viper) {
		if f == nil {
			return
		}

		v.finder = f
	})
}

// Finder looks for files and directories in an [afero.Fs] filesystem.
type Finder interface {
	Find(fsys afero.Fs) ([]string, error)
}

// Finders combines multiple finders into one.
func Finders(finders ...Finder) Finder {
	return &combinedFinder{finders: finders}
}

// combinedFinder is a Finder that combines multiple finders.
type combinedFinder struct {
	finders []Finder
}

// Find implements the [Finder] interface.
func (c *combinedFinder) Find(fsys afero.Fs) ([]string, error) {
	var results []string
	var errs []error

	for _, finder := range c.finders {
		if finder == nil {
			continue
		}

		r, err := finder.Find(fsys)
		if err != nil {
			errs = append(errs, err)
			continue
		}

		results = append(results, r...)
	}

	return results, errors.Join(errs...)
}
