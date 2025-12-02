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

// +build windows

package open

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	// "syscall"
)

var (
	cmd      = "url.dll,FileProtocolHandler"
	runDll32 = filepath.Join(os.Getenv("SYSTEMROOT"), "System32", "rundll32.exe")
)

func cleaninput(input string) string {
	r := strings.NewReplacer("&", "^&")
	return r.Replace(input)
}

func open(input string) *exec.Cmd {
	cmd := exec.Command(runDll32, cmd, input)
	//cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	return cmd
}

func openWith(input string, appName string) *exec.Cmd {
	cmd := exec.Command("cmd", "/C", "start", "", appName, cleaninput(input))
	//cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	return cmd
}
