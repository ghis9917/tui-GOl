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
}

func (cfg *Config) Println(strs ...string) {
	for _, s := range strs {
		fmt.Print(ColoredString(s, THEME.Primary, THEME.Background))
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
		fmt.Print(ColoredString(" ", THEME.Primary, THEME.Background))
	}
	fmt.Println()
}
