package main

import (
	"fmt"
)

func ClearScreen() {
	fmt.Print("\033[H\033[2J")
}

func ColorString(s, color string) string {
	return fmt.Sprintf("%s%s%s", color, s, Colors.RESET)
}
