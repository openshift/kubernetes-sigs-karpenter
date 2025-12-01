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

// This package provides utilities for efficiently performing Win32 IO operations in Go.
// Currently, this package is provides support for genreal IO and management of
//   - named pipes
//   - files
//   - [Hyper-V sockets]
//
// This code is similar to Go's [net] package, and uses IO completion ports to avoid
// blocking IO on system threads, allowing Go to reuse the thread to schedule other goroutines.
//
// This limits support to Windows Vista and newer operating systems.
//
// Additionally, this package provides support for:
//   - creating and managing GUIDs
//   - writing to [ETW]
//   - opening and manageing VHDs
//   - parsing [Windows Image files]
//   - auto-generating Win32 API code
//
// [Hyper-V sockets]: https://docs.microsoft.com/en-us/virtualization/hyper-v-on-windows/user-guide/make-integration-service
// [ETW]: https://docs.microsoft.com/en-us/windows-hardware/drivers/devtest/event-tracing-for-windows--etw-
// [Windows Image files]: https://docs.microsoft.com/en-us/windows-hardware/manufacture/desktop/work-with-windows-images
package winio
