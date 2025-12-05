package main

import (
	"log"
	"os"
	"sort"
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

func (p *parser) Insert(rnge Range) {
	p.MergedRanges = append(p.MergedRanges, &rnge)
}

func (p *parser) Part2() int {
	for _, rnge := range p.Ranges {
		p.Insert(rnge)
	}

	sort.Slice(p.MergedRanges, func(i, j int) bool {
		return p.MergedRanges[i].Low < p.MergedRanges[j].Low
	})

	var mergedRanges []*Range
	if len(p.MergedRanges) > 0 {
		mergedRanges = append(mergedRanges, p.MergedRanges[0])
	}

	for i := 1; i < len(p.MergedRanges); i++ {
		lastMerged := mergedRanges[len(mergedRanges)-1]
		current := p.MergedRanges[i]

		if current.Low <= lastMerged.High {
			lastMerged.High = max(lastMerged.High, current.High) // Using built-in max
		} else {
			mergedRanges = append(mergedRanges, current)
		}
	}
	p.MergedRanges = mergedRanges

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
