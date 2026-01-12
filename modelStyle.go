package main

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
