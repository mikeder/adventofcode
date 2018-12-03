package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

const inputFile = "input.txt"

var Answer = 0
var Repeats = make(map[int]int)

func main() {
	data := readFile(inputFile)
	lines := strings.Split(data, "\n")
	for {
		repeat, stop := processLines(lines)
		fmt.Println(repeat, stop)
		if stop {
			fmt.Println("Repeat found: " + strconv.Itoa(repeat))
			break
		}
	}
}

func processLines(lines []string) (int, bool) {
	for i := 0; i < len(lines); i++ {
		if lines[i] != "" {
			operation, value := parseLine(lines[i])
			doMaths(operation, value)
			Repeats[Answer] += 1
      if Repeats[Answer] == 2 {
        return Answer, true
      }
		}
	}
	return Answer, false
}

func readFile(path string) string {
	dat, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(dat)
}

func doMaths(o string, v int) {
	if o == "+" {
		Answer += v
	} else if o == "-" {
		Answer -= v
	}
}

func parseLine(line string) (string, int) {
	split := strings.Split(line, "")
	operation := split[0]
	return operation, convert(strings.Join(split[1:], ""))
}

func convert(num string) int {
	ret, err := strconv.Atoi(num)
	if err != nil {
		panic(err)
	}
	return ret
}
