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

package command

import "fmt"

type AbortDetails struct {
	ExitCode  int
	Error     error
	EmitUsage bool
}

func Abort(details AbortDetails) {
	panic(details)
}

func AbortGracefullyWith(format string, args ...any) {
	Abort(AbortDetails{
		ExitCode:  0,
		Error:     fmt.Errorf(format, args...),
		EmitUsage: false,
	})
}

func AbortWith(format string, args ...any) {
	Abort(AbortDetails{
		ExitCode:  1,
		Error:     fmt.Errorf(format, args...),
		EmitUsage: false,
	})
}

func AbortWithUsage(format string, args ...any) {
	Abort(AbortDetails{
		ExitCode:  1,
		Error:     fmt.Errorf(format, args...),
		EmitUsage: true,
	})
}

func AbortIfError(preamble string, err error) {
	if err != nil {
		Abort(AbortDetails{
			ExitCode:  1,
			Error:     fmt.Errorf("%s\n%s", preamble, err.Error()),
			EmitUsage: false,
		})
	}
}

func AbortIfErrors(preamble string, errors []error) {
	if len(errors) > 0 {
		out := ""
		for _, err := range errors {
			out += err.Error()
		}
		Abort(AbortDetails{
			ExitCode:  1,
			Error:     fmt.Errorf("%s\n%s", preamble, out),
			EmitUsage: false,
		})
	}
}
