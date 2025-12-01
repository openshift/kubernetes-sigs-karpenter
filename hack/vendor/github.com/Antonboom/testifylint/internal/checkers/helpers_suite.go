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

package checkers

import (
	"go/ast"
	"strings"

	"golang.org/x/tools/go/analysis"
)

func isSuiteMethod(pass *analysis.Pass, fDecl *ast.FuncDecl) bool {
	if fDecl.Recv == nil || len(fDecl.Recv.List) != 1 {
		return false
	}

	rcv := fDecl.Recv.List[0]
	return implementsTestifySuite(pass, rcv.Type)
}

func isSuiteTestMethod(name string) bool {
	return strings.HasPrefix(name, "Test")
}

func isSuiteServiceMethod(name string) bool {
	// https://github.com/stretchr/testify/blob/master/suite/interfaces.go
	switch name {
	case "T", "SetT", "SetS", "SetupSuite", "SetupTest", "TearDownSuite", "TearDownTest",
		"BeforeTest", "AfterTest", "HandleStats", "SetupSubTest", "TearDownSubTest":
		return true
	}
	return false
}

func isSuiteAfterTestMethod(name string) bool {
	// https://github.com/stretchr/testify/blob/master/suite/interfaces.go
	switch name {
	case "TearDownSuite", "TearDownTest", "AfterTest", "HandleStats", "TearDownSubTest":
		return true
	}
	return false
}
