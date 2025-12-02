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

package cmd

import (
	"strconv"

	"github.com/spf13/pflag"
)

type boolFlag interface {
	pflag.Value
	IsExplicitlySet() bool
	IsSet() bool
}

type unwrapScalarFlagStrc struct {
	explicitlySet bool
	value         bool
}

func newUnwrapFlag() boolFlag {
	return &unwrapScalarFlagStrc{value: true}
}

func (f *unwrapScalarFlagStrc) IsExplicitlySet() bool {
	return f.explicitlySet
}

func (f *unwrapScalarFlagStrc) IsSet() bool {
	return f.value
}

func (f *unwrapScalarFlagStrc) String() string {
	return strconv.FormatBool(f.value)
}

func (f *unwrapScalarFlagStrc) Set(value string) error {

	v, err := strconv.ParseBool(value)
	f.value = v
	f.explicitlySet = true
	return err
}

func (*unwrapScalarFlagStrc) Type() string {
	return "bool"
}
