// Copyright (c) 2024  The Go-CoreLibs Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package scanners provides useful text scanners that alleviate the need for
// reading entire file contents into memory when processing one string at a time
// would suffice.
package scanners

import (
	"bufio"
	"io"
	"os"
	"strings"
)

// ScanLinesFn is the callback func for receiving lines of text as they are
// read and if the func returns true, the scanning process will immediately
// stop any further scanning
type ScanLinesFn func(line string) (stop bool)

// ScanNulls reads null-terminated strings of text from the given reader,
// calling the ScanLinesFn for each null terminated string of text
func ScanNulls(reader io.Reader, fn ScanLinesFn) (stopped bool) {
	var line string
	b := make([]byte, 1)
	for {
		if count, err := reader.Read(b); err == nil && count == 1 {
			if rune(b[0]) == rune(0) {
				if stopped = fn(line); stopped {
					return
				}
				line = ""
			} else {
				line += string(b)
			}
			// zero bytes read, nothing to add in line
			continue
		}
		break
	}

	if line != "" {
		// last line ended with EOF instead of NULL
		stopped = fn(line)
	}

	return
}

// ScanLines reads newline terminated strings of text from the given reader,
// calling the ScanLinesFn for each newline terminated string of text
func ScanLines(reader io.Reader, fn ScanLinesFn) (stopped bool) {
	for s := bufio.NewScanner(reader); s.Scan(); {
		// prune nulls from inputs, go strings are not null-terminated so nulls
		// just clutter the strings when a null is a zero-width rune
		cleaned := strings.ReplaceAll(s.Text(), string(rune(0)), "")
		if stopped = fn(cleaned); stopped {
			return
		}
	}
	return
}

// ScanFileLines opens the given file for reading, passing the open *os.File
// and the given ScanLinesFn to the ScanLines function for reading lines one
// at a time
func ScanFileLines(path string, fn ScanLinesFn) (stopped bool, err error) {
	var fh *os.File
	if fh, err = os.Open(path); err != nil {
		return
	}
	defer fh.Close()
	stopped = ScanLines(fh, fn)
	return
}
