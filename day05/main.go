package main

import (
	"log"
	"os"
	"time"
)

type Range struct {
	Low, High int
}

func (p *parser) Part1() int {
	count := 0

	for _, item := range p.Items {
		for _, rnge := range p.Ranges {
			if item >= rnge.Low && item <= rnge.High {
				count++
				break
			}
		}
	}

	return count
}

func main() {
	txt, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatalf("Cannot read input: %s", err)
	}

	t := time.Now()

	p := &parser{
		Items:  make([]int, 0),
		Ranges: make([]Range, 0),
		Buffer: string(txt),
	}

	err = p.Init()
	if err != nil {
		log.Fatalf("Cannot init parser: %s", err)
	}

	err = p.Parse()
	if err != nil {
		log.Fatalf("Cannot parse input: %s", err)
	}

	p.Execute()

	log.Printf("Part one answer: %d", p.Part1())
	log.Printf("Time taken: %s", time.Since(t))
}
