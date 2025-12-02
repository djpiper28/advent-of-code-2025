package main

import (
	"fmt"
	"log"
	"os"
)

func isInvalid(id int64) bool {
	idStr := fmt.Sprintf("%d", id)
	if len(idStr)%2 == 1 {
		return false
	}

	lowerStr := idStr[0 : len(idStr)/2]
	upperStr := idStr[len(idStr)/2:]

	return lowerStr == upperStr
}

func addInvalidIds(ids [2]int64) int64 {
	var ret int64
	for id := ids[0]; id <= ids[1]; id++ {
		if isInvalid(id) {
			log.Printf("Invalid ID: %d", id)
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

	var num int64
	var part1Acc int64
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
			part1Acc += addInvalidIds(nums)

			// Reset the state
			nums = [2]int64{}
		}
	}

	log.Printf("Part one answer: %d", part1Acc)
}
