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

package awsrulesfn

import (
	"net"
	"strings"

	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// IsVirtualHostableS3Bucket returns if the input is a DNS compatible bucket
// name and can be used with Amazon S3 virtual hosted style addressing. Similar
// to [rulesfn.IsValidHostLabel] with the added restriction that the length of label
// must be [3:63] characters long, all lowercase, and not formatted as an IP
// address.
func IsVirtualHostableS3Bucket(input string, allowSubDomains bool) bool {
	// input should not be formatted as an IP address
	// NOTE: this will technically trip up on IPv6 hosts with zone IDs, but
	// validation further down will catch that anyway (it's guaranteed to have
	// unfriendly characters % and : if that's the case)
	if net.ParseIP(input) != nil {
		return false
	}

	var labels []string
	if allowSubDomains {
		labels = strings.Split(input, ".")
	} else {
		labels = []string{input}
	}

	for _, label := range labels {
		// validate special length constraints
		if l := len(label); l < 3 || l > 63 {
			return false
		}

		// Validate no capital letters
		for _, r := range label {
			if r >= 'A' && r <= 'Z' {
				return false
			}
		}

		// Validate valid host label
		if !smithyhttp.ValidHostLabel(label) {
			return false
		}
	}

	return true
}
