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

package terminfo

import (
	"os"
	"os/user"
	"path"
	"strings"
	"sync"
)

// termCache is the terminfo cache.
var termCache = struct {
	db map[string]*Terminfo
	sync.RWMutex
}{
	db: make(map[string]*Terminfo),
}

// Load follows the behavior described in terminfo(5) to find correct the
// terminfo file using the name, reads the file and then returns a Terminfo
// struct that describes the file.
func Load(name string) (*Terminfo, error) {
	if name == "" {
		return nil, ErrEmptyTermName
	}
	termCache.RLock()
	ti, ok := termCache.db[name]
	termCache.RUnlock()
	if ok {
		return ti, nil
	}
	var checkDirs []string
	// check $TERMINFO
	if dir := os.Getenv("TERMINFO"); dir != "" {
		checkDirs = append(checkDirs, dir)
	}
	// check $HOME/.terminfo
	u, err := user.Current()
	if err != nil {
		return nil, err
	}
	checkDirs = append(checkDirs, path.Join(u.HomeDir, ".terminfo"))
	// check $TERMINFO_DIRS
	if dirs := os.Getenv("TERMINFO_DIRS"); dirs != "" {
		checkDirs = append(checkDirs, strings.Split(dirs, ":")...)
	}
	// check fallback directories
	checkDirs = append(checkDirs, "/etc/terminfo", "/lib/terminfo", "/usr/share/terminfo")
	for _, dir := range checkDirs {
		ti, err = Open(dir, name)
		if err != nil && err != ErrFileNotFound && !os.IsNotExist(err) {
			return nil, err
		} else if err == nil {
			return ti, nil
		}
	}
	return nil, ErrDatabaseDirectoryNotFound
}

// LoadFromEnv loads the terminal info based on the name contained in
// environment variable TERM.
func LoadFromEnv() (*Terminfo, error) {
	return Load(os.Getenv("TERM"))
}
