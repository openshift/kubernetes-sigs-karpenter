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

import "context"

// ConfigRemove removes a config.
func (cli *Client) ConfigRemove(ctx context.Context, id string) error {
	id, err := trimID("config", id)
	if err != nil {
		return err
	}
	if err := cli.NewVersionError(ctx, "1.30", "config remove"); err != nil {
		return err
	}
	resp, err := cli.delete(ctx, "/configs/"+id, nil, nil)
	defer ensureReaderClosed(resp)
	return err
}
