package main

import (
	"fmt"
	"strings"
)

func Println(strs ...string) {
	for _, s := range strs {
		fmt.Print(ColoredString(s, THEME.Primary, THEME.Background))
	}
	fmt.Println()
}

func ClearScreen() {
	fmt.Print("\033[H\033[2J")
	fmt.Print("\033[3J")
}

func ColoredString(s string, foreground, background Color) string {
	return fmt.Sprintf("%s%s%s%s", foreground, background, s, Colors.RESET)
}

func ApplyStyle(s string, styles ...string) string {
	return strings.Join(styles, "") + s + "\033[0m"
}

func countRune(s string, r rune) int {
	count := 0
	for _, c := range s {
		if c == r {
			count++
		}
	}
	return count
}
