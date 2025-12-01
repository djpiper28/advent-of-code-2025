package main

import (
	"log"
	"os"
)

type data struct {
	inputData    []byte
	index        int
	dialPosition int
	numZeros     int
}

const dialMax = 99

func (d *data) current() byte {
	return d.inputData[d.index]
}

func (d *data) readNextInstruction() (bool, error) {
	if d.index >= len(d.inputData) {
		return true, nil
	}

	left := false
	num := 0

	for ; d.index < len(d.inputData) && d.current() != '\n'; d.index++ {
		switch d.current() {
		case 'l':
			fallthrough
		case 'L':
			left = true
		case 'r':
			fallthrough
		case 'R':
			left = false
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			num *= 10
			num += int(d.current() - '0')
		}
	}

	if d.current() == '\n' {
		d.index++
	}

	if left {
		d.dialPosition -= num
	} else {
		d.dialPosition += num
	}

	for d.dialPosition < 0 {
		d.dialPosition += dialMax + 1
	}

	for d.dialPosition > dialMax {
		d.dialPosition -= dialMax + 1
	}

	if d.dialPosition == 0 {
		d.numZeros++
	}

	return false, nil
}

func main() {
	inputData, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatalf("Cannot read file: %s", err)
	}

	data := data{
		inputData:    inputData,
		dialPosition: 50,
	}

	for {
		done, err := data.readNextInstruction()
		if err != nil {
			log.Fatalf("Cannot parse file: %s, data: %v", err, data)
		}

		if done {
			break
		}
	}

	log.Printf("Dial Position: %d", data.dialPosition)
	log.Printf("Numer of zeros (output): %d", data.numZeros)
}
