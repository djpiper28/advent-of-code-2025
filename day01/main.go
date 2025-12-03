package main

import (
	"log"
	"os"
	"time"
)

type data struct {
	inputData       []byte
	index           int
	dialPosition    int
	numZeros        int
	numZerosAnyTime int
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
		case 'L', 'l':
			left = true
		case 'R', 'r':
			left = false
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			num *= 10
			num += int(d.current() - '0')
    default:
      log.Fatalf("Cannot parse char %c", d.current())
		}
	}

	if d.current() == '\n' {
		d.index++
	}

	if left {
		for range num {
			d.dialPosition--
			if d.dialPosition == 0 {
				d.numZerosAnyTime++
			}

			if d.dialPosition < 0 {
				d.dialPosition = dialMax
			}
		}
	} else {
		for range num {
			d.dialPosition++

			if d.dialPosition > dialMax {
				d.dialPosition = 0
				d.numZerosAnyTime++
			}
		}
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

  t := time.Now()
	numZerosAnyTime := data{
		inputData:    inputData,
		dialPosition: 50,
	}

	for {
		done, err := numZerosAnyTime.readNextInstruction()
		if err != nil {
			log.Fatalf("Cannot parse file: %s, data: %v", err, numZerosAnyTime)
		}

		if done {
			break
		}
	}

	log.Printf("Dial Position: %d", numZerosAnyTime.dialPosition)
	log.Printf("Numer of zeros (part 1 output): %d", numZerosAnyTime.numZeros)
	log.Printf("Numer of zeros at any time (part 2 output): %d", numZerosAnyTime.numZerosAnyTime)
  log.Printf("Time taken: %s", time.Since(t))
}
