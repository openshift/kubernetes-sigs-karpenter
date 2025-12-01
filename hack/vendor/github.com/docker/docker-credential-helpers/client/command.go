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

package client

import (
	"io"
	"os"
	"os/exec"
)

// Program is an interface to execute external programs.
type Program interface {
	Output() ([]byte, error)
	Input(in io.Reader)
}

// ProgramFunc is a type of function that initializes programs based on arguments.
type ProgramFunc func(args ...string) Program

// NewShellProgramFunc creates a [ProgramFunc] to run command in a [Shell].
func NewShellProgramFunc(command string) ProgramFunc {
	return func(args ...string) Program {
		return createProgramCmdRedirectErr(command, args, nil)
	}
}

// NewShellProgramFuncWithEnv creates a [ProgramFunc] tu run command
// in a [Shell] with the given environment variables.
func NewShellProgramFuncWithEnv(command string, env *map[string]string) ProgramFunc {
	return func(args ...string) Program {
		return createProgramCmdRedirectErr(command, args, env)
	}
}

func createProgramCmdRedirectErr(command string, args []string, env *map[string]string) *Shell {
	ec := exec.Command(command, args...)
	if env != nil {
		for k, v := range *env {
			ec.Env = append(ec.Environ(), k+"="+v)
		}
	}
	ec.Stderr = os.Stderr
	return &Shell{cmd: ec}
}

// Shell invokes shell commands to talk with a remote credentials-helper.
type Shell struct {
	cmd *exec.Cmd
}

// Output returns responses from the remote credentials-helper.
func (s *Shell) Output() ([]byte, error) {
	return s.cmd.Output()
}

// Input sets the input to send to a remote credentials-helper.
func (s *Shell) Input(in io.Reader) {
	s.cmd.Stdin = in
}
