package main

type Color string

type AvailableColor struct {
	NIL   Color
	RESET Color
	// FOREGROUND COLORS
	FG_GREY           Color
	FG_MAGENTA        Color
	FG_RED            Color
	FG_CYAN           Color
	FG_BRIGHT_RED     Color
	FG_BRIGHT_YELLOW  Color
	FG_BRIGHT_CYAN    Color
	FG_BRIGHT_MAGENTA Color

	//BACKGROUND COLORS
	BG_GREY        Color
	BG_RED         Color
	BG_WHITE       Color
	BG_CYAN        Color
	BG_BRIGHT_CYAN Color
	BG_BRIGHT_RED  Color
}

var Colors = AvailableColor{
	NIL:               "",
	RESET:             "\033[0m",
	FG_RED:            "\033[31m",
	FG_MAGENTA:        "\033[35m",
	FG_CYAN:           "\033[36m",
	FG_GREY:           "\033[90m",
	FG_BRIGHT_RED:     "\033[91m",
	FG_BRIGHT_YELLOW:  "\033[93m",
	FG_BRIGHT_MAGENTA: "\033[95m",
	FG_BRIGHT_CYAN:    "\033[96m",

	BG_CYAN:        "\033[46m",
	BG_RED:         "\033[41m",
	BG_WHITE:       "\033[47m",
	BG_GREY:        "\033[100m",
	BG_BRIGHT_RED:  "\033[101m",
	BG_BRIGHT_CYAN: "\033[106m",
}
