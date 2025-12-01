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

//go:build windows
// +build windows

package lipgloss

import (
	"sync"

	"github.com/muesli/termenv"
)

var enableANSI sync.Once

// enableANSIColors enables support for ANSI color sequences in the Windows
// default console (cmd.exe and the PowerShell application). Note that this
// only works with Windows 10. Also note that Windows Terminal supports colors
// by default.
func enableLegacyWindowsANSI() {
	enableANSI.Do(func() {
		_, _ = termenv.EnableWindowsANSIConsole()
	})
}
