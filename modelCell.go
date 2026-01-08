package main

type Cell struct {
	status   bool
	selected bool
}

func (c *Cell) String() string {
	representation := ""

	if c.status {
		representation = ColoredString(ALIVE_STRING, Colors.FG_MAGENTA) //ALIVE_STRING
	} else {
		representation = ColoredString(DEAD_STRING, Colors.FG_GREY) //DEAD_STRING
	}

	if c.selected {
		return ColoredString(representation, Colors.BG_GREY)
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
