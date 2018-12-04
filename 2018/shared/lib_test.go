package shared

import (
	"adventofcode/2018/shared"
	"testing"
)

func TestGetInputPathSingleDigit(t *testing.T) {
	input := 1
	result := shared.GetInputPath(input)
	if result != "./day_01/input.txt" {
		t.Error("Wrong result")
	}
}

func TestGetInputPathDoubleDigit(t *testing.T) {
	input := 10
	result := shared.GetInputPath(input)
	if result != "./day_10/input.txt" {
		t.Error("Wrong result")
	}
}

func TestGetInputPathDoubleDigit2(t *testing.T) {
	input := 20
	result := shared.GetInputPath(input)
	if result != "./day_20/input.txt" {
		t.Error("Wrong result")
	}
}
