package main

import (
	"time"

	tm "github.com/buger/goterm"
	"github.com/eiannone/keyboard"
)

func main() {

	board := NewBoard(
		WIDTH,
		HEIGHT,
	)

	tm.Clear()

	for {

		tm.MoveCursor(1, 1)

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

		tm.Flush()

		if board.started {
			time.Sleep(SIMULATION_FRAME_RATE)
		} else {
			time.Sleep(UI_FRAME_RATE)
		}
	}

}
