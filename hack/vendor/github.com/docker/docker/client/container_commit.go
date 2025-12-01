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

package client

import (
	"context"
	"encoding/json"
	"errors"
	"net/url"

	"github.com/distribution/reference"
	"github.com/docker/docker/api/types/container"
)

// ContainerCommit applies changes to a container and creates a new tagged image.
func (cli *Client) ContainerCommit(ctx context.Context, containerID string, options container.CommitOptions) (container.CommitResponse, error) {
	containerID, err := trimID("container", containerID)
	if err != nil {
		return container.CommitResponse{}, err
	}

	var repository, tag string
	if options.Reference != "" {
		ref, err := reference.ParseNormalizedNamed(options.Reference)
		if err != nil {
			return container.CommitResponse{}, err
		}

		if _, isCanonical := ref.(reference.Canonical); isCanonical {
			return container.CommitResponse{}, errors.New("refusing to create a tag with a digest reference")
		}
		ref = reference.TagNameOnly(ref)

		if tagged, ok := ref.(reference.Tagged); ok {
			tag = tagged.Tag()
		}
		repository = ref.Name()
	}

	query := url.Values{}
	query.Set("container", containerID)
	query.Set("repo", repository)
	query.Set("tag", tag)
	query.Set("comment", options.Comment)
	query.Set("author", options.Author)
	for _, change := range options.Changes {
		query.Add("changes", change)
	}
	if !options.Pause {
		query.Set("pause", "0")
	}

	var response container.CommitResponse
	resp, err := cli.post(ctx, "/commit", query, options.Config, nil)
	defer ensureReaderClosed(resp)
	if err != nil {
		return response, err
	}

	err = json.NewDecoder(resp.Body).Decode(&response)
	return response, err
}
