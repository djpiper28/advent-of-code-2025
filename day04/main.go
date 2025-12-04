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
	// xPos, yPos is always iterated over
	sum := -1

	yMax := min(yPos+1, len(g.Grid)-1)
	xMax := min(xPos+1, len(g.Grid[0])-1)
	for y := max(0, yPos-1); y <= yMax; y++ {
		for x := max(0, xPos-1); x <= xMax; x++ {
			if g.Grid[y][x] {
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

func (g *Grid) fastPart2R(xPos, yPos int) int {
	if g.CountGrid[yPos][xPos] >= 4 {
		return 0
	}

	sum := 1
	g.Grid[yPos][xPos] = false

	// All of the neighbours have one less now
	yMax := min(yPos+1, len(g.Grid)-1)
	xMax := min(xPos+1, len(g.Grid[0])-1)
	for y := max(0, yPos-1); y <= yMax; y++ {
		for x := max(0, xPos-1); x <= xMax; x++ {
			g.CountGrid[y][x]--
		}
	}

	for y := max(0, yPos-1); y <= yMax; y++ {
		for x := max(0, xPos-1); x <= xMax; x++ {
			if g.Grid[y][x] {
				sum += g.fastPart2R(x, y)
			}
		}
	}
	return sum
}

func (g *Grid) FastPart2() int {
	g.CountGrid = make([][]int, len(g.Grid))

	yLen := len(g.Grid)
	xLen := len(g.Grid[0])
	for y := range yLen {
		g.CountGrid[y] = make([]int, len(g.Grid[y]))

		for x := range xLen {
			if g.Grid[y][x] {
				g.CountGrid[y][x] = g.CountNeighbours(x, y)
			}
		}
	}

	sum := 0

	for y := range yLen {
		for x := range xLen {
			if g.Grid[y][x] {
				sum += g.fastPart2R(x, y)
			}
		}
	}

	return sum
}

func (grid *Grid) parse(bytes []byte) {
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
}

func main() {
	bytes, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatalf("Cannot read input: %s", err)
	}

	t := time.Now()

	// grid[y][x]
	var grid Grid
	grid.parse(bytes)

	log.Printf("Part one answer: %d", grid.Part1())
	// log.Printf("Part two answer: %d", grid.Part2())
	log.Printf("Part two (fast) answer: %d", grid.FastPart2())
	log.Printf("Time taken: %s", time.Since(t))

	const tries = 1000
	log.Printf("Running fast part 2 %d times", tries)
	t = time.Now()

	for range tries {
		var grid Grid
		grid.parse(bytes)

		x := grid.FastPart2()
		if x != 8899 {
			log.Fatalf("Broken: %d", x)
		}
	}

	log.Printf("Time taken: %s", time.Since(t)/tries)

	log.Printf("Running part 2 %d times", tries)
	t = time.Now()

	for range tries {
		var grid Grid
		grid.parse(bytes)

		x := grid.Part2()
		if x != 8899 {
			log.Fatalf("Broken: %d", x)
		}
	}

	log.Printf("Time taken: %s", time.Since(t)/tries)
}
