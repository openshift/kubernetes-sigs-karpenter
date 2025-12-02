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
	"bytes"
	"context"
	"encoding/json"
	"io"

	"github.com/docker/docker/api/types/volume"
)

// VolumeInspect returns the information about a specific volume in the docker host.
func (cli *Client) VolumeInspect(ctx context.Context, volumeID string) (volume.Volume, error) {
	vol, _, err := cli.VolumeInspectWithRaw(ctx, volumeID)
	return vol, err
}

// VolumeInspectWithRaw returns the information about a specific volume in the docker host and its raw representation
func (cli *Client) VolumeInspectWithRaw(ctx context.Context, volumeID string) (volume.Volume, []byte, error) {
	volumeID, err := trimID("volume", volumeID)
	if err != nil {
		return volume.Volume{}, nil, err
	}

	resp, err := cli.get(ctx, "/volumes/"+volumeID, nil, nil)
	defer ensureReaderClosed(resp)
	if err != nil {
		return volume.Volume{}, nil, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return volume.Volume{}, nil, err
	}

	var vol volume.Volume
	rdr := bytes.NewReader(body)
	err = json.NewDecoder(rdr).Decode(&vol)
	return vol, body, err
}
