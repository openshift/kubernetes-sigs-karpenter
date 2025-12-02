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

package strcase

// We use a lightweight replacement for testify/assert to reduce dependencies

// testingT interface allows us to test our assert functions
type testingT interface {
	Logf(format string, args ...interface{})
	Fail()
}

// assertTrue will fail if the value is not true
func assertTrue(t testingT, value bool) {
	if !value {
		t.Fail()
	}
}

// assertEqual will fail if the two strings are not equal
func assertEqual(t testingT, expected, actual string) {
	if expected != actual {
		t.Logf("Expected: %s Actual: %s", expected, actual)
		t.Fail()
	}
}
