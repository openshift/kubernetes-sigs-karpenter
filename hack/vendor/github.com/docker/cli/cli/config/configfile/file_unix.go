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

//go:build !windows

package configfile

import (
	"os"
	"syscall"
)

// copyFilePermissions copies file ownership and permissions from "src" to "dst",
// ignoring any error during the process.
func copyFilePermissions(src, dst string) {
	var (
		mode     os.FileMode = 0o600
		uid, gid int
	)

	fi, err := os.Stat(src)
	if err != nil {
		return
	}
	if fi.Mode().IsRegular() {
		mode = fi.Mode()
	}
	if err := os.Chmod(dst, mode); err != nil {
		return
	}

	uid = int(fi.Sys().(*syscall.Stat_t).Uid)
	gid = int(fi.Sys().(*syscall.Stat_t).Gid)

	if uid > 0 && gid > 0 {
		_ = os.Chown(dst, uid, gid)
	}
}
