// Copyright (c) 2024  The Go-Curses Authors
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

package scanners

import (
	"bytes"
	"os"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestScanners(t *testing.T) {
	Convey("ScanNulls", t, func(c C) {
		data := []byte{
			'f', 'i', 'r', 's', 't', '\n', 0,
			'n', 'e', 'x', 't', '\n', 'o', 'n', 'e', 0,
			'l', 'a', 's', 't', '\n', 'o', 'n', 'e',
		}

		Convey("Not Stopped", func(c C) {
			buf := bytes.NewBuffer(data)
			var count int
			So(ScanNulls(buf, func(line string) (stop bool) {
				c.So(count, ShouldBeLessThan, 4)
				switch count {
				case 0:
					c.So(line, ShouldEqual, "first\n")
				case 1:
					c.So(line, ShouldEqual, "next\none")
				case 2:
					c.So(line, ShouldEqual, "last\none")
				}
				count += 1
				return
			}), ShouldEqual, false)
		})

		Convey("Stopped Early", func(c C) {
			buf := bytes.NewBuffer(data)
			var count int
			So(ScanNulls(buf, func(line string) (stop bool) {
				stop = true
				count += 1
				c.So(count, ShouldEqual, 1)
				return
			}), ShouldEqual, true)
		})

	})

	Convey("ScanLines", t, func(c C) {
		data := []byte{
			'f', 'i', 'r', 's', 't', '\n', 0,
			'n', 'e', 'x', 't', '\n', 'o', 'n', 'e', '\n',
			'l', 'a', 's', 't', '\n', 'o', 'n', 'e',
		}

		Convey("Not Stopped", func(c C) {
			buf := bytes.NewBuffer(data)
			var count int
			So(ScanLines(buf, func(line string) (stop bool) {
				c.So(count, ShouldBeLessThan, 6)
				switch count {
				case 0:
					c.So(line, ShouldEqual, "first")
				case 1:
					c.So(line, ShouldEqual, "next")
				case 2:
					c.So(line, ShouldEqual, "one")
				case 3:
					c.So(line, ShouldEqual, "last")
				case 4:
					c.So(line, ShouldEqual, "one")
				}
				count += 1
				return
			}), ShouldEqual, false)
		})

		Convey("Stopped Early", func(c C) {
			buf := bytes.NewBuffer(data)
			var count int
			So(ScanLines(buf, func(line string) (stop bool) {
				stop = true
				count += 1
				c.So(count, ShouldEqual, 1)
				return
			}), ShouldEqual, true)
		})
	})

	Convey("ScanFileLines", t, func(c C) {
		data := []byte{
			'f', 'i', 'r', 's', 't', '\n', 0,
			'n', 'e', 'x', 't', '\n', 'o', 'n', 'e', '\n',
			'l', 'a', 's', 't', '\n', 'o', 'n', 'e',
		}
		fh, err := os.CreateTemp("", "corelibs-scanners-%.tmp")
		tmpName := fh.Name()
		fh.Close()
		defer os.Remove(tmpName)
		So(err, ShouldEqual, nil)
		So(tmpName, ShouldNotEqual, "")
		err = os.WriteFile(tmpName, data, 0660)
		So(err, ShouldEqual, nil)
		var count int
		var stopped bool
		stopped, err = ScanFileLines(tmpName, func(line string) (stop bool) {
			c.So(count, ShouldBeLessThan, 6)
			switch count {
			case 0:
				c.So(line, ShouldEqual, "first")
			case 1:
				c.So(line, ShouldEqual, "next")
			case 2:
				c.So(line, ShouldEqual, "one")
			case 3:
				c.So(line, ShouldEqual, "last")
			case 4:
				c.So(line, ShouldEqual, "one")
				stop = true
			}
			count += 1
			return
		})
		So(err, ShouldEqual, nil)
		So(stopped, ShouldEqual, true)
		stopped, err = ScanFileLines(tmpName+".nope", func(line string) (stop bool) {
			return
		})
		So(err, ShouldNotEqual, nil)
		So(stopped, ShouldEqual, false)
	})
}
