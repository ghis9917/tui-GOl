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

func PrintSetUpInstructions() {
	Println(ColoredString(ApplyStyle("Simulation Set Up", TextStyle.BOLD, TextStyle.REVERSE), THEME.Secondary, THEME.Background))
	Println(ApplyStyle("Commands:"))

	Println("    - ", ColoredString(ApplyStyle("'a'", TextStyle.BOLD, TextStyle.UNDERLINE), THEME.Secondary, THEME.Background), ": move to the cell to the left;")
	Println("    - ", ColoredString(ApplyStyle("'w'", TextStyle.BOLD, TextStyle.UNDERLINE), THEME.Secondary, THEME.Background), ": move to the cell above;")
	Println("    - ", ColoredString(ApplyStyle("'s'", TextStyle.BOLD, TextStyle.UNDERLINE), THEME.Secondary, THEME.Background), ": move to the cell below;")
	Println("    - ", ColoredString(ApplyStyle("'d'", TextStyle.BOLD, TextStyle.UNDERLINE), THEME.Secondary, THEME.Background), ": move to the cell to the right;")
	Println("    - ", ColoredString(ApplyStyle("space bar", TextStyle.BOLD, TextStyle.UNDERLINE), THEME.Secondary, THEME.Background), ": toggle cell status (alive/dead);")
	Println("    - ", ColoredString(ApplyStyle("any other key", TextStyle.BOLD, TextStyle.UNDERLINE), THEME.Secondary, THEME.Background), ": start simulation;")
	Println("")
}
