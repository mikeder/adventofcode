package day_01

import (
  "testing"
)

func TestCalculateFrequencyDriftResult1(t *testing.T) {
  input := []string{"0", "+1"}
  result := CalculateFrequencyDrift(input)
  if result != 1 {
    t.Error("Wrong result")
  }
}

func TestCalculateFrequencyDriftResultNegative(t *testing.T) {
  input := []string{"1", "-2"}
  result := CalculateFrequencyDrift(input)
  if result != -1 {
    t.Error("Wrong result")
  }
}

func TestCalculateFrequencyDriftResultPositive(t *testing.T) {
  input := []string{"-1", "+2", "+3"}
  result := CalculateFrequencyDrift(input)
  if result != 4 {
    t.Error("Wrong result")
  }
}
