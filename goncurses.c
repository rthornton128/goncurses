#include <stdbool.h>
#include <ncurses.h>

void ncurses_getyx(WINDOW *win, int *y, int *x)
{
	getyx(win, *y, *x);
}

void ncurses_getmaxyx(WINDOW *win, int *y, int *x)
{
	getmaxyx(win, *y, *x);
}

bool ncurses_is_cleared(const WINDOW *win)
{
	return is_cleared(win);
}

bool ncurses_is_keypad(const WINDOW *win)
{
	return is_keypad(win);
}

bool ncurses_is_pad(const WINDOW *win)
{
	return is_pad(win);
}

bool ncurses_is_subwin(const WINDOW *win)
{
	return is_subwin(win);
}


bool ncurses_has_mouse(void)
{
#if NCURSES_VERSION_MINOR < 8
	return false;
#else
	return has_mouse();
#endif
}
