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

package aws

// RequestChecksumCalculation controls request checksum calculation workflow
type RequestChecksumCalculation int

const (
	// RequestChecksumCalculationUnset is the unset value for RequestChecksumCalculation
	RequestChecksumCalculationUnset RequestChecksumCalculation = iota

	// RequestChecksumCalculationWhenSupported indicates request checksum will be calculated
	// if the operation supports input checksums
	RequestChecksumCalculationWhenSupported

	// RequestChecksumCalculationWhenRequired indicates request checksum will be calculated
	// if required by the operation or if user elects to set a checksum algorithm in request
	RequestChecksumCalculationWhenRequired
)

// ResponseChecksumValidation controls response checksum validation workflow
type ResponseChecksumValidation int

const (
	// ResponseChecksumValidationUnset is the unset value for ResponseChecksumValidation
	ResponseChecksumValidationUnset ResponseChecksumValidation = iota

	// ResponseChecksumValidationWhenSupported indicates response checksum will be validated
	// if the operation supports output checksums
	ResponseChecksumValidationWhenSupported

	// ResponseChecksumValidationWhenRequired indicates response checksum will only
	// be validated if the operation requires output checksum validation
	ResponseChecksumValidationWhenRequired
)
