Installation
------------
The go tool is probably the best method of installing goncurses.

go get "code.google.com/p/goncurses"

Alternatively you could use mercurial to install goncurses.


Windows Users
-------------

Goncurses is, strictly speaking, an ncurses library. Ncurses is not available
on Windows without using either a VM running an OS with ncurses or Cygwin.
Instead, Goncurses uses PDCurses which is not 100% compatible with ncurses.
As a result, 

NOTES
-----

No functions which operate only on stdscr have been implemented because 
it makes little sense to do so in a Go implementation. Stdscr is treated the
same as any other window.

Whenever possible, versions of ncurses functions which could potentially
have a buffer overflow, like the getstr() family of functions, have not been
implemented. Instead, only the mvwgetnstr() and wgetnstr() are used.
