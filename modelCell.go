package main

type Cell struct {
	status   bool
	selected bool
}

func (c *Cell) String() string {
	representation := ""

	if c.status {
		representation = ColoredString(ALIVE_STRING, Colors.FG_MAGENTA) //Applied only to unicode symbols
	} else {
		representation = ColoredString(DEAD_STRING, Colors.FG_GREY) //Applied only to unicode symbols
	}

	if c.selected {
		return ColoredString(representation, Colors.BG_BRIGHT_CYAN) //Change accto keep contrast with emoji/unicode color
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
