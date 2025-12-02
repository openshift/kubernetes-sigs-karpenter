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

package revocation

import (
	"fmt"
	"sort"
	"strings"

	"golang.org/x/crypto/ocsp"
)

// Reason is used to specify a certificate revocation reason
type Reason int

// ReasonToString provides a map from reason code to string
var ReasonToString = map[Reason]string{
	ocsp.Unspecified:          "unspecified",
	ocsp.KeyCompromise:        "keyCompromise",
	ocsp.CACompromise:         "cACompromise",
	ocsp.AffiliationChanged:   "affiliationChanged",
	ocsp.Superseded:           "superseded",
	ocsp.CessationOfOperation: "cessationOfOperation",
	ocsp.CertificateHold:      "certificateHold",
	// 7 is unused
	ocsp.RemoveFromCRL:      "removeFromCRL",
	ocsp.PrivilegeWithdrawn: "privilegeWithdrawn",
	ocsp.AACompromise:       "aAcompromise",
}

// UserAllowedReasons contains the subset of Reasons which users are
// allowed to use
var UserAllowedReasons = map[Reason]struct{}{
	ocsp.Unspecified:          {},
	ocsp.KeyCompromise:        {},
	ocsp.Superseded:           {},
	ocsp.CessationOfOperation: {},
}

// AdminAllowedReasons contains the subset of Reasons which admins are allowed
// to use. Reasons not found here will soon be forbidden from appearing in CRLs
// or OCSP responses by root programs.
var AdminAllowedReasons = map[Reason]struct{}{
	ocsp.Unspecified:          {},
	ocsp.KeyCompromise:        {},
	ocsp.Superseded:           {},
	ocsp.CessationOfOperation: {},
	ocsp.PrivilegeWithdrawn:   {},
}

// UserAllowedReasonsMessage contains a string describing a list of user allowed
// revocation reasons. This is useful when a revocation is rejected because it
// is not a valid user supplied reason and the allowed values must be
// communicated. This variable is populated during package initialization.
var UserAllowedReasonsMessage = ""

func init() {
	// Build a slice of ints from the allowed reason codes.
	// We want a slice because iterating `UserAllowedReasons` will change order
	// and make the message unpredictable and cumbersome for unit testing.
	// We use []ints instead of []Reason to use `sort.Ints` without fuss.
	var allowed []int
	for reason := range UserAllowedReasons {
		allowed = append(allowed, int(reason))
	}
	sort.Ints(allowed)

	var reasonStrings []string
	for _, reason := range allowed {
		reasonStrings = append(reasonStrings, fmt.Sprintf("%s (%d)",
			ReasonToString[Reason(reason)], reason))
	}
	UserAllowedReasonsMessage = strings.Join(reasonStrings, ", ")
}
