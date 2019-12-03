package main

import (
	"flag"
	"fmt"

	"github.com/mikeder/adventofcode/2019/internal/day1"
	"github.com/mikeder/adventofcode/2019/internal/day2"
)

func main() {
	day := flag.Int("day", 1, "puzzle by day 1-25")
	part := flag.Int("part", 1, "part of the puzzle for the day 1/2")
	flag.Parse()

	if *part < 1 || *part > 2 {
		panic(fmt.Sprintf("Invalid part %d", *part))
	}

	fmt.Printf("Solving for Day %d Problem %d\n", *day, *part)
	answer := solve(*day, *part)
	fmt.Println("Answer: ", answer)
}

func solve(day int, part int) int64 {
	switch day {
	case 1:
		switch part {
		case 1:
			answer, err := day1.Problem1()
			if err != nil {
				panic(err)
			}
			return answer
		case 2:
			answer, err := day1.Problem2()
			if err != nil {
				panic(err)
			}
			return answer
		}
	case 2:
		switch part {
		case 1:
			answer, err := day2.Problem1()
			if err != nil {
				panic(err)
			}
			return answer
		case 2:
			answer, err := day2.Problem2()
			if err != nil {
				panic(err)
			}
			return answer
		}
	}
	return 0
}
