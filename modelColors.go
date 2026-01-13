package main

type Color string

type AvailableColor struct {
	NIL   Color
	RESET Color

	// FOREGROUND COLORS
	FG_BLACK          Color
	FG_RED            Color
	FG_GREEN          Color
	FG_YELLOW         Color
	FG_BLUE           Color
	FG_MAGENTA        Color
	FG_CYAN           Color
	FG_GREY           Color
	FG_DARK_GREY      Color
	FG_BRIGHT_RED     Color
	FG_BRIGHT_GREEN   Color
	FG_BRIGHT_YELLOW  Color
	FG_BRIGHT_BLUE    Color
	FG_BRIGHT_MAGENTA Color
	FG_BRIGHT_CYAN    Color
	FG_WHITE          Color

	//BACKGROUND COLORS
	BG_BLACK          Color
	BG_RED            Color
	BG_GREEN          Color
	BG_YELLOW         Color
	BG_BLUE           Color
	BG_MAGENTA        Color
	BG_CYAN           Color
	BG_GREY           Color
	BG_DARK_GREY      Color
	BG_BRIGHT_RED     Color
	BG_BRIGHT_GREEN   Color
	BG_BRIGHT_YELLOW  Color
	BG_BRIGHT_BLUE    Color
	BG_BRIGHT_MAGENTA Color
	BG_BRIGHT_CYAN    Color
	BG_WHITE          Color
}

var Colors = AvailableColor{
	NIL:   "",
	RESET: "\033[0m",

	FG_BLACK:          "\033[30m",
	FG_RED:            "\033[31m",
	FG_GREEN:          "\033[32m",
	FG_YELLOW:         "\033[33m",
	FG_BLUE:           "\033[34m",
	FG_MAGENTA:        "\033[35m",
	FG_CYAN:           "\033[36m",
	FG_GREY:           "\033[37m",
	FG_DARK_GREY:      "\033[90m",
	FG_BRIGHT_RED:     "\033[91m",
	FG_BRIGHT_GREEN:   "\033[92m",
	FG_BRIGHT_YELLOW:  "\033[93m",
	FG_BRIGHT_BLUE:    "\033[94m",
	FG_BRIGHT_MAGENTA: "\033[95m",
	FG_BRIGHT_CYAN:    "\033[96m",
	FG_WHITE:          "\033[97m",

	BG_BLACK:          "\033[40m",
	BG_RED:            "\033[41m",
	BG_GREEN:          "\033[42m",
	BG_YELLOW:         "\033[43m",
	BG_BLUE:           "\033[44m",
	BG_MAGENTA:        "\033[45m",
	BG_CYAN:           "\033[46m",
	BG_GREY:           "\033[47m",
	BG_DARK_GREY:      "\033[100m",
	BG_BRIGHT_RED:     "\033[101m",
	BG_BRIGHT_GREEN:   "\033[102m",
	BG_BRIGHT_YELLOW:  "\033[103m",
	BG_BRIGHT_BLUE:    "\033[104m",
	BG_BRIGHT_MAGENTA: "\033[105m",
	BG_BRIGHT_CYAN:    "\033[106m",
	BG_WHITE:          "\033[107m",
}
