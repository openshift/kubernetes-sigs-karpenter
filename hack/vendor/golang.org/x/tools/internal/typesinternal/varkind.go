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

// Copyright 2024 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package typesinternal

// TODO(adonovan): when CL 645115 lands, define the go1.25 version of
// this API that actually does something.

import "go/types"

type VarKind uint8

const (
	_          VarKind = iota // (not meaningful)
	PackageVar                // a package-level variable
	LocalVar                  // a local variable
	RecvVar                   // a method receiver variable
	ParamVar                  // a function parameter variable
	ResultVar                 // a function result variable
	FieldVar                  // a struct field
)

func (kind VarKind) String() string {
	return [...]string{
		0:          "VarKind(0)",
		PackageVar: "PackageVar",
		LocalVar:   "LocalVar",
		RecvVar:    "RecvVar",
		ParamVar:   "ParamVar",
		ResultVar:  "ResultVar",
		FieldVar:   "FieldVar",
	}[kind]
}

// GetVarKind returns an invalid VarKind.
func GetVarKind(v *types.Var) VarKind { return 0 }

// SetVarKind has no effect.
func SetVarKind(v *types.Var, kind VarKind) {}
