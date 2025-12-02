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

package watch

import (
	"fmt"
	"math"
	"time"

	"github.com/onsi/ginkgo/v2/ginkgo/internal"
)

type Suite struct {
	Suite        internal.TestSuite
	RunTime      time.Time
	Dependencies Dependencies

	sharedPackageHashes *PackageHashes
}

func NewSuite(suite internal.TestSuite, maxDepth int, sharedPackageHashes *PackageHashes) (*Suite, error) {
	deps, err := NewDependencies(suite.Path, maxDepth)
	if err != nil {
		return nil, err
	}

	sharedPackageHashes.Add(suite.Path)
	for dep := range deps.Dependencies() {
		sharedPackageHashes.Add(dep)
	}

	return &Suite{
		Suite:        suite,
		Dependencies: deps,

		sharedPackageHashes: sharedPackageHashes,
	}, nil
}

func (s *Suite) Delta() float64 {
	delta := s.delta(s.Suite.Path, true, 0) * 1000
	for dep, depth := range s.Dependencies.Dependencies() {
		delta += s.delta(dep, false, depth)
	}
	return delta
}

func (s *Suite) MarkAsRunAndRecomputedDependencies(maxDepth int) error {
	s.RunTime = time.Now()

	deps, err := NewDependencies(s.Suite.Path, maxDepth)
	if err != nil {
		return err
	}

	s.sharedPackageHashes.Add(s.Suite.Path)
	for dep := range deps.Dependencies() {
		s.sharedPackageHashes.Add(dep)
	}

	s.Dependencies = deps

	return nil
}

func (s *Suite) Description() string {
	numDeps := len(s.Dependencies.Dependencies())
	pluralizer := "ies"
	if numDeps == 1 {
		pluralizer = "y"
	}
	return fmt.Sprintf("%s [%d dependenc%s]", s.Suite.Path, numDeps, pluralizer)
}

func (s *Suite) delta(packagePath string, includeTests bool, depth int) float64 {
	return math.Max(float64(s.dt(packagePath, includeTests)), 0) / float64(depth+1)
}

func (s *Suite) dt(packagePath string, includeTests bool) time.Duration {
	packageHash := s.sharedPackageHashes.Get(packagePath)
	var modifiedTime time.Time
	if includeTests {
		modifiedTime = packageHash.TestModifiedTime
	} else {
		modifiedTime = packageHash.CodeModifiedTime
	}

	return modifiedTime.Sub(s.RunTime)
}
