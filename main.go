package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

type Cells uint

const (
	Dead Cells = iota
	Alive
)

const (
	HEIGHT = 6
	WEIGHT = 6
)

type (
	// Board
	Board [HEIGHT][WEIGHT]Cells

	// Coords
	Coords struct {
		Row int
		Col int
	}

	// Range
	Range struct {
		start int
		end   int
	}

	// Neighbours
	Neighbours []Cells
)

func (ns Neighbours) Has(cell Cells) (r int) {
	for _, n := range ns {
		if n == cell {
			r++
		}
	}
	return
}

func getRange(board Board, index int) (r Range) {
	switch index {
	case 0:
		r.start, r.end = 0, 1
	case len(board) - 1:
		r.start, r.end = index-1, index
	default:
		r.start, r.end = index-1, index+1
	}
	return
}

func getNeighbours(board Board, coords Coords) (neighbours Neighbours) {

	rowRange := getRange(board, coords.Row)
	colRange := getRange(board, coords.Col)

	//fmt.Println("DEBUG", rowRange, colRange)

	for i := range board {
		for j := range board[i] {
			if i >= rowRange.start && i <= rowRange.end &&
				j >= colRange.start && j <= colRange.end {

				if i != coords.Row && j != coords.Col {
					neighbours = append(neighbours, board[i][j])
				} else if i == coords.Row && j != coords.Col {
					neighbours = append(neighbours, board[i][j])
				} else if i != coords.Row && j == coords.Col {
					neighbours = append(neighbours, board[i][j])
				}

			}
		}
	}
	return
}

func checkRules(currentCell Cells, neighbours Neighbours) (cell Cells) {
	AliveNeighbours := neighbours.Has(Alive)

	switch currentCell {
	case Alive:
		if AliveNeighbours == 2 || AliveNeighbours == 3 {
			cell = Alive
		}
	case Dead:
		if AliveNeighbours == 3 {
			cell = Alive
		}
	}
	return
}

func nexGen(board Board) (next Board) {
	// TODO: ...
	// if alive cell has 2 || 3 live neighbours => alive
	// if dead cell has 3 alive neighbours => alive
	// 3 in next gen any live cells becomes dead and dead cells stays dead.
	//  in 3 i guess what they mean with that is any other cells without neighbour
	// that is alive should be dead in next gen and dead ones stays dead.
	for i := range board {
		for j := range board[i] {
			neighbours := getNeighbours(board, Coords{i, j})
			next[i][j] = checkRules(board[i][j], neighbours)
		}
	}
	return next
}

func renderBoard(board Board) {
	for i := range board {
		for _, c := range board[i] {
			switch c {
			case Alive:
				fmt.Printf("*")
			case Dead:
				fmt.Printf(".")
			}
		}
		fmt.Println()
	}
}

func clearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {

	board := Board{
		{Dead, Dead, Dead, Dead, Dead},
		{Dead, Dead, Alive, Dead, Dead},
		{Dead, Dead, Dead, Alive, Dead},
		{Dead, Alive, Alive, Alive, Dead},
		{Dead, Dead, Dead, Dead, Dead},
		{Dead, Dead, Dead, Dead, Dead},
	}

	for i := 0; i < 13; i++ {
		clearScreen()
		renderBoard(board)
		board = nexGen(board)
		time.Sleep(1 * time.Second)
	}
}
