Installation
------------
The go tool is the recommended method of installing goncurses. Issue the
following command on the command line:

$ go get "github.com/rthornton128/goncurses"

Alternatively, you could use _git_ to install goncurses.


Windows Users
-------------

Goncurses is, strictly speaking, an ncurses library. Ncurses is not available
on Windows without using either a VM running an OS with ncurses or Cygwin.
Instead, Goncurses uses PDCurses which is not 100% compatible with ncurses.
As a result, some features of ncurses will not be available to you. Namely,
the "extended" curses functionality found in the form and menu extentions.

For installation instructions, please see the
[wiki](https://github.com/rthornton128/goncurses/wiki/WindowsInstallation).

OSX Users
---------
For installation instructions, please refer
[here](http://mrcook.uk/how-to-install-go-ncurses-on-mac-osx).


Notes
-----

No functions which operate only on stdscr have been implemented because 
it makes little sense to do so in a Go implementation. Stdscr is treated the
same as any other window.

Whenever possible, versions of ncurses functions which could potentially
have a buffer overflow, like the getstr() family of functions, have not been
implemented. Instead, only mvwgetnstr() and wgetnstr() are used.
