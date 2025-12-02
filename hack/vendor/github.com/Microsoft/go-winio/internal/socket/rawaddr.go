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

package socket

import (
	"unsafe"
)

// RawSockaddr allows structs to be used with [Bind] and [ConnectEx]. The
// struct must meet the Win32 sockaddr requirements specified here:
// https://docs.microsoft.com/en-us/windows/win32/winsock/sockaddr-2
//
// Specifically, the struct size must be least larger than an int16 (unsigned short)
// for the address family.
type RawSockaddr interface {
	// Sockaddr returns a pointer to the RawSockaddr and its struct size, allowing
	// for the RawSockaddr's data to be overwritten by syscalls (if necessary).
	//
	// It is the callers responsibility to validate that the values are valid; invalid
	// pointers or size can cause a panic.
	Sockaddr() (unsafe.Pointer, int32, error)
}
