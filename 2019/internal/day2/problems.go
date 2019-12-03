package day2

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

type opcode int8

const (
	opcodeAdd      = opcode(1)
	opcodeMultiply = opcode(2)
	opcodeHalt     = opcode(99)
)

type intcode struct {
	oc  opcode
	n1  int64
	n2  int64
	out int64
}

var input []int64

// Problem1 solves the Intcode computer bug
func Problem1() (int64, error) {
	input = parseInput()
	pos := 0
	input[1] = 12
	input[2] = 2
	for {
		instruction := intcode{
			oc:  opcode(input[pos]),
			n1:  input[pos+1],
			n2:  input[pos+2],
			out: input[pos+3],
		}
		if instruction.oc == opcodeHalt {
			break
		}
		doLogic(instruction)
		if input[0] == 19690720 {
			fmt.Println(instruction)
		}
		pos += 4
	}

	return input[0], nil
}

// Problem2 solves the Intcode computer bug
func Problem2() (int64, error) {
	min := 0
	max := 99
	want := int64(19690720)
	pos := 0
	input = parseInput()
	for {
		instruction := intcode{
			oc:  opcode(input[pos]),
			n1:  input[pos+1],
			n2:  input[pos+2],
			out: input[pos+3],
		}
		if instruction.oc == opcodeHalt {
			input = parseInput()
			input[1] = int64(rand.Intn(max-min) + min)
			input[2] = int64(rand.Intn(max-min) + min)
			// fmt.Println(input[1], input[2])
			pos = 0
			continue
		}
		doLogic(instruction)
		if input[0] == want {
			vals := fmt.Sprintf("%02d%02d", input[1], input[2])
			want, err := strconv.ParseInt(vals, 10, 64)
			return want, err
		}
		pos += 4
	}

	return want, nil
}

func parseInput() []int64 {
	var vals []int64
	for _, i := range strings.Split(rawInput, ",") {
		num, err := strconv.ParseInt(i, 10, 64)
		if err != nil {
			return nil
		}
		vals = append(vals, num)
	}
	return vals
}

func doLogic(c intcode) {
	switch c.oc {
	case opcodeAdd:
		input[c.out] = input[c.n1] + input[c.n2]
	case opcodeMultiply:
		input[c.out] = input[c.n1] * input[c.n2]
	case opcodeHalt:
		return
	}
}
