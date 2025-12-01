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

// Package cpuinfo gives runtime info about the current CPU.
//
// This is a very limited module meant for use internally
// in this project. For more versatile solution check
// https://github.com/klauspost/cpuid.
package cpuinfo

// HasBMI1 checks whether an x86 CPU supports the BMI1 extension.
func HasBMI1() bool {
	return hasBMI1
}

// HasBMI2 checks whether an x86 CPU supports the BMI2 extension.
func HasBMI2() bool {
	return hasBMI2
}

// DisableBMI2 will disable BMI2, for testing purposes.
// Call returned function to restore previous state.
func DisableBMI2() func() {
	old := hasBMI2
	hasBMI2 = false
	return func() {
		hasBMI2 = old
	}
}

// HasBMI checks whether an x86 CPU supports both BMI1 and BMI2 extensions.
func HasBMI() bool {
	return HasBMI1() && HasBMI2()
}

var hasBMI1 bool
var hasBMI2 bool
