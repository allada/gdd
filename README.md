# GDD - Go Debugger Devtools

Go debugger using Chrome Devtools as a front-end environment.

## Notes

This is in an alpha stage. Most of the features do not work. This is to experiment there is no promise that this project will continue to exist either. I could use as much help as possible on this project, so if you want to do some of the wiring to Delve or Devtools it would be much appreciated.

## Installation

1. Install Go
2. Install Chrome
3. Install and setup Delve (DLV) [This is the most difficult part. Instructions here: https://github.com/derekparker/delve]
4. Run `go get github.com/allada/gdd`
5. Run `go install github.com/allada/gdd`

## Demo Usage

* Run `cd $GOPATH/src/github.com/allada/gdd/demo`
* Run `$GOPATH/bin/gdd helloworld.go`

This should output a link in the terminal. Open that link in a recent version of Chrome and begin playing.

## Usage Help

```
GDD is a simple debuger that binds Chrome Devtools debugger as a front-end and
uses DLV as a backend.

Usage:

  gdd go-file-to-debug [options] [-- args-to-pass-to-file ...]

Options:

  --port=PORT        Port Number to use to communicate to Chrome Devtools
                     front-end. Default 9922.
  --dlv=FILE         Location of DLV installed on machine. Default (tries to
                     find 'dlv' in system path.
  --help             Prints this help dialog.
```

## License
Copyright 2017 Nathan Bruer

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.