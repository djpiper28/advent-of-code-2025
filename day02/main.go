package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func isInvalidPart1(id int64) bool {
	idStr := fmt.Sprintf("%d", id)
	if len(idStr)%2 == 1 {
		return false
	}

	lowerStr := idStr[0 : len(idStr)/2]
	upperStr := idStr[len(idStr)/2:]

	return lowerStr == upperStr
}

func addInvalidIdsPart1(ids [2]int64) int64 {
	var ret int64
	for id := ids[0]; id <= ids[1]; id++ {
		if isInvalidPart1(id) {
			log.Printf("Invalid ID: %d", id)
			ret += id
		}
	}

	return ret
}

func isInvalidPart2(id int64) bool {
	idStr := fmt.Sprintf("%d", id)

	// Special case factor of 1
	eq := true
	for i := 1; i < len(idStr); i++ {
		if idStr[i] != idStr[0] {
			eq = false
			break
		}
	}

	if eq {
		log.Printf("Invalid Id Factor: 1, Id: %d", id)
		return true
	}

	// Slower checks for larger symmetries
	for factor := 2; factor <= len(idStr)/2; factor++ {
		if len(idStr)%factor != 0 {
			continue
		}

		parts := make([]string, factor)
		partLen := len(idStr) / factor

		for i := range factor {
			parts[i] = idStr[i*int(partLen) : (i+1)*partLen]
		}

		eq := true
		for i := 1; i < len(parts); i++ {
			if parts[i] != parts[0] {
				eq = false
				break
			}
		}

		if eq {
			log.Printf("Invalid Id Factor: %d, Part: %s, Id: %d", factor, parts[0], id)
			return true
		}
	}

	return false
}

func addInvalidIdsPart2(ids [2]int64) int64 {
	var ret int64
	for id := ids[0]; id <= ids[1]; id++ {
		if isInvalidPart2(id) {
			ret += id
		}
	}

	return ret
}

func main() {
	bytes, err := os.ReadFile("./input.txt")
	if err != nil {
		log.Fatalf("Cannot read file: %s", err)
	}

	t := time.Now()
	var num int64
	var part1Acc int64
	var part2Acc int64
	nums := [2]int64{}
	for i, c := range bytes {
		switch {
		case ('0' <= c && c <= '9'):
			num *= 10
			num += int64(c - '0')
		case c == '-':
			nums[0] = num
			num = 0
		case i == len(bytes)-1:
			fallthrough
		case c == ',':
			nums[1] = num
			num = 0
			part1Acc += addInvalidIdsPart1(nums)
			part2Acc += addInvalidIdsPart2(nums)

			// Reset the state
			nums = [2]int64{}
		}
	}

	log.Printf("Time: %s", time.Since(t))
	log.Printf("Part one answer: %d", part1Acc)
	log.Printf("Part two answer: %d", part2Acc)
}
