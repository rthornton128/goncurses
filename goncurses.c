// Copyright 2011 Rob Thornton. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include <stdbool.h>
#include <stdlib.h>
#include <curses.h>

#ifdef PDCURSES
bool is_term_resized(int y, int x) { return is_termresized(); }
int resizeterm(int y, int x) { return resize_term(y, x); }
int ncurses_getmouse(MEVENT *me) { return nc_getmouse(me); }
int ncurses_has_key(int ch) { return has_key(ch) == true ? 1 : 0; }
int ncurses_ungetch(int ch) { return PDC_ungetch(ch); }
int ncurses_wattroff(WINDOW *win, int attr ) {
	return wattroff(win, (chtype) attr);
}
int ncurses_wattron(WINDOW *win, int attr) {
	return wattron(win, (chtype) attr);
}
#else
int ncurses_getmouse(MEVENT *me) { return getmouse(me); }
int ncurses_has_key(int ch) { return has_key(ch); }
int ncurses_ungetch(int ch) { return ungetch(ch); }
int ncurses_wattroff(WINDOW *win, int attr) { return wattroff(win, attr); }
int ncurses_wattron(WINDOW *win, int attr) { return wattron(win, attr); }
#endif

int ncurses_COLOR_PAIR(int p) { return COLOR_PAIR(p); }
chtype ncurses_getbkgd(WINDOW *win) { return getbkgd(win); }
void ncurses_getyx(WINDOW *win, int *y, int *x) { getyx(win, *y, *x); }
void ncurses_getbegyx(WINDOW *win, int *y, int *x) { getbegyx(win, *y, *x); }
void ncurses_getmaxyx(WINDOW *win, int *y, int *x) { getmaxyx(win, *y, *x); }

WINDOW *ncurses_wgetparent(const WINDOW *win) {
#ifdef PDCURSES
	return win->_parent;
#else
	return wgetparent(win);
#endif
}

bool ncurses_is_cleared(const WINDOW *win) {
#ifdef PDCURSES
	return win->_clear;
#else
	return is_cleared(win);
#endif
}

bool ncurses_is_keypad(const WINDOW *win) {
#ifdef PDCURSES
	return win->_use_keypad;
#else
	return is_keypad(win);
#endif
}

bool ncurses_is_pad(const WINDOW *win) {
#if defined(PDCURSES) || NCURSES_VERSION_MAJOR < 6
	return false; /* no known built-in way to test for this */
#else
	return is_pad(win);
#endif
}

bool ncurses_is_subwin(const WINDOW *win) {
#ifdef PDCURSES
	return win->_parent != NULL;
#elseif NCURSES_VERSION_MAJOR > 5
	return is_subwin(win);
#else
	return false; /* FIXME */
#endif
}


bool ncurses_has_mouse(void) {
#if NCURSES_VERSION_MINOR < 8
	return false;
#else
	return has_mouse();
#endif
}

int ncurses_touchwin(WINDOW *win) { return touchwin(win); }
int ncurses_untouchwin(WINDOW *win) { return untouchwin(win); }
int ncurses_wattrset(WINDOW *win, int attr) { return wattrset(win, attr); }
int ncurses_wstandend(WINDOW *win) { return wstandend(win); }
int ncurses_wstandout(WINDOW *win) { return wstandout(win); }
