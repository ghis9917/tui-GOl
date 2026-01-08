package main

import (
	"time"

	"github.com/eiannone/keyboard"
)

func main() {

	board := NewBoard(
		WIDTH,
		HEIGHT,
	)

	for {

		ClearScreen()

		if board.started {
			if board.Evolve() {
				break
			}
		} else {
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
