package main

import (
	"log"
	"os"
	"time"
)

type Grid struct {
	Grid [][]bool
}

func (g *Grid) CountNeighbours(xPos, yPos int) int {
	sum := 0

	for y := yPos - 1; y <= yPos+1; y++ {
		for x := xPos - 1; x <= xPos+1; x++ {
			if x == xPos && y == yPos {
				continue
			} else if y < 0 || x < 0 {
				continue
			} else if y >= len(g.Grid) || x >= len(g.Grid[y]) {
				continue
			} else if g.Grid[y][x] {
				sum++
			}
		}
	}

	return sum
}

func (g *Grid) Part1() int {
	count := 0
	for y := 0; y < len(g.Grid); y++ {
		for x := 0; x < len(g.Grid[y]); x++ {
			if g.Grid[y][x] {
				if g.CountNeighbours(x, y) < 4 {
					count++
				}
			}
		}
	}

	return count
}

func main() {
	bytes, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatalf("Cannot read input: %s", err)
	}

	t := time.Now()

	// grid[y][x]
	var grid Grid
	var row []bool
	for i, c := range bytes {
		if c == '\n' || i == len(bytes)-1 {
			grid.Grid = append(grid.Grid, row)
			row = make([]bool, 0)
		} else if c == '@' {
			row = append(row, true)
		} else if c == '.' {
			row = append(row, false)
		} else {
			log.Fatalf("Invalid char '%c'", c)
		}
	}

	log.Printf("Part one answer: %d", grid.Part1())
	log.Printf("Time taken: %s", time.Since(t))
}
