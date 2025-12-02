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

package load

import (
	"sync"

	"golang.org/x/tools/go/packages"
)

type Guard struct {
	loadMutexes map[*packages.Package]*sync.Mutex
	mutex       sync.Mutex
}

func NewGuard() *Guard {
	return &Guard{
		loadMutexes: map[*packages.Package]*sync.Mutex{},
	}
}

func (g *Guard) AddMutexForPkg(pkg *packages.Package) {
	g.loadMutexes[pkg] = &sync.Mutex{}
}

func (g *Guard) MutexForPkg(pkg *packages.Package) *sync.Mutex {
	return g.loadMutexes[pkg]
}

func (g *Guard) Mutex() *sync.Mutex {
	return &g.mutex
}
