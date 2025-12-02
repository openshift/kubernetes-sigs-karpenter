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

// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package asn1 contains supporting types for parsing and building ASN.1
// messages with the cryptobyte package.
package asn1

// Tag represents an ASN.1 identifier octet, consisting of a tag number
// (indicating a type) and class (such as context-specific or constructed).
//
// Methods in the cryptobyte package only support the low-tag-number form, i.e.
// a single identifier octet with bits 7-8 encoding the class and bits 1-6
// encoding the tag number.
type Tag uint8

const (
	classConstructed     = 0x20
	classContextSpecific = 0x80
)

// Constructed returns t with the constructed class bit set.
func (t Tag) Constructed() Tag { return t | classConstructed }

// ContextSpecific returns t with the context-specific class bit set.
func (t Tag) ContextSpecific() Tag { return t | classContextSpecific }

// The following is a list of standard tag and class combinations.
const (
	BOOLEAN           = Tag(1)
	INTEGER           = Tag(2)
	BIT_STRING        = Tag(3)
	OCTET_STRING      = Tag(4)
	NULL              = Tag(5)
	OBJECT_IDENTIFIER = Tag(6)
	ENUM              = Tag(10)
	UTF8String        = Tag(12)
	SEQUENCE          = Tag(16 | classConstructed)
	SET               = Tag(17 | classConstructed)
	PrintableString   = Tag(19)
	T61String         = Tag(20)
	IA5String         = Tag(22)
	UTCTime           = Tag(23)
	GeneralizedTime   = Tag(24)
	GeneralString     = Tag(27)
)
