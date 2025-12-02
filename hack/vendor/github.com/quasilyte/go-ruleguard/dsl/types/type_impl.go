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

// Method stubs to make various types implement Type interface.
//
// Nothing interesting here, hence it's moved to a separate file.

func (*Array) String() string     { return "" }
func (*Slice) String() string     { return "" }
func (*Pointer) String() string   { return "" }
func (*Interface) String() string { return "" }
func (*Struct) String() string    { return "" }

func (*Array) Underlying() Type     { return nil }
func (*Slice) Underlying() Type     { return nil }
func (*Pointer) Underlying() Type   { return nil }
func (*Interface) Underlying() Type { return nil }
func (*Struct) Underlying() Type    { return nil }
