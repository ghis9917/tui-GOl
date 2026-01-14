package main

import (
	"flag"
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

	board := NewBoard(
		Config{
			width:  *width,
			height: *height,
			fill:   *fill,
			fps:    *fps,
		},
	)

	for {

		board.cfg.ResetOutput()
		board.PrintBanner(BANNER)

		if board.started {
			if board.Evolve() {
				board.cfg.Draw()
				break
			}
			ClearScreen()
			board.cfg.Draw()
		} else {
			board.PrintSetUpInstructions()
			board.PrintBoard()
			board.cfg.Draw()
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

	board.cfg.ResetOutput()
	board.PrintBanner(END_BANNER)
	board.PrintSummary()
	board.cfg.Draw()

}
