package day_01

import (
  "strconv"
)

// CalculateFrequencyDrift for Day 1 Part 1
func CalculateFrequencyDrift(lines []string) int {
  var result = 0
  for i := 0; i < len(lines); i++ {
    d, err := strconv.Atoi(lines[i])
    if err != nil {
      panic(err)
    }
    result += d
  }
  return result
}

// FindFirstFrequencyRepeat for Day 1 Part 2
func FindFirstFrequencyRepeat(lines []string) int {
  var repeats = make(map[int]int)
  var result = 0
  for {
    for i := 0; i < len(lines); i++ {
      d, err := strconv.Atoi(lines[i])
      if err != nil {
        panic(err)
      }
      result += d
      repeats[result] += 1
      if repeats[result] == 2 {
        return result
      }
    }
  }
}
