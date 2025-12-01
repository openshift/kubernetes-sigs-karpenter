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

package internal

import "golang.org/x/sys/unix"

var names = []struct {
	n string
	m uint32
}{
	{"NOTE_ATTRIB", unix.NOTE_ATTRIB},
	{"NOTE_CHILD", unix.NOTE_CHILD},
	{"NOTE_DELETE", unix.NOTE_DELETE},
	{"NOTE_EXEC", unix.NOTE_EXEC},
	{"NOTE_EXIT", unix.NOTE_EXIT},
	{"NOTE_EXTEND", unix.NOTE_EXTEND},
	{"NOTE_FFAND", unix.NOTE_FFAND},
	{"NOTE_FFCOPY", unix.NOTE_FFCOPY},
	{"NOTE_FFCTRLMASK", unix.NOTE_FFCTRLMASK},
	{"NOTE_FFLAGSMASK", unix.NOTE_FFLAGSMASK},
	{"NOTE_FFNOP", unix.NOTE_FFNOP},
	{"NOTE_FFOR", unix.NOTE_FFOR},
	{"NOTE_FORK", unix.NOTE_FORK},
	{"NOTE_LINK", unix.NOTE_LINK},
	{"NOTE_LOWAT", unix.NOTE_LOWAT},
	{"NOTE_OOB", unix.NOTE_OOB},
	{"NOTE_PCTRLMASK", unix.NOTE_PCTRLMASK},
	{"NOTE_PDATAMASK", unix.NOTE_PDATAMASK},
	{"NOTE_RENAME", unix.NOTE_RENAME},
	{"NOTE_REVOKE", unix.NOTE_REVOKE},
	{"NOTE_TRACK", unix.NOTE_TRACK},
	{"NOTE_TRACKERR", unix.NOTE_TRACKERR},
	{"NOTE_TRIGGER", unix.NOTE_TRIGGER},
	{"NOTE_WRITE", unix.NOTE_WRITE},
}
