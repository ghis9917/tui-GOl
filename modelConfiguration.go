package main

import (
	"fmt"
	"log"
	"strings"

	tsize "github.com/kopoli/go-terminal-size"
	"github.com/leaanthony/go-ansi-parser"
)

type Config struct {
	width  int
	height int
	fps    int
	fill   float64
	output string
}

func (cfg *Config) Draw() {
	ClearScreen()
	fmt.Print(cfg.output)
}

func (cfg *Config) ResetOutput() {
	cfg.output = ""
}

func (cfg *Config) Println(strs ...string) {

	lineStart := "  "
	cfg.output += fmt.Sprint(ColoredString(lineStart, THEME.Primary, THEME.Background))

	for _, s := range strs {
		cfg.output += fmt.Sprint(ColoredString(s, THEME.Primary, THEME.Background))
	}

	s, err := tsize.GetSize()
	if err != nil {
		log.Fatal(err)
	}

	count, err := ansi.Length(lineStart + strings.Join(strs, ""))
	if err != nil {
		log.Fatal(err)
	}

	for range s.Width - count {
		cfg.output += fmt.Sprint(ColoredString(" ", THEME.Primary, THEME.Background))
	}
	cfg.output += fmt.Sprintln()
}

func (cfg *Config) FillVerticalSpace() {

	s, err := tsize.GetSize()
	if err != nil {
		log.Fatal(err)
	}

	countRows := countRune(cfg.output, '\n')

	for range s.Height - countRows - 1 {

		cfg.Println()

	}

}

func (cfg *Config) PrintConfig() {

	cfg.Println(
		ColoredString(
			ApplyStyle(
				"Configuration",
				TextStyle.BOLD,
				TextStyle.REVERSE,
			),
			THEME.Secondary,
			THEME.Background,
		),
	)

	cfg.Println(
		ApplyStyle("* Board Size (W x H): ", TextStyle.BOLD),
		ColoredString(
			ApplyStyle(
				fmt.Sprintf("%v x %v", cfg.width, cfg.height),
				TextStyle.BOLD,
				TextStyle.UNDERLINE,
			),
			THEME.Secondary,
			THEME.Background,
		),
	)

	cfg.Println(
		ApplyStyle("* Initial Fill: ", TextStyle.BOLD),
		ColoredString(
			ApplyStyle(
				fmt.Sprintf("%v", cfg.fill*100)+"%",
				TextStyle.BOLD,
				TextStyle.UNDERLINE,
			),
			THEME.Secondary,
			THEME.Background,
		),
	)

	cfg.Println(
		ApplyStyle("* Frame Rate: ", TextStyle.BOLD),
		ColoredString(
			ApplyStyle(
				fmt.Sprintf("%v frames/s", cfg.fps),
				TextStyle.BOLD,
				TextStyle.UNDERLINE,
			),
			THEME.Secondary,
			THEME.Background,
		),
	)

	cfg.Println(
		ApplyStyle("* Theme: ", TextStyle.BOLD),
		ColoredString(
			ApplyStyle(
				fmt.Sprintf("%s", THEME.Name),
				TextStyle.BOLD,
				TextStyle.UNDERLINE,
			),
			THEME.Secondary,
			THEME.Background,
		),
	)

	cfg.Println()

}
