package main

import (
	"fmt"
	"time"

	"github.com/eiannone/keyboard"
)

func main() {

	// TODO: Set up initialization of simulation through args such as:
	// TODO:     * size of the simulation board (width, height)
	// TODO:     * sparsity of the random initialization of cells (how many can randomly be initialized as alive?)
	// TODO:     * FPS
	// TODO:     * Representation of cells (alive and dead symbols)
	// TODO:     * Text colors

	// TODO: Write function to print 'help' cmd with all available options

	// TODO: Create theme struct to simplify editing of color

	board := NewBoard(
		WIDTH,
		HEIGHT,
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
				panic(err)
			}
			board.HandleKeyStrokes(keysEvents)
		}

		if board.started {
			time.Sleep(SIMULATION_FRAME_RATE)
		} else {
			time.Sleep(UI_FRAME_RATE)
		}
	}

}
