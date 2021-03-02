# Overview
Goncurses is an ncurses library for the Go programming language. It
requires both pkg-config and ncurses C development files be installed.

# Installation
The go tool is the recommended method of installing goncurses. Issue the
following command on the command line:

$ go get github.com/rthornton128/goncurses

OSX and Windows users should visit the 
[Wiki](https://github.com/rthornton128/goncurses/wiki) for installation
instructions.

## Pkg-config Flags Error
**Cgo** will fail to build with an invalid or unknown flag error with recent
versions of **ncurses**. Unfortunately, the **cgo** tool only provides one
mechanism for overcoming this. You need to set \*\_ALLOW environment variables
to overcome the issue. There are no **cgo** directives or any other clever
ways (that I know of) to fix this.

This package provides a [Makefile](../master/Makefile) as one solution.
Another would be to set the variables in your shell in whatever way makes
you feel comfortable.

See Issues: [#55](https://github.com/rthornton128/goncurses/issues/55) and
[#56](https://github.com/rthornton128/goncurses/issues/56)

# Notes

No functions which operate only on stdscr have been implemented because 
it makes little sense to do so in a Go implementation. Stdscr is treated the
same as any other window.

Whenever possible, versions of ncurses functions which could potentially
have a buffer overflow, like the getstr() family of functions, have not been
implemented. Instead, only mvwgetnstr() and wgetnstr() are used.
