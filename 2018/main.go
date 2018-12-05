package main

import (
	"adventofcode/2018/day_01"
	"adventofcode/2018/day_02"
	"adventofcode/2018/shared"
	"flag"
	"fmt"
)

func main() {

	day := flag.Int("day", 1, "puzzle by day 1-25")
	part := flag.Int("part", 1, "part of the puzzle for the day 1/2")
	flag.Parse()

	inputPath := shared.GetInputPath(*day)
	lines := shared.ReadInputLines(inputPath)

	var result interface{}
	switch *day {
	case 1:
		switch *part {
		case 1:
			result = day_01.CalculateFrequencyDrift(lines)
		case 2:
			result = day_01.FindFirstFrequencyRepeat(lines)
		}
	case 2:
		switch *part {
		case 1:
			result = day_02.CalculateInventoryChecksum(lines)
		case 2:
			result = day_02.FindPrototypeFabricBoxes(lines)
		}
	default:
		panic("Requested day not yet implemented.")
	}

	fmt.Printf("Day %v Part %v Result: %v\n", *day, *part, result)
}
