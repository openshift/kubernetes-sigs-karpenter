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

// Package imds provides the API client for interacting with the Amazon EC2
// Instance Metadata Service.
//
// All Client operation calls have a default timeout. If the operation is not
// completed before this timeout expires, the operation will be canceled. This
// timeout can be overridden through the following:
//   - Set the options flag DisableDefaultTimeout
//   - Provide a Context with a timeout or deadline with calling the client's operations.
//
// See the EC2 IMDS user guide for more information on using the API.
// https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ec2-instance-metadata.html
package imds
