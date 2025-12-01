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

// Package ini implements parsing of the AWS shared config file.
//
//	Example:
//	sections, err := ini.OpenFile("/path/to/file")
//	if err != nil {
//		panic(err)
//	}
//
//	profile := "foo"
//	section, ok := sections.GetSection(profile)
//	if !ok {
//		fmt.Printf("section %q could not be found", profile)
//	}
package ini

import (
	"fmt"
	"io"
	"os"
	"strings"
)

// OpenFile parses shared config from the given file path.
func OpenFile(path string) (sections Sections, err error) {
	f, oerr := os.Open(path)
	if oerr != nil {
		return Sections{}, &UnableToReadFile{Err: oerr}
	}

	defer func() {
		closeErr := f.Close()
		if err == nil {
			err = closeErr
		} else if closeErr != nil {
			err = fmt.Errorf("close error: %v, original error: %w", closeErr, err)
		}
	}()

	return Parse(f, path)
}

// Parse parses shared config from the given reader.
func Parse(r io.Reader, path string) (Sections, error) {
	contents, err := io.ReadAll(r)
	if err != nil {
		return Sections{}, fmt.Errorf("read all: %v", err)
	}

	lines := strings.Split(string(contents), "\n")
	tokens, err := tokenize(lines)
	if err != nil {
		return Sections{}, fmt.Errorf("tokenize: %v", err)
	}

	return parse(tokens, path), nil
}
