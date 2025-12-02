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

package characters

var invalidAsciiTable = [256]bool{
	0x00: true,
	0x01: true,
	0x02: true,
	0x03: true,
	0x04: true,
	0x05: true,
	0x06: true,
	0x07: true,
	0x08: true,
	// 0x09 TAB
	// 0x0A LF
	0x0B: true,
	0x0C: true,
	// 0x0D CR
	0x0E: true,
	0x0F: true,
	0x10: true,
	0x11: true,
	0x12: true,
	0x13: true,
	0x14: true,
	0x15: true,
	0x16: true,
	0x17: true,
	0x18: true,
	0x19: true,
	0x1A: true,
	0x1B: true,
	0x1C: true,
	0x1D: true,
	0x1E: true,
	0x1F: true,
	// 0x20 - 0x7E Printable ASCII characters
	0x7F: true,
}

func InvalidAscii(b byte) bool {
	return invalidAsciiTable[b]
}
