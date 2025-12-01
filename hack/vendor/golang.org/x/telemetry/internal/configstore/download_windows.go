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

//go:build windows

package configstore

import (
	"os/exec"
	"syscall"

	"golang.org/x/sys/windows"
)

func init() {
	needNoConsole = needNoConsoleWindows
}

func needNoConsoleWindows(cmd *exec.Cmd) {
	// The uploader main process is likely a daemonized process with no console.
	// (see x/telemetry/start_windows.go) The console creation behavior when
	// a parent is a console process without console is not clearly documented
	// but empirically we observed the new console is created and attached to the
	// subprocess in the default setup.
	//
	// Ensure no new console is attached to the subprocess by setting CREATE_NO_WINDOW.
	//   https://learn.microsoft.com/en-us/windows/console/creation-of-a-console
	//   https://learn.microsoft.com/en-us/windows/win32/procthread/process-creation-flags
	cmd.SysProcAttr = &syscall.SysProcAttr{
		CreationFlags: windows.CREATE_NO_WINDOW,
	}
}
