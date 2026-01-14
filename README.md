# TUI Game Of Life

The Game of Life, also known simply as *Life*, is a cellular automaton. Based on simple rules, cells can live, die, or revive their neighbors.
![simulation](https://github.com/ghis9917/tui-GOl/blob/main/resources/SampleVideo.mp4)

## Rules

Cells in *Life* are elements of $\mathbb{Z}^2$, and their *neighbors* are the eight adjacent points, horizontally, vertically, or diagonally. Each cell can be either *alive* or *dead*. At each step in time, the state of a cell can change based on the following rules:


- Any live cell with fewer than two live neighbours dies, as if by underpopulation.
- Any live cell with two or three live neighbours lives on to the next generation.
- Any live cell with more than three live neighbours dies, as if by overpopulation.
- Any dead cell with exactly three live neighbours becomes a live cell, as if by reproduction.

Generation 0 is known as the *seed* of the system, and each new generation is created by applying these rules to the previous generation. As new generations arise, different types of patterns may appear. For further information, you can refer to [this article](https://en.wikipedia.org/wiki/Conway%27s_Game_of_Life).

## Getting Started

### Prerequisites

- Go (if you want to rebuild for your own platform, code was compiled on MacOS)

### Installation

1. Clone the repository:
```
git clone https://github.com/ghis9917/tui-GOl
```
2. Grant run permission to 'tui-GOl':
```
chmod +x tui-GOl
```
3. Execute 'tui-GOl' to start:
```
./tui-GOl
```

### Controls during Set Up:
- *a* to move the selector to the left
- *d* to move the selector to the right
- *w* to move the selector up
- *s* to move the selector down
- *space* to switch cell status
- *any other key* to start simulation