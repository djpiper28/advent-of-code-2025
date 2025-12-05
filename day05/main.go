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

func (p *parser) Insert(rnge *Range) {
	for i, mergedRange := range p.MergedRanges {
		if rnge.Low >= mergedRange.Low && rnge.Low <= mergedRange.High {
			/*
			        |--------------| a
			   |--------------|      merged
			*/
		} else if rnge.High <= mergedRange.High && rnge.High >= mergedRange.Low {
			/*
			   |--------------|     a
			       |--------------| merged
			*/
		} else {
			continue
		}

		mergedRange.Low = min(mergedRange.Low, rnge.Low)
		mergedRange.High = max(mergedRange.High, rnge.High)

		// Check for new overlapping range
		p.MergedRanges = p.MergedRanges[0:i]
		if i < len(p.MergedRanges) {
			p.MergedRanges = append(p.MergedRanges, p.MergedRanges[i+1:]...)
		}

		p.Insert(mergedRange)
		return
	}

	p.MergedRanges = append(p.MergedRanges, rnge)
}

func (p *parser) Part2() int {
	for _, rnge := range p.Ranges {
		p.Insert(&rnge)
	}

	sum := 0
	for _, rnge := range p.MergedRanges {
		log.Printf("Merged range: %v, ids: %d", *rnge, rnge.High-rnge.Low+1)
		sum += rnge.High - rnge.Low + 1
	}
	return sum
}

func main() {
	txt, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatalf("Cannot read input: %s", err)
	}

	t := time.Now()

	p := &parser{
		Items:        make([]int, 0),
		Ranges:       make([]Range, 0),
		MergedRanges: make([]*Range, 0),
		Buffer:       string(txt),
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
	log.Printf("Part two answer: %d", p.Part2())
	log.Printf("Time taken: %s", time.Since(t))
}
