package main

import (
	"fmt"
	"log"
	"os"

	"github.com/hymkor/go-cursorposition"
	tsize "github.com/kopoli/go-terminal-size"
	"golang.org/x/term"
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
	// _ := stripansi.String(strings.Join(strs, ""))

	// Switch terminal to raw-mode.
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		panic(err.Error())
	}

	// Query and display the cursor position with ESC[6n
	_, col, err := cursorposition.Request(os.Stderr)
	if err != nil {
		log.Fatal(err)
	} else {
		term.Restore(int(os.Stdin.Fd()), oldState)
	}

	for range s.Width - col + 1 {
		fmt.Print(ColoredString(" ", THEME.Primary, THEME.Background))
	}
	fmt.Printf("\n\r")
}
