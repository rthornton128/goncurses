Overview
--------
Goncurses is an ncurses library for the Go programming language. It
requires both pkg-config and ncurses C development files be installed.

Installation
------------
The go tool is the recommended method of installing goncurses. Issue the
following command on the command line:

$ go get github.com/rthornton128/goncurses

OSX and Windows users should visit the 
[Wiki](https://github.com/rthornton128/goncurses/wiki) for installation
instructions.

Notes
-----

No functions which operate only on stdscr have been implemented because 
it makes little sense to do so in a Go implementation. Stdscr is treated the
same as any other window.

Whenever possible, versions of ncurses functions which could potentially
have a buffer overflow, like the getstr() family of functions, have not been
implemented. Instead, only mvwgetnstr() and wgetnstr() are used.
