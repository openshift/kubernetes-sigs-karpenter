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

package container

import "github.com/docker/docker/api/types/filters"

// ResizeOptions holds parameters to resize a TTY.
// It can be used to resize container TTYs and
// exec process TTYs too.
type ResizeOptions struct {
	Height uint
	Width  uint
}

// AttachOptions holds parameters to attach to a container.
type AttachOptions struct {
	Stream     bool
	Stdin      bool
	Stdout     bool
	Stderr     bool
	DetachKeys string
	Logs       bool
}

// CommitOptions holds parameters to commit changes into a container.
type CommitOptions struct {
	Reference string
	Comment   string
	Author    string
	Changes   []string
	Pause     bool
	Config    *Config
}

// RemoveOptions holds parameters to remove containers.
type RemoveOptions struct {
	RemoveVolumes bool
	RemoveLinks   bool
	Force         bool
}

// StartOptions holds parameters to start containers.
type StartOptions struct {
	CheckpointID  string
	CheckpointDir string
}

// ListOptions holds parameters to list containers with.
type ListOptions struct {
	Size    bool
	All     bool
	Latest  bool
	Since   string
	Before  string
	Limit   int
	Filters filters.Args
}

// LogsOptions holds parameters to filter logs with.
type LogsOptions struct {
	ShowStdout bool
	ShowStderr bool
	Since      string
	Until      string
	Timestamps bool
	Follow     bool
	Tail       string
	Details    bool
}
