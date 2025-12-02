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

package fsutils

import (
	"fmt"
	"os"
	"sync"

	"github.com/golangci/golangci-lint/v2/pkg/logutils"
)

type FileCache struct {
	files sync.Map
}

func NewFileCache() *FileCache {
	return &FileCache{}
}

func (fc *FileCache) GetFileBytes(filePath string) ([]byte, error) {
	cachedBytes, ok := fc.files.Load(filePath)
	if ok {
		return cachedBytes.([]byte), nil
	}

	fileBytes, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("can't read file %s: %w", filePath, err)
	}

	fc.files.Store(filePath, fileBytes)
	return fileBytes, nil
}

func PrettifyBytesCount(n int64) string {
	const (
		Multiplexer = 1024
		KiB         = 1 * Multiplexer
		MiB         = KiB * Multiplexer
		GiB         = MiB * Multiplexer
	)

	if n >= GiB {
		return fmt.Sprintf("%.1fGiB", float64(n)/GiB)
	}
	if n >= MiB {
		return fmt.Sprintf("%.1fMiB", float64(n)/MiB)
	}
	if n >= KiB {
		return fmt.Sprintf("%.1fKiB", float64(n)/KiB)
	}
	return fmt.Sprintf("%dB", n)
}

func (fc *FileCache) PrintStats(log logutils.Log) {
	var size int64
	var mapLen int
	fc.files.Range(func(_, fileBytes any) bool {
		mapLen++
		size += int64(len(fileBytes.([]byte)))

		return true
	})

	log.Infof("File cache stats: %d entries of total size %s", mapLen, PrettifyBytesCount(size))
}
