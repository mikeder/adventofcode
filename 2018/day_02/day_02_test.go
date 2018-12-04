package day_02

import (
  "testing"
)


func TestCalculateInventoryChecksum0Double0Triple(t *testing.T){
  input := []string{"abcdef", "zyxwqu"}
  result := CalculateInventoryChecksum(input)
  if result != 0 {
    t.Error("Wrong Result")
  }
}

func TestCalculateInventoryChecksum1Double0Triple(t *testing.T){
  input := []string{"aabcde", "zyxwqu"}
  result := CalculateInventoryChecksum(input)
  if result != 0 {
    t.Error("Wrong Result")
  }
}

func TestCalculateInventoryChecksum2Double0Triple(t *testing.T){
  input := []string{"aabcde", "zzyxwq"}
  result := CalculateInventoryChecksum(input)
  if result != 0 {
    t.Error("Wrong Result")
  }
}

func TestCalculateInventoryChecksum1Double1Triple(t *testing.T){
  input := []string{"aabcde", "zzzxwq"}
  result := CalculateInventoryChecksum(input)
  if result != 1 {
    t.Error("Wrong Result")
  }
}

func TestCalculateInventoryChecksum3Double0Triple(t *testing.T){
  input := []string{"aabbde", "zzyxwq"}
  result := CalculateInventoryChecksum(input)
  if result != 0 {
    t.Error("Wrong Result")
  }
}

func TestCalculateInventoryChecksum2Double2Triple(t *testing.T){
  input := []string{"aabbddd", "zzzxwq"}
  result := CalculateInventoryChecksum(input)
  if result != 4 {
    t.Error("Wrong Result")
  }
}
