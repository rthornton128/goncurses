include $(GOROOT)/src/Make.inc

TARG=goncurses
CGOFILES=\
	goncurses.go\
	panel.go\

include $(GOROOT)/src/Make.pkg
