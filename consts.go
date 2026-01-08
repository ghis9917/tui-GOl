package main

import "time"

const BANNER = `
▄▄███▄▄·    ██╗       ██████╗  █████╗ ███╗   ███╗███████╗     ██████╗ ███████╗    ██╗     ██╗███████╗███████╗
██╔════╝    ╚██╗     ██╔════╝ ██╔══██╗████╗ ████║██╔════╝    ██╔═══██╗██╔════╝    ██║     ██║██╔════╝██╔════╝
███████╗     ╚██╗    ██║  ███╗███████║██╔████╔██║█████╗      ██║   ██║█████╗      ██║     ██║█████╗  █████╗  
╚════██║     ██╔╝    ██║   ██║██╔══██║██║╚██╔╝██║██╔══╝      ██║   ██║██╔══╝      ██║     ██║██╔══╝  ██╔══╝  
███████║    ██╔╝     ╚██████╔╝██║  ██║██║ ╚═╝ ██║███████╗    ╚██████╔╝██║         ███████╗██║██║     ███████╗
╚═▀▀▀══╝    ╚═╝       ╚═════╝ ╚═╝  ╚═╝╚═╝     ╚═╝╚══════╝     ╚═════╝ ╚═╝         ╚══════╝╚═╝╚═╝     ╚══════╝
                                                                                                                                                                                                                                   
`

const FPS = 30
const SIMULATION_FRAME_RATE = time.Second / FPS
const UI_FRAME_RATE = time.Second / 60

const WIDTH int = 100
const HEIGHT int = 25

const SPARSITY int = 10
const MAX_INITIAL_POPULATION = (WIDTH * HEIGHT) / SPARSITY

const ALIVE bool = true
const ALIVE_STRING string = "●"
const DEAD bool = false
const DEAD_STRING string = "○"

type Color struct {
	RESET string
	// FOREGROUND COLORS
	FG_GREY          string
	FG_MAGENTA       string
	FG_RED           string
	FG_CYAN          string
	FG_BRIGHT_RED    string
	FG_BRIGHT_YELLOW string
	FG_BRIGHT_CYAN   string

	//BACKGROUND COLORS
	BG_GREY        string
	BG_WHITE       string
	BG_CYAN        string
	BG_BRIGHT_CYAN string
}

var Colors = Color{
	RESET:            "\033[0m",
	FG_RED:           "\033[31m",
	FG_MAGENTA:       "\033[35m",
	FG_GREY:          "\033[90m",
	FG_CYAN:          "\033[36m",
	FG_BRIGHT_RED:    "\033[91m",
	FG_BRIGHT_YELLOW: "\033[93m",
	FG_BRIGHT_CYAN:   "\033[96m",

	BG_GREY:        "\033[100m",
	BG_WHITE:       "\033[47m",
	BG_BRIGHT_CYAN: "\033[106m",
	BG_CYAN:        "\033[46m",
}

type Style struct {
	BOLD      string
	UNDERLINE string
	OVERLINE  string
	BLINK     string
	REVERSE   string
}

var TextStyle = Style{
	BOLD:      "\033[1m",
	UNDERLINE: "\033[4m",
	OVERLINE:  "\033[53m",
	BLINK:     "\033[5m",
	REVERSE:   "\033[7m",
}
