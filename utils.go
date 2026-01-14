package main

import (
	"flag"
	"fmt"
	"log"
	"strings"
)

func ParseArgs() (width, height, fps int, fill float64) {
	width = *flag.Int("width", WIDTH, "Override the default width of the board")
	height = *flag.Int("height", HEIGHT, "Override the default height of the board")
	fill = *flag.Float64("fill", FILL, "How filled should the randomly generated matrix be? Value should be between 0 and 1")
	fps = *flag.Int("fps", FPS, "Simulation frame rate (during set up the frame rate is set to a smooth 60fps)")
	flag.Parse()

	if fill < 0 || fill > 1 {
		log.Fatal("-fill value should be contained inside the [0, 1] interval")
	}

	return width, height, fps, fill
}

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
