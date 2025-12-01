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

package defaults

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"runtime"
	"strings"
)

var getGOOS = func() string {
	return runtime.GOOS
}

// ResolveDefaultsModeAuto is used to determine the effective aws.DefaultsMode when the mode
// is set to aws.DefaultsModeAuto.
func ResolveDefaultsModeAuto(region string, environment aws.RuntimeEnvironment) aws.DefaultsMode {
	goos := getGOOS()
	if goos == "android" || goos == "ios" {
		return aws.DefaultsModeMobile
	}

	var currentRegion string
	if len(environment.EnvironmentIdentifier) > 0 {
		currentRegion = environment.Region
	}

	if len(currentRegion) == 0 && len(environment.EC2InstanceMetadataRegion) > 0 {
		currentRegion = environment.EC2InstanceMetadataRegion
	}

	if len(region) > 0 && len(currentRegion) > 0 {
		if strings.EqualFold(region, currentRegion) {
			return aws.DefaultsModeInRegion
		}
		return aws.DefaultsModeCrossRegion
	}

	return aws.DefaultsModeStandard
}
