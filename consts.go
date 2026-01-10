package main

import "time"

const BANNER = `
▄▄███▄▄·    ██╗       ██████╗  █████╗ ███╗   ███╗███████╗     ██████╗ ███████╗    ██╗     ██╗███████╗███████╗
██╔════╝    ╚██╗     ██╔════╝ ██╔══██╗████╗ ████║██╔════╝    ██╔═══██╗██╔════╝    ██║     ██║██╔════╝██╔════╝
███████╗     ╚██╗    ██║  ███╗███████║██╔████╔██║█████╗      ██║   ██║█████╗      ██║     ██║█████╗  █████╗  
╚════██║     ██╔╝    ██║   ██║██╔══██║██║╚██╔╝██║██╔══╝      ██║   ██║██╔══╝      ██║     ██║██╔══╝  ██╔══╝  
███████║    ██╔╝     ╚██████╔╝██║  ██║██║ ╚═╝ ██║███████╗    ╚██████╔╝██║         ███████╗██║██║     ███████╗
╚═▀▀▀══╝    ╚═╝       ╚═════╝ ╚═╝  ╚═╝╚═╝     ╚═╝╚══════╝     ╚═════╝ ╚═╝         ╚══════╝╚═╝╚═╝     ╚══════╝
                                                                                                                                                                                                                                   
`

const FPS = 30
const SIMULATION_FRAME_RATE = time.Second / FPS
const UI_FRAME_RATE = time.Second / 60

const WIDTH int = 100
const HEIGHT int = 25

const SPARSITY int = 10
const MAX_INITIAL_POPULATION = (WIDTH * HEIGHT) / SPARSITY

const ALIVE bool = true
const DEAD bool = false

var ALIVE_STRING string = AvailableSymbols.SQUARE_EMOJI.Filled
var DEAD_STRING string = AvailableSymbols.SQUARE_EMOJI.Empty
