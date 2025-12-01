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

package pflag

// -- func Value
type boolfuncValue func(string) error

func (f boolfuncValue) Set(s string) error { return f(s) }

func (f boolfuncValue) Type() string { return "boolfunc" }

func (f boolfuncValue) String() string { return "" } // same behavior as stdlib 'flag' package

func (f boolfuncValue) IsBoolFlag() bool { return true }

// BoolFunc defines a func flag with specified name, callback function and usage string.
//
// The callback function will be called every time "--{name}" (or any form that matches the flag) is parsed
// on the command line.
func (f *FlagSet) BoolFunc(name string, usage string, fn func(string) error) {
	f.BoolFuncP(name, "", usage, fn)
}

// BoolFuncP is like BoolFunc, but accepts a shorthand letter that can be used after a single dash.
func (f *FlagSet) BoolFuncP(name, shorthand string, usage string, fn func(string) error) {
	var val Value = boolfuncValue(fn)
	flag := f.VarPF(val, name, shorthand, usage)
	flag.NoOptDefVal = "true"
}

// BoolFunc defines a func flag with specified name, callback function and usage string.
//
// The callback function will be called every time "--{name}" (or any form that matches the flag) is parsed
// on the command line.
func BoolFunc(name string, usage string, fn func(string) error) {
	CommandLine.BoolFuncP(name, "", usage, fn)
}

// BoolFuncP is like BoolFunc, but accepts a shorthand letter that can be used after a single dash.
func BoolFuncP(name, shorthand string, usage string, fn func(string) error) {
	CommandLine.BoolFuncP(name, shorthand, usage, fn)
}
