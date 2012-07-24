#include <stdbool.h>

bool ncurses_has_mouse(void)
{
#if NCURSES_VERSION_MINOR < 8
	return false;
#else
	return has_mouse();
#endif
}
