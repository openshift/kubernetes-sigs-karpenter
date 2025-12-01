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

/*
Wrapper APIs for in-toto attestation Statement layer protos.
*/

package v1

import "errors"

const StatementTypeUri = "https://in-toto.io/Statement/v1"

var (
	ErrInvalidStatementType  = errors.New("wrong statement type")
	ErrSubjectRequired       = errors.New("at least one subject required")
	ErrDigestRequired        = errors.New("at least one digest required")
	ErrPredicateTypeRequired = errors.New("predicate type required")
	ErrPredicateRequired     = errors.New("predicate object required")
)

func (s *Statement) Validate() error {
	if s.GetType() != StatementTypeUri {
		return ErrInvalidStatementType
	}

	if s.GetSubject() == nil || len(s.GetSubject()) == 0 {
		return ErrSubjectRequired
	}

	// check all resource descriptors in the subject
	subject := s.GetSubject()
	for _, rd := range subject {
		if err := rd.Validate(); err != nil {
			return err
		}

		// v1 statements require the digest to be set in the subject
		if len(rd.GetDigest()) == 0 {
			return ErrDigestRequired
		}
	}

	if s.GetPredicateType() == "" {
		return ErrPredicateTypeRequired
	}

	if s.GetPredicate() == nil {
		return ErrPredicateRequired
	}

	return nil
}
