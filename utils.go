package main

import (
	"fmt"
	"strings"
)

func ClearScreen() {
	fmt.Print("\033[H\033[2J")
	fmt.Print("\033[3J")
}

func ColoredString(s, color string) string {
	return fmt.Sprintf("%s%s%s", color, s, Colors.RESET)
}

func ApplyStyle(s string, styles ...string) string {
	return strings.Join(styles, "") + s + "\033[0m"
}

func PrintSetUpInstructions() {
	fmt.Println(ColoredString(ApplyStyle("Simulation Set Up", TextStyle.BOLD, TextStyle.REVERSE), Colors.FG_MAGENTA))
	fmt.Println(ApplyStyle("Commands:"))

	fmt.Println("\t- ", ColoredString(ApplyStyle("'a'", TextStyle.BOLD, TextStyle.UNDERLINE), Colors.FG_BRIGHT_MAGENTA), ": move to the cell to the left;")
	fmt.Println("\t- ", ColoredString(ApplyStyle("'w'", TextStyle.BOLD, TextStyle.UNDERLINE), Colors.FG_BRIGHT_MAGENTA), ": move to the cell above;")
	fmt.Println("\t- ", ColoredString(ApplyStyle("'s'", TextStyle.BOLD, TextStyle.UNDERLINE), Colors.FG_BRIGHT_MAGENTA), ": move to the cell below;")
	fmt.Println("\t- ", ColoredString(ApplyStyle("'d'", TextStyle.BOLD, TextStyle.UNDERLINE), Colors.FG_BRIGHT_MAGENTA), ": move to the cell to the right;")
	fmt.Println("\t- ", ColoredString(ApplyStyle("space bar", TextStyle.BOLD, TextStyle.UNDERLINE), Colors.FG_BRIGHT_MAGENTA), ": toggle cell status (alive/dead);")
	fmt.Println("\t- ", ColoredString(ApplyStyle("any other key", TextStyle.BOLD, TextStyle.UNDERLINE), Colors.FG_BRIGHT_MAGENTA), ": start simulation;")
	fmt.Println()
}
