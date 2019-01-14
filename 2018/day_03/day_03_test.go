package day_03

import (
  "fmt"
  "testing"
)

func TestCreateFabric(t *testing.T) {
	result := len(CreateFabric())
	if result != 1000 {
		t.Error("Wrong Result")
	}
}

func TestCreateClaim(t *testing.T) {
  input := "#123 @ 3,2: 5x4"
  result := CreateClaim(input)
  fmt.Println(result)
}
