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

//go:build linux

package yqlib

import (
	"io/fs"
	"os"
	"syscall"
)

func changeOwner(info fs.FileInfo, file *os.File) error {
	if stat, ok := info.Sys().(*syscall.Stat_t); ok {
		uid := int(stat.Uid)
		gid := int(stat.Gid)

		err := os.Chown(file.Name(), uid, gid)
		if err != nil {
			// this happens with snap confinement
			// not really a big issue as users can chown
			// the file themselves if required.
			log.Info("Skipping chown: %v", err)
		}
	}
	return nil
}
