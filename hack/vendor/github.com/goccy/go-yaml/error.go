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

package yaml

import (
	"fmt"

	"github.com/goccy/go-yaml/ast"
	"github.com/goccy/go-yaml/internal/errors"
)

var (
	ErrInvalidQuery               = errors.New("invalid query")
	ErrInvalidPath                = errors.New("invalid path instance")
	ErrInvalidPathString          = errors.New("invalid path string")
	ErrNotFoundNode               = errors.New("node not found")
	ErrUnknownCommentPositionType = errors.New("unknown comment position type")
	ErrInvalidCommentMapValue     = errors.New("invalid comment map value. it must be not nil value")
	ErrDecodeRequiredPointerType  = errors.New("required pointer type value")
	ErrExceededMaxDepth           = errors.New("exceeded max depth")
	FormatErrorWithToken          = errors.FormatError
)

type (
	SyntaxError             = errors.SyntaxError
	TypeError               = errors.TypeError
	OverflowError           = errors.OverflowError
	DuplicateKeyError       = errors.DuplicateKeyError
	UnknownFieldError       = errors.UnknownFieldError
	UnexpectedNodeTypeError = errors.UnexpectedNodeTypeError
	Error                   = errors.Error
)

func ErrUnsupportedHeadPositionType(node ast.Node) error {
	return fmt.Errorf("unsupported comment head position for %s", node.Type())
}

func ErrUnsupportedLinePositionType(node ast.Node) error {
	return fmt.Errorf("unsupported comment line position for %s", node.Type())
}

func ErrUnsupportedFootPositionType(node ast.Node) error {
	return fmt.Errorf("unsupported comment foot position for %s", node.Type())
}

// IsInvalidQueryError whether err is ErrInvalidQuery or not.
func IsInvalidQueryError(err error) bool {
	return errors.Is(err, ErrInvalidQuery)
}

// IsInvalidPathError whether err is ErrInvalidPath or not.
func IsInvalidPathError(err error) bool {
	return errors.Is(err, ErrInvalidPath)
}

// IsInvalidPathStringError whether err is ErrInvalidPathString or not.
func IsInvalidPathStringError(err error) bool {
	return errors.Is(err, ErrInvalidPathString)
}

// IsNotFoundNodeError whether err is ErrNotFoundNode or not.
func IsNotFoundNodeError(err error) bool {
	return errors.Is(err, ErrNotFoundNode)
}

// IsInvalidTokenTypeError whether err is ast.ErrInvalidTokenType or not.
func IsInvalidTokenTypeError(err error) bool {
	return errors.Is(err, ast.ErrInvalidTokenType)
}

// IsInvalidAnchorNameError whether err is ast.ErrInvalidAnchorName or not.
func IsInvalidAnchorNameError(err error) bool {
	return errors.Is(err, ast.ErrInvalidAnchorName)
}

// IsInvalidAliasNameError whether err is ast.ErrInvalidAliasName or not.
func IsInvalidAliasNameError(err error) bool {
	return errors.Is(err, ast.ErrInvalidAliasName)
}
