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

package credentials

import (
	"github.com/docker/cli/cli/config/types"
)

// Store is the interface that any credentials store must implement.
type Store interface {
	// Erase removes credentials from the store for a given server.
	Erase(serverAddress string) error
	// Get retrieves credentials from the store for a given server.
	Get(serverAddress string) (types.AuthConfig, error)
	// GetAll retrieves all the credentials from the store.
	GetAll() (map[string]types.AuthConfig, error)
	// Store saves credentials in the store.
	Store(authConfig types.AuthConfig) error
}
