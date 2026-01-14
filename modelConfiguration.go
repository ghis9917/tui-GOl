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
	fmt.Print(cfg.output)
}

func (cfg *Config) ResetOutput() {
	cfg.output = ""
}

func (cfg *Config) Println(strs ...string) {
	for _, s := range strs {
		cfg.output += fmt.Sprint(ColoredString(s, THEME.Primary, THEME.Background))
	}

	s, err := tsize.GetSize()
	if err != nil {
		log.Fatal(err)
	}

	count, err := ansi.Length(strings.Join(strs, ""))
	if err != nil {
		log.Fatal(err)
	}

	for range s.Width - count {
		cfg.output += fmt.Sprint(ColoredString(" ", THEME.Primary, THEME.Background))
	}
	cfg.output += fmt.Sprintln()
}
