package main

import (
	"log"
	"os"
	"time"
)

type Grid struct {
	Grid      [][]bool
	CountGrid [][]int
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

func (g *Grid) Part2() int {
	count := 0
	dirty := true
	for dirty {
		dirty = false

		for y := 0; y < len(g.Grid); y++ {
			for x := 0; x < len(g.Grid[y]); x++ {
				if g.Grid[y][x] {
					if g.CountNeighbours(x, y) < 4 {
						g.Grid[y][x] = false
						count++
						dirty = true
					}
				}
			}
		}
	}

	return count
}

func (g *Grid) FastPart2() int {
	g.CountGrid = make([][]int, len(g.Grid))

	for y := 0; y < len(g.Grid); y++ {
		g.CountGrid[y] = make([]int, len(g.Grid[y]))

		for x := 0; x < len(g.Grid[y]); x++ {
			if g.Grid[y][x] {
				g.CountGrid[y][x] = g.CountNeighbours(x, y)
			}
		}
	}

	count := 0
	dirty := true

	for dirty {
		dirty = false

		for yPos := 0; yPos < len(g.CountGrid); yPos++ {
			for xPos := 0; xPos < len(g.CountGrid[yPos]); xPos++ {
				if !g.Grid[yPos][xPos] {
					continue
				}

				if g.CountGrid[yPos][xPos] < 4 {
					g.Grid[yPos][xPos] = false
					count++
					dirty = true

					// All of the neighbours have one less now
					for y := yPos - 1; y <= yPos+1; y++ {
						for x := xPos - 1; x <= xPos+1; x++ {
							if y < 0 || x < 0 {
								continue
							} else if y >= len(g.Grid) || x >= len(g.Grid[y]) {
								continue
							}

							g.CountGrid[y][x]--
						}
					}
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
	// log.Printf("Part two answer: %d", grid.Part2())
	log.Printf("Part two (fast) answer: %d", grid.FastPart2())
	log.Printf("Time taken: %s", time.Since(t))
}
