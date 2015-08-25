# Go-Environ [![GoDoc](https://godoc.org/github.com/codehack/go-environ?status.svg)](https://godoc.org/github.com/codehack/go-environ)

*Unix-like environment variable functions in [Go](http://golang.org)*

**Go-Environ** is a system to implement the similar functionality as environment lists found in all Unix-based OS'. Basically, all the functions found at "man 3 setenv" from a Unix prompt. With some additions to support Go's basic types.

## Features

- for convenience, it defaults to "overwrite" from the setenv() context.
- Uses ``sync.pool`` to efficiently use resources when under heavy load.
- get* functions for ``bool``, ``float``, and ``int``.
- ``Index`` and ``Contains`` functions to check for variable existence.

## Installation

Using "go get":

	go get github.com/codehack/go-environ

Then import from source:

	import "github.com/codehack/go-environ"

## Documentation

The full code documentation is located at GoDoc:

[http://godoc.org/github.com/codehack/go-environ](http://godoc.org/github.com/codehack/go-environ)

**Go-Environ** is Copyright (c) 2014-present [Codehack](http://codehack.com).
Published under [MIT License](https://raw.githubusercontent.com/codehack/go-environ/master/LICENSE)



