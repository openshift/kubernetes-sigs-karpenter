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
	"net/url"

	"github.com/docker/docker/api/types/container"
)

// ContainerStats returns near realtime stats for a given container.
// It's up to the caller to close the io.ReadCloser returned.
func (cli *Client) ContainerStats(ctx context.Context, containerID string, stream bool) (container.StatsResponseReader, error) {
	containerID, err := trimID("container", containerID)
	if err != nil {
		return container.StatsResponseReader{}, err
	}

	query := url.Values{}
	query.Set("stream", "0")
	if stream {
		query.Set("stream", "1")
	}

	resp, err := cli.get(ctx, "/containers/"+containerID+"/stats", query, nil)
	if err != nil {
		return container.StatsResponseReader{}, err
	}

	return container.StatsResponseReader{
		Body:   resp.Body,
		OSType: resp.Header.Get("Ostype"),
	}, nil
}

// ContainerStatsOneShot gets a single stat entry from a container.
// It differs from `ContainerStats` in that the API should not wait to prime the stats
func (cli *Client) ContainerStatsOneShot(ctx context.Context, containerID string) (container.StatsResponseReader, error) {
	containerID, err := trimID("container", containerID)
	if err != nil {
		return container.StatsResponseReader{}, err
	}

	query := url.Values{}
	query.Set("stream", "0")
	query.Set("one-shot", "1")

	resp, err := cli.get(ctx, "/containers/"+containerID+"/stats", query, nil)
	if err != nil {
		return container.StatsResponseReader{}, err
	}

	return container.StatsResponseReader{
		Body:   resp.Body,
		OSType: resp.Header.Get("Ostype"),
	}, nil
}
