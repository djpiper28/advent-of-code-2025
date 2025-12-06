package main

import (
	"log"
	"os"
	"time"
)

type Operator int

const (
	Operator_Addition Operator = iota + 1
	Operator_Multiplication
)

func (o Operator) Apply(a, b int) int {
	switch o {
	case Operator_Addition:
		return a + b
	case Operator_Multiplication:
		return a * b
	}

	log.Fatalf("Unknown operator: %v", o)
	return -1
}

func (p *parser) Part1() int {
	result := 0
	for i, operator := range p.Operators {
		calc := 0
		if operator == Operator_Multiplication {
			calc = 1
		}

		for _, row := range p.Numbers {
			calc = operator.Apply(calc, row[i])
		}

		result += calc
	}

	return result
}

func main() {
	txt, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatalf("Cannot read input: %s", err)
	}

	t := time.Now()

	p := &parser{
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
