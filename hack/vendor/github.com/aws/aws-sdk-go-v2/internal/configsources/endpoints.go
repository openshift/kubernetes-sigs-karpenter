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

package configsources

import (
	"context"
)

// ServiceBaseEndpointProvider is needed to search for all providers
// that provide a configured service endpoint
type ServiceBaseEndpointProvider interface {
	GetServiceBaseEndpoint(ctx context.Context, sdkID string) (string, bool, error)
}

// IgnoreConfiguredEndpointsProvider is needed to search for all providers
// that provide a flag to disable configured endpoints.
//
// Currently duplicated from github.com/aws/aws-sdk-go-v2/config because
// service packages cannot import github.com/aws/aws-sdk-go-v2/config
// due to result import cycle error.
type IgnoreConfiguredEndpointsProvider interface {
	GetIgnoreConfiguredEndpoints(ctx context.Context) (bool, bool, error)
}

// GetIgnoreConfiguredEndpoints is used in knowing when to disable configured
// endpoints feature.
//
// Currently duplicated from github.com/aws/aws-sdk-go-v2/config because
// service packages cannot import github.com/aws/aws-sdk-go-v2/config
// due to result import cycle error.
func GetIgnoreConfiguredEndpoints(ctx context.Context, configs []interface{}) (value bool, found bool, err error) {
	for _, cfg := range configs {
		if p, ok := cfg.(IgnoreConfiguredEndpointsProvider); ok {
			value, found, err = p.GetIgnoreConfiguredEndpoints(ctx)
			if err != nil || found {
				break
			}
		}
	}
	return
}

// ResolveServiceBaseEndpoint is used to retrieve service endpoints from configured sources
// while allowing for configured endpoints to be disabled
func ResolveServiceBaseEndpoint(ctx context.Context, sdkID string, configs []interface{}) (value string, found bool, err error) {
	if val, found, _ := GetIgnoreConfiguredEndpoints(ctx, configs); found && val {
		return "", false, nil
	}

	for _, cs := range configs {
		if p, ok := cs.(ServiceBaseEndpointProvider); ok {
			value, found, err = p.GetServiceBaseEndpoint(context.Background(), sdkID)
			if err != nil || found {
				break
			}
		}
	}
	return
}
