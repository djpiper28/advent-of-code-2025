package main

import (
	"log"
	"os"
	"time"
)

type Bank struct {
	Cells []int
}

func (b *Bank) MaximumJoltage() int {
	left := 0
	leftIdx := 0
	for i, joltage := range b.Cells {
		if i == len(b.Cells)-1 {
			break
		}

		if joltage > left {
			left = joltage
			leftIdx = i
		}

		if joltage >= 9 {
			break
		}
	}

	right := 0
	for _, joltage := range b.Cells[leftIdx+1:] {
		if joltage > right {
			right = joltage
		}

		if joltage >= 9 {
			break
		}
	}

	return left*10 + right
}

func main() {
	bytes, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatalf("Cannot read input: %s", err)
	}

	t := time.Now()
	banks := make([]Bank, 0)
	bank := Bank{}

	for i, b := range bytes {
		if '0' <= b && b <= '9' {
			num := int(b - '0')
			bank.Cells = append(bank.Cells, num)
		} else if i == len(bytes)-1 || b == '\n' {
			banks = append(banks, bank)
			bank = Bank{}
		}
	}

	sum := 0
	for _, bank := range banks {
		val := bank.MaximumJoltage()
		log.Printf("Found %d", val)
		sum += val
	}

  log.Printf("Part one answer: %d", sum)
	log.Printf("Time taken: %s", time.Since(t))
}
