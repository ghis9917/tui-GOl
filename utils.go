package main

import (
	"fmt"
	"strings"
)

func ClearScreen() {
	fmt.Print("\033[H\033[2J")
}

func ColoredString(s, color string) string {
	return fmt.Sprintf("%s%s%s", color, s, Colors.RESET)
}

func ApplyStyle(s string, styles ...string) string {
	return strings.Join(styles, "") + s + "\033[0m"
}
