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

package garif

type ResultKind string

// declare JSON values
const (
	_pass          ResultKind = "pass"
	_open          ResultKind = "open"
	_informational ResultKind = "informational"
	_notApplicable ResultKind = "notApplicable"
	_review        ResultKind = "review"
	_fail          ResultKind = "fail"
)

// create public visible constants with a namespace as enums
const (
	ResultKind_Pass          ResultKind = _pass
	ResultKind_Open          ResultKind = _open
	ResultKind_Informational ResultKind = _informational
	ResultKind_NotApplicable ResultKind = _notApplicable
	ResultKind_Review        ResultKind = _review
	ResultKind_Fail          ResultKind = _fail
)

type ResultLevel string

// declare JSON values
const (
	_warning ResultLevel = "warning"
	_error   ResultLevel = "error"
	_note    ResultLevel = "note"
	_none    ResultLevel = "none"
)

// create public visible constants with a namespace as enums
const (
	ResultLevel_Warning ResultLevel = _warning
	ResultLevel_Error   ResultLevel = _error
	ResultLevel_Note    ResultLevel = _note
	ResultLevel_None    ResultLevel = _none
)
