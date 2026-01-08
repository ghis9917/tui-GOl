package main

import "fmt"

type Coordinate struct {
	i int
	j int
}

func (c *Coordinate) String() string {
	return fmt.Sprintf("(i: %d, j: %d)", c.i, c.j)
}
