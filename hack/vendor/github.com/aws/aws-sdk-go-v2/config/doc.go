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

// Package config provides utilities for loading configuration from multiple
// sources that can be used to configure the SDK's API clients, and utilities.
//
// The config package will load configuration from environment variables, AWS
// shared configuration file (~/.aws/config), and AWS shared credentials file
// (~/.aws/credentials).
//
// Use the LoadDefaultConfig to load configuration from all the SDK's supported
// sources, and resolve credentials using the SDK's default credential chain.
//
// LoadDefaultConfig allows for a variadic list of additional Config sources that can
// provide one or more configuration values which can be used to programmatically control the resolution
// of a specific value, or allow for broader range of additional configuration sources not supported by the SDK.
// A Config source implements one or more provider interfaces defined in this package. Config sources passed in will
// take precedence over the default environment and shared config sources used by the SDK. If one or more Config sources
// implement the same provider interface, priority will be handled by the order in which the sources were passed in.
//
// A number of helpers (prefixed by “With“)  are provided in this package that implement their respective provider
// interface. These helpers should be used for overriding configuration programmatically at runtime.
package config
