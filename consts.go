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

const FPS = 24
const UI_REFRESH_RATE = time.Second / 60

const WIDTH int = 75
const HEIGHT int = 25

const SPARSITY int = 50
const MAX_POPULATION = (WIDTH * HEIGHT)

const ALIVE bool = true
const DEAD bool = false

var ALIVE_STRING string = AvailableSymbols.LETTERS.Filled
var DEAD_STRING string = AvailableSymbols.LETTERS.Empty
