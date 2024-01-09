[![godoc](https://img.shields.io/badge/godoc-reference-blue.svg)](https://pkg.go.dev/github.com/go-corelibs/scanners)
[![codecov](https://codecov.io/gh/go-corelibs/scanners/graph/badge.svg?token=Uy5PFwIWqA)](https://codecov.io/gh/go-corelibs/scanners)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-corelibs/scanners)](https://goreportcard.com/report/github.com/go-corelibs/scanners)

# scanners - useful string scanner utilities

scanners provides useful text scanners that alleviate the need for reading
entire file contents into memory when processing one string at a time would
suffice.

# Installation

``` shell
> go get github.com/go-corelibs/scanners@latest
```

# Examples

## ScanLines

``` go
func main() {
    // read one line at a time from os.Stdin and write to os.Stdout
    stopped := ScanLines(os.Stding, func(line string) (stop bool) {
        fmt.Fprintf(os.Stdout, line + "\n")
        return
    })
    // stopped == false because the ScanLinesFn given never returns true
}
```

## ScanNulls

``` go
func main() {
    // read one null-terminated string at a time from os.Stdin and write to
    // os.Stdout
    stopped := ScanNulls(os.Stding, func(line string) (stop bool) {
        fmt.Fprintf(os.Stdout, line + "\n")
        return
    })
    // stopped == false because the ScanLinesFn given never returns true
}
```

# Go-CoreLibs

[Go-CoreLibs] is a repository of shared code between the [Go-Curses] and
[Go-Enjin] projects.

# License

```
Copyright 2024 The Go-CoreLibs Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use file except in compliance with the License.
You may obtain a copy of the license at

 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
```

[Go-CoreLibs]: https://github.com/go-corelibs
[Go-Curses]: https://github.com/go-curses
[Go-Enjin]: https://github.com/go-enjin
