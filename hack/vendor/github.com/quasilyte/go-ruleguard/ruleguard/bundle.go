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

package ruleguard

import (
	"path/filepath"

	"github.com/quasilyte/go-ruleguard/internal/golist"
)

func findBundleFiles(pkgPath string) ([]string, error) { // nolint
	pkg, err := golist.JSON(pkgPath)
	if err != nil {
		return nil, err
	}
	files := make([]string, 0, len(pkg.GoFiles))
	for _, f := range pkg.GoFiles {
		files = append(files, filepath.Join(pkg.Dir, f))
	}
	return files, nil
}
