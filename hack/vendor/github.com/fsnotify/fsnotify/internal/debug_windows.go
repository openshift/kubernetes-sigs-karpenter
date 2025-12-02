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

package internal

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"golang.org/x/sys/windows"
)

func Debug(name string, mask uint32) {
	names := []struct {
		n string
		m uint32
	}{
		{"FILE_ACTION_ADDED", windows.FILE_ACTION_ADDED},
		{"FILE_ACTION_REMOVED", windows.FILE_ACTION_REMOVED},
		{"FILE_ACTION_MODIFIED", windows.FILE_ACTION_MODIFIED},
		{"FILE_ACTION_RENAMED_OLD_NAME", windows.FILE_ACTION_RENAMED_OLD_NAME},
		{"FILE_ACTION_RENAMED_NEW_NAME", windows.FILE_ACTION_RENAMED_NEW_NAME},
	}

	var (
		l       []string
		unknown = mask
	)
	for _, n := range names {
		if mask&n.m == n.m {
			l = append(l, n.n)
			unknown ^= n.m
		}
	}
	if unknown > 0 {
		l = append(l, fmt.Sprintf("0x%x", unknown))
	}
	fmt.Fprintf(os.Stderr, "FSNOTIFY_DEBUG: %s  %-65s → %q\n",
		time.Now().Format("15:04:05.000000000"), strings.Join(l, " | "), filepath.ToSlash(name))
}
