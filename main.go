package main

import (
	"flag"
	"log"
)

// TODO: General refactoring needed
// TODO: Clarify especially the structure for the printing of different sections

func main() {

	width := flag.Int("width", WIDTH, "Override the default width of the board")
	height := flag.Int("height", HEIGHT, "Override the default height of the board")
	fill := flag.Float64("fill", FILL, "How filled should the randomly generated matrix be? Value should be between 0 and 1")
	fps := flag.Int("fps", FPS, "Simulation frame rate (during set up the frame rate is set to a smooth 60fps)")
	flag.Parse()

	if *fill < 0 || *fill > 1 {
		log.Fatal("-fill value should be contained inside the [0, 1] interval")
	}

	game := NewGame(
		Config{
			width:  *width,
			height: *height,
			fill:   *fill,
			fps:    *fps,
		},
	)

	game.SetUp()
	game.Run()

}
