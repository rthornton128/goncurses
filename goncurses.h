// Copyright 2011 Rob Thornton. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
#ifndef _GONCURSES_
#define _GONCURSES_ 1

#ifdef PDCURSES
bool is_term_resized(int y, int x);
int resizeterm(int y, int x);
#endif

int ncurses_COLOR_PAIR(int p);
chtype ncurses_getbkgd(WINDOW *win);
void ncurses_getbegyx(WINDOW *win, int *y, int *x);
void ncurses_getmaxyx(WINDOW *win, int *y, int *x);
int ncurses_getmouse(MEVENT *me);
void ncurses_getyx(WINDOW *win, int *y, int *x);
int ncurses_has_key(int);
bool ncurses_has_mouse(void);
bool ncurses_is_cleared(const WINDOW *win);
bool ncurses_is_keypad(const WINDOW *win);
bool ncurses_is_pad(const WINDOW *win);
bool ncurses_is_subwin(const WINDOW *win);
int ncurses_touchwin(WINDOW *win);
int ncurses_ungetch(int ch);
int ncurses_untouchwin(WINDOW *win);
int ncurses_wattroff(WINDOW *, int);
int ncurses_wattron(WINDOW *, int);
int ncurses_wattrset(WINDOW *win, int attr);
WINDOW * ncurses_wgetparent(const WINDOW *win);
int ncurses_wstandend(WINDOW *win);
int ncurses_wstandout(WINDOW *win);

#endif /* _GONCURSES_ */
