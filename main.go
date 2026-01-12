package main

import (
	"fmt"
	"time"
	"flag"
	"log"

	"github.com/eiannone/keyboard"
)

func main() {

	width := flag.Int("width", WIDTH, "Override the default width of the board")
	height := flag.Int("height", HEIGHT, "Override the default height of the board")
	sparsity := flag.Int("sparsity", SPARSITY, "How filled should the randomly generated matrix be?")
	fps := flag.Int("fps", FPS, "Simulation frame rate (during set up the frame rate is set to a smooth 60fps)")
	flag.Parse()

	simulationRefreshRate := time.Second / time.Duration(*fps)

	// TODO: Set up initialization of simulation through args such as:
	// TODO:     * Representation of cells (alive and dead symbols)
	// TODO:     * Text colors
	// TODO:     * Max Generations

	// TODO: Create theme struct to simplify editing of color

	board := NewBoard(
		*width,
		*height,
		*sparsity,
	)

	for {

		ClearScreen()
		fmt.Print(ColoredString(BANNER, Colors.FG_BRIGHT_CYAN))

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

}
