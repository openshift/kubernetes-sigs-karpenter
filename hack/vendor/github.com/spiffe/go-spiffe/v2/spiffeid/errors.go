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

package spiffeid

import "errors"

var (
	errBadTrustDomainChar = errors.New("trust domain characters are limited to lowercase letters, numbers, dots, dashes, and underscores")
	errBadPathSegmentChar = errors.New("path segment characters are limited to letters, numbers, dots, dashes, and underscores")
	errDotSegment         = errors.New("path cannot contain dot segments")
	errNoLeadingSlash     = errors.New("path must have a leading slash")
	errEmpty              = errors.New("cannot be empty")
	errEmptySegment       = errors.New("path cannot contain empty segments")
	errMissingTrustDomain = errors.New("trust domain is missing")
	errTrailingSlash      = errors.New("path cannot have a trailing slash")
	errWrongScheme        = errors.New("scheme is missing or invalid")
)
