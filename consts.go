package main

import "time"

const FPS = 30
const SIMULATION_FRAME_RATE = time.Second / FPS
const UI_FRAME_RATE = time.Second / 60

const WIDTH int = 200
const HEIGHT int = 50

const SPARSITY int = 5
const MAX_INITIAL_POPULATION = (WIDTH * HEIGHT) / SPARSITY

const ALIVE bool = true
const ALIVE_STRING string = "●"
const DEAD bool = false
const DEAD_STRING string = "○"
