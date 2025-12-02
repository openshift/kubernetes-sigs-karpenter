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

package swarm

import (
	"os"

	"github.com/docker/docker/api/types/filters"
)

// Secret represents a secret.
type Secret struct {
	ID string
	Meta
	Spec SecretSpec
}

// SecretSpec represents a secret specification from a secret in swarm
type SecretSpec struct {
	Annotations

	// Data is the data to store as a secret. It must be empty if a
	// [Driver] is used, in which case the data is loaded from an external
	// secret store. The maximum allowed size is 500KB, as defined in
	// [MaxSecretSize].
	//
	// This field is only used to create the secret, and is not returned
	// by other endpoints.
	//
	// [MaxSecretSize]: https://pkg.go.dev/github.com/moby/swarmkit/v2@v2.0.0-20250103191802-8c1959736554/api/validation#MaxSecretSize
	Data []byte `json:",omitempty"`

	// Driver is the name of the secrets driver used to fetch the secret's
	// value from an external secret store. If not set, the default built-in
	// store is used.
	Driver *Driver `json:",omitempty"`

	// Templating controls whether and how to evaluate the secret payload as
	// a template. If it is not set, no templating is used.
	Templating *Driver `json:",omitempty"`
}

// SecretReferenceFileTarget is a file target in a secret reference
type SecretReferenceFileTarget struct {
	Name string
	UID  string
	GID  string
	Mode os.FileMode
}

// SecretReference is a reference to a secret in swarm
type SecretReference struct {
	File       *SecretReferenceFileTarget
	SecretID   string
	SecretName string
}

// SecretCreateResponse contains the information returned to a client
// on the creation of a new secret.
type SecretCreateResponse struct {
	// ID is the id of the created secret.
	ID string
}

// SecretListOptions holds parameters to list secrets
type SecretListOptions struct {
	Filters filters.Args
}
