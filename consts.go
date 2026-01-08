package main

import "time"

const FPS = 30
const SIMULATION_FRAME_RATE = time.Second / FPS
const UI_FRAME_RATE = time.Second / 60

const WIDTH int = 100
const HEIGHT int = 25

const SPARSITY int = 10
const MAX_INITIAL_POPULATION = (WIDTH * HEIGHT) / SPARSITY

const ALIVE bool = true
const ALIVE_STRING string = "●"
const DEAD bool = false
const DEAD_STRING string = "○"

type Color struct {
	RESET   string
	GREY    string
	MAGENTA string
	RED     string
}

var Colors = Color{
	RESET:   "\033[0m",
	RED:     "\033[31m",
	MAGENTA: "\033[35m",
	GREY:    "\033[90m",
}
