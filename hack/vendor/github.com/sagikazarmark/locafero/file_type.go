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

package locafero

import "io/fs"

// FileType represents the kind of entries [Finder] can return.
type FileType int

// FileType represents the kind of entries [Finder] can return.
const (
	FileTypeAny FileType = iota
	FileTypeFile
	FileTypeDir

	// Deprecated: Use [FileTypeAny] instead.
	FileTypeAll = FileTypeAny
)

func (ft FileType) match(info fs.FileInfo) bool {
	switch ft {
	case FileTypeAny:
		return true

	case FileTypeFile:
		return info.Mode().IsRegular()

	case FileTypeDir:
		return info.IsDir()

	default:
		return false
	}
}
