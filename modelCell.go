package main

import (
	"fmt"
)

type Cell struct {
	status   bool
	selected bool
}

func (c *Cell) String() string {
	representation := ""

	if c.status {
		representation = ColorString(ALIVE_STRING, Colors.MAGENTA)
	} else {
		representation = ColorString(DEAD_STRING, Colors.GREY)
	}

	if c.selected {
		representation = representation + ColorString("<", Colors.RED)
	} else {
		representation = fmt.Sprintf("%s", representation)
	}

	return representation
}

func (c *Cell) Fate(neighbours int) {
	if c.status == ALIVE {

		if neighbours < 2 || neighbours >= 4 {
			c.status = DEAD
		}

	} else {

		if neighbours == 3 {
			c.status = ALIVE
		}

	}
}
