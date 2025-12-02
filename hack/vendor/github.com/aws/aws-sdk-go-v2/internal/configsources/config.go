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
	"github.com/aws/aws-sdk-go-v2/aws"
)

// EnableEndpointDiscoveryProvider is an interface for retrieving external configuration value
// for Enable Endpoint Discovery
type EnableEndpointDiscoveryProvider interface {
	GetEnableEndpointDiscovery(ctx context.Context) (value aws.EndpointDiscoveryEnableState, found bool, err error)
}

// ResolveEnableEndpointDiscovery extracts the first instance of a EnableEndpointDiscoveryProvider from the config slice.
// Additionally returns a aws.EndpointDiscoveryEnableState to indicate if the value was found in provided configs,
// and error if one is encountered.
func ResolveEnableEndpointDiscovery(ctx context.Context, configs []interface{}) (value aws.EndpointDiscoveryEnableState, found bool, err error) {
	for _, cfg := range configs {
		if p, ok := cfg.(EnableEndpointDiscoveryProvider); ok {
			value, found, err = p.GetEnableEndpointDiscovery(ctx)
			if err != nil || found {
				break
			}
		}
	}
	return
}

// UseDualStackEndpointProvider is an interface for retrieving external configuration values for UseDualStackEndpoint
type UseDualStackEndpointProvider interface {
	GetUseDualStackEndpoint(context.Context) (value aws.DualStackEndpointState, found bool, err error)
}

// ResolveUseDualStackEndpoint extracts the first instance of a UseDualStackEndpoint from the config slice.
// Additionally returns a boolean to indicate if the value was found in provided configs, and error if one is encountered.
func ResolveUseDualStackEndpoint(ctx context.Context, configs []interface{}) (value aws.DualStackEndpointState, found bool, err error) {
	for _, cfg := range configs {
		if p, ok := cfg.(UseDualStackEndpointProvider); ok {
			value, found, err = p.GetUseDualStackEndpoint(ctx)
			if err != nil || found {
				break
			}
		}
	}
	return
}

// UseFIPSEndpointProvider is an interface for retrieving external configuration values for UseFIPSEndpoint
type UseFIPSEndpointProvider interface {
	GetUseFIPSEndpoint(context.Context) (value aws.FIPSEndpointState, found bool, err error)
}

// ResolveUseFIPSEndpoint extracts the first instance of a UseFIPSEndpointProvider from the config slice.
// Additionally, returns a boolean to indicate if the value was found in provided configs, and error if one is encountered.
func ResolveUseFIPSEndpoint(ctx context.Context, configs []interface{}) (value aws.FIPSEndpointState, found bool, err error) {
	for _, cfg := range configs {
		if p, ok := cfg.(UseFIPSEndpointProvider); ok {
			value, found, err = p.GetUseFIPSEndpoint(ctx)
			if err != nil || found {
				break
			}
		}
	}
	return
}
