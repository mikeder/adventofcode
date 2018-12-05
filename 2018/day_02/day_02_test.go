package day_02

import (
	"testing"
)

func TestCalculateInventoryChecksum0Double0Triple(t *testing.T) {
	input := []string{"abcdef", "zyxwqu"}
	result := CalculateInventoryChecksum(input)
	if result != 0 {
		t.Error("Wrong Result")
	}
}

func TestCalculateInventoryChecksum1Double0Triple(t *testing.T) {
	input := []string{"aabcde", "zyxwqu"}
	result := CalculateInventoryChecksum(input)
	if result != 0 {
		t.Error("Wrong Result")
	}
}

func TestCalculateInventoryChecksum2Double0Triple(t *testing.T) {
	input := []string{"aabcde", "zzyxwq"}
	result := CalculateInventoryChecksum(input)
	if result != 0 {
		t.Error("Wrong Result")
	}
}

func TestCalculateInventoryChecksum1Double1Triple(t *testing.T) {
	input := []string{"aabcde", "zzzxwq"}
	result := CalculateInventoryChecksum(input)
	if result != 1 {
		t.Error("Wrong Result")
	}
}

func TestCalculateInventoryChecksum3Double0Triple(t *testing.T) {
	input := []string{"aabbde", "zzyxwq"}
	result := CalculateInventoryChecksum(input)
	if result != 0 {
		t.Error("Wrong Result")
	}
}

func TestCalculateInventoryChecksum2Double2Triple(t *testing.T) {
	input := []string{"aabbddd", "zzzxwq"}
	result := CalculateInventoryChecksum(input)
	if result != 2 {
		t.Error("Wrong Result")
	}
}

func TestFindPrototypeFabricBoxes1Diff(t *testing.T) {
	input := []string{"abcde", "abgde"}
	result := FindPrototypeFabricBoxes(input)
	if result != "abde" {
		t.Error("Wrong Result")
	}
}

func TestFindPrototypeFabricBoxes2Diff(t *testing.T) {
	input := []string{"abcde", "axcye"}
	result := FindPrototypeFabricBoxes(input)
	if result != "" {
		t.Error("Wrong Result")
	}
}
