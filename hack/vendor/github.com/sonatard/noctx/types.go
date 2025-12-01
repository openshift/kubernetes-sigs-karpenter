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

package noctx

import (
	"errors"
	"go/types"
	"strings"

	"github.com/gostaticanalysis/analysisutil"
	"golang.org/x/tools/go/analysis"
)

var errNotFound = errors.New("function not found")

func typeFuncs(pass *analysis.Pass, funcs []string) []*types.Func {
	fs := make([]*types.Func, 0, len(funcs))

	for _, fn := range funcs {
		f, err := typeFunc(pass, fn)
		if err != nil {
			continue
		}

		fs = append(fs, f)
	}

	return fs
}

func typeFunc(pass *analysis.Pass, funcName string) (*types.Func, error) {
	nameParts := strings.Split(strings.TrimSpace(funcName), ".")

	switch len(nameParts) {
	case 2:
		// package function: pkgname.Func
		f, ok := analysisutil.ObjectOf(pass, nameParts[0], nameParts[1]).(*types.Func)
		if !ok || f == nil {
			return nil, errNotFound
		}

		return f, nil
	case 3:
		// method: (*pkgname.Type).Method
		pkgName := strings.TrimLeft(nameParts[0], "(")
		typeName := strings.TrimRight(nameParts[1], ")")

		if pkgName != "" && pkgName[0] == '*' {
			pkgName = pkgName[1:]
			typeName = "*" + typeName
		}

		typ := analysisutil.TypeOf(pass, pkgName, typeName)
		if typ == nil {
			return nil, errNotFound
		}

		m := analysisutil.MethodOf(typ, nameParts[2])
		if m == nil {
			return nil, errNotFound
		}

		return m, nil
	}

	return nil, errNotFound
}
