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

package parser

// Action is a DEC ANSI parser action.
type Action = byte

// These are the actions that the parser can take.
const (
	NoneAction Action = iota
	ClearAction
	CollectAction
	PrefixAction
	DispatchAction
	ExecuteAction
	StartAction // Start of a data string
	PutAction   // Put into the data string
	ParamAction
	PrintAction

	IgnoreAction = NoneAction
)

// nolint: unused
var ActionNames = []string{
	"NoneAction",
	"ClearAction",
	"CollectAction",
	"PrefixAction",
	"DispatchAction",
	"ExecuteAction",
	"StartAction",
	"PutAction",
	"ParamAction",
	"PrintAction",
}

// State is a DEC ANSI parser state.
type State = byte

// These are the states that the parser can be in.
const (
	GroundState State = iota
	CsiEntryState
	CsiIntermediateState
	CsiParamState
	DcsEntryState
	DcsIntermediateState
	DcsParamState
	DcsStringState
	EscapeState
	EscapeIntermediateState
	OscStringState
	SosStringState
	PmStringState
	ApcStringState

	// Utf8State is not part of the DEC ANSI standard. It is used to handle
	// UTF-8 sequences.
	Utf8State
)

// nolint: unused
var StateNames = []string{
	"GroundState",
	"CsiEntryState",
	"CsiIntermediateState",
	"CsiParamState",
	"DcsEntryState",
	"DcsIntermediateState",
	"DcsParamState",
	"DcsStringState",
	"EscapeState",
	"EscapeIntermediateState",
	"OscStringState",
	"SosStringState",
	"PmStringState",
	"ApcStringState",
	"Utf8State",
}
