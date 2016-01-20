# USFM Toolkit

[![Build Status](https://travis-ci.org/baijum/usfm.svg?branch=master)](https://travis-ci.org/baijum/usfm)
[![Coverage Status](https://coveralls.io/repos/baijum/usfm/badge.svg?branch=master&service=github)](https://coveralls.io/github/baijum/usfm?branch=master)
[![GoDoc](https://godoc.org/github.com/baijum/usfm?status.svg)](https://godoc.org/github.com/baijum/usfm)
[![MIT License](https://img.shields.io/badge/license-MIT-blue.svg)](https://opensource.org/licenses/MIT)

A toolkit to analyze, parse & convert [USFM formatted text
files](http://paratext.org/about/usfm).

Visit the [wiki for
documentation](https://github.com/baijum/usfm/wiki) about the usage,
architecture etc.,

## Development

If you are interested to contribute to this project, please follow the
instruction given here.

1. Install [Go
compiler](http://muthukadan.net/golang/an-introduction-to-go-programming.html)
2. Run this command to get source code: `go get github.com/baijum/usfm`
   ([Git](http://git-scm.com/) is a prerequisite)

The code will be available under `$GOPATH/src/github.com/baijum/usfm`
You will also a get an executable binary under `$GOPATH/bin/usfm`.

You can fork the project from [Github](https://github.com/baijum/usfm)
and push your changes there.  Later you can send pull request with
your changes.  Before sending the pull request, make sure the tests
are running locally using this command:

    go test ./...

I also recommend to run these tools:

1. [go fmt](https://golang.org/cmd/gofmt/) - format the code as per the Go community standard
2. [go vet](https://golang.org/cmd/vet/) - reports suspicious constructs
3. [golint](https://github.com/golang/lint) - reports more lint issues

Please run these commands before sending pull request.  This will help
us to maintain the quality of code.  You can run the `run.sh` shell
script to run all the above mentioned tools including tests.

## Credits

The parser is created based on the [article written by Ben
Johnson](https://blog.gopheracademy.com/advent-2014/parsers-lexers/)
