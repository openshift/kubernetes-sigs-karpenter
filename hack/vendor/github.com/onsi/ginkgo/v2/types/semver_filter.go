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

package types

import (
	"fmt"

	"github.com/Masterminds/semver/v3"
)

type SemVerFilter func([]string) bool

func MustParseSemVerFilter(input string) SemVerFilter {
	filter, err := ParseSemVerFilter(input)
	if err != nil {
		panic(err)
	}
	return filter
}

func ParseSemVerFilter(filterVersion string) (SemVerFilter, error) {
	if filterVersion == "" {
		return func(_ []string) bool { return true }, nil
	}

	targetVersion, err := semver.NewVersion(filterVersion)
	if err != nil {
		return nil, fmt.Errorf("invalid filter version: %w", err)
	}

	return func(constraints []string) bool {
		// unconstrained specs always run
		if len(constraints) == 0 {
			return true
		}

		for _, constraintStr := range constraints {
			constraint, err := semver.NewConstraint(constraintStr)
			if err != nil {
				return false
			}

			if !constraint.Check(targetVersion) {
				return false
			}
		}

		return true
	}, nil
}

func ValidateAndCleanupSemVerConstraint(semVerConstraint string, cl CodeLocation) (string, error) {
	if len(semVerConstraint) == 0 {
		return "", GinkgoErrors.InvalidEmptySemVerConstraint(cl)
	}
	_, err := semver.NewConstraint(semVerConstraint)
	if err != nil {
		return "", GinkgoErrors.InvalidSemVerConstraint(semVerConstraint, err.Error(), cl)
	}

	return semVerConstraint, nil
}
