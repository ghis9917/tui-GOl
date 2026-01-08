package main

type Cell struct {
	status   bool
	selected bool
}

func (c *Cell) String() string {
	representation := ""

	if c.status {
		representation = ALIVE_STRING
	} else {
		representation = DEAD_STRING
	}

	if c.selected {
		return ApplyStyle(
			ColoredString(representation, Colors.BRIGHT_CYAN),
			TextStyle.UNDERLINE,
		)
	} else {
		return ColoredString(representation, Colors.MAGENTA)
	}

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
