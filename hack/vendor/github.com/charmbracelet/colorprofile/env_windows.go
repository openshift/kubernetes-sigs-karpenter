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

//go:build windows
// +build windows

package colorprofile

import (
	"strconv"

	"golang.org/x/sys/windows"
)

func windowsColorProfile(env map[string]string) (Profile, bool) {
	if env["ConEmuANSI"] == "ON" {
		return TrueColor, true
	}

	if len(env["WT_SESSION"]) > 0 {
		// Windows Terminal supports TrueColor
		return TrueColor, true
	}

	major, _, build := windows.RtlGetNtVersionNumbers()
	if build < 10586 || major < 10 {
		// No ANSI support before WindowsNT 10 build 10586
		if len(env["ANSICON"]) > 0 {
			ansiconVer := env["ANSICON_VER"]
			cv, err := strconv.Atoi(ansiconVer)
			if err != nil || cv < 181 {
				// No 8 bit color support before ANSICON 1.81
				return ANSI, true
			}

			return ANSI256, true
		}

		return NoTTY, true
	}

	if build < 14931 {
		// No true color support before build 14931
		return ANSI256, true
	}

	return TrueColor, true
}
