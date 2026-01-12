package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/eiannone/keyboard"
)

func main() {

	width := flag.Int("width", WIDTH, "Override the default width of the board")
	height := flag.Int("height", HEIGHT, "Override the default height of the board")
	fill := flag.Float64("fill", FILL, "How filled should the randomly generated matrix be? Value should be between 0 and 1")
	fps := flag.Int("fps", FPS, "Simulation frame rate (during set up the frame rate is set to a smooth 60fps)")
	flag.Parse()

	simulationRefreshRate := time.Second / time.Duration(*fps)

	if *fill < 0 || *fill > 1 {
		log.Fatal("-fill value should be contained inside the [0, 1] interval")
	}

	// TODO: Set up initialization of simulation through args such as:
	// TODO:     * Representation of cells (alive and dead symbols)
	// TODO:     * Text colors

	// TODO: Create theme struct to simplify editing of color

	board := NewBoard(
		*width,
		*height,
		*fill,
	)

	for {

		ClearScreen()
		fmt.Print(ColoredString(BANNER, THEME.Primary, THEME.Background))

		if board.started {
			if board.Evolve() {
				break
			}
		} else {
			PrintSetUpInstructions()
			board.PrintBoard()
			keysEvents, err := keyboard.GetKeys(10)
			if err != nil {
				log.Fatal(err)
			}
			board.HandleKeyStrokes(keysEvents)
		}

		if board.started {
			time.Sleep(simulationRefreshRate)
		} else {
			time.Sleep(UI_REFRESH_RATE)
		}
	}

	ClearScreen()
	fmt.Print(ColoredString(END_BANNER, THEME.Primary, THEME.Background))
	board.PrintSummary()

}
