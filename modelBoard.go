package main

import (
	"math/rand"
)

type Board struct {
	cells        [][]Cell
	height       int
	width        int
	selectedCell Coordinate
}

type Stats struct {
	generation       int
	population       []float64
	populationChange []float64
}

func NewBoard(cfg Config) (board *Board) {

	b := &Board{
		cells:        make([][]Cell, cfg.height),
		height:       cfg.height,
		width:        cfg.width,
		selectedCell: Coordinate{i: 0, j: 0},
	}
	limit := int(float64(cfg.width*cfg.height) * cfg.fill)

	for i := range b.cells {
		b.cells[i] = make([]Cell, cfg.width)
		for j := range b.cells[i] {
			if rand.Intn(int(1.0/cfg.fill)) == 0 && limit > 0 {
				b.cells[i][j] = Cell{status: ALIVE, selected: Coordinate{i: i, j: j} == b.selectedCell}
				limit -= 1
			} else {
				b.cells[i][j] = Cell{status: DEAD, selected: Coordinate{i: i, j: j} == b.selectedCell}
			}
		}
	}

	return b

}

func (b *Board) SwitchCellStatus() {

	b.cells[b.selectedCell.i][b.selectedCell.j].status = !b.cells[b.selectedCell.i][b.selectedCell.j].status

}

func (b *Board) SwitchCellSelection() {

	b.cells[b.selectedCell.i][b.selectedCell.j].selected = !b.cells[b.selectedCell.i][b.selectedCell.j].selected

}

func (b *Board) CleanSelection() {

	for i := range b.cells {
		for j := range b.cells[i] {

			b.cells[i][j].selected = false

		}
	}

}

func (b *Board) PolliceVerso() {

	neighboursMap := make([][]int, b.height)

	for i := range b.cells {
		neighboursMap[i] = make([]int, b.width)
		for j := range b.cells[i] {

			neighbours := CountLiveNeighbours(i, j, b)
			neighboursMap[i][j] = neighbours

		}
	}

	for i := range b.height {
		for j := range b.width {

			b.cells[i][j].Fate(neighboursMap[i][j])

		}
	}

}

func (b *Board) GetCell(i, j int) (status, valid bool) {

	iValid := i >= 0 && i < len(b.cells)
	jValid := j >= 0 && j < len(b.cells[0])

	if iValid && jValid {
		return b.cells[i][j].status, true
	}

	return false, false

}

func CountLiveNeighbours(i, j int, b *Board) (neighboursNumber int) {

	count := 0

	directions := []Coordinate{
		{i: -1, j: -1}, {i: -1, j: +0}, {i: -1, j: +1}, // 1 ROW ABOVE
		{i: +0, j: -1}, {i: +0, j: +1}, // SAME ROW
		{i: +1, j: -1}, {i: +1, j: +0}, {i: +1, j: +1}, // 1 ROW DOWN
	}

	for _, c := range directions {

		if status, valid := b.GetCell(i+c.i, j+c.j); valid {
			if status {
				count += 1
			}
		}

	}

	return count
}

func (b *Board) Strings() (rows []string) {

	rows = []string{}

	for i := range b.cells {
		representation := ""
		for j := range b.cells[i] {
			representation += b.cells[i][j].String()
		}
		rows = append(rows, representation)
	}

	return rows

}

func (b *Board) GetPopulation() (population float64) {

	population = 0

	for i := range b.cells {
		for j := range b.cells[i] {
			if b.cells[i][j].status {
				population += 1
			}
		}
	}

	return float64(population)
}

func (b *Board) Move(axys *int, direction, upperLimit int) {

	if *axys+direction >= 0 && *axys+direction < upperLimit {
		b.SwitchCellSelection()
		*axys += direction
		b.SwitchCellSelection()
	}

}
