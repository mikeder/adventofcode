package day_02

import (
	"strings"
)

// CalculateInventoryChecksum for Day 2 Part 1, calculate rudimentary
// checksum of boxes found in warehouse.
func CalculateInventoryChecksum(lines []string) int {
	doubleCounts := 0
	tripleCounts := 0
	for i := 0; i < len(lines); i++ {
		chars := strings.Split(lines[i], "")
		found := make(map[string]bool)
		found["double"] = false
		found["triple"] = false
		for ic := 0; ic < len(chars); ic++ {
			if strings.Count(lines[i], chars[ic]) == 2 {
				found["double"] = true
			} else if strings.Count(lines[i], chars[ic]) == 3 {
				found["triple"] = true
			}
		}
		if found["double"] {
			doubleCounts += 1
		}
		if found["triple"] {
			tripleCounts += 1
		}
	}
	checksum := doubleCounts * tripleCounts
	return checksum
}

// FindPrototypeFabricBoxes for Day 2 Part 2, finds two boxes with ID's that
// only differ by a single character and returns the characters that are the
// same between them.
func FindPrototypeFabricBoxes(lines []string) string {
	likeIds := []string{}
	for i := 0; i < len(lines); i++ {
		for ii := 0; ii < len(lines); ii++ {
			if diffBoxIds(lines[i], lines[ii]) == 1 {
				likeIds = append(likeIds, lines[ii])
			}
		}
	}
	if len(likeIds) == 2 {
		return findLikeChars(likeIds)
	} else {
		return ""
	}
}

// findLikeChars compares two box ID's and returns the like characters
func findLikeChars(ids []string) string {
	chars0 := strings.Split(ids[0], "")
	chars1 := strings.Split(ids[1], "")
	like := []string{}
	for i := 0; i < len(chars0); i++ {
		if chars0[i] == chars1[i] {
			like = append(like, chars0[i])
		}
	}
	return strings.Join(like, "")
}

// diffBoxIds compares two box ID's and returns the number of characters that
// differ.
func diffBoxIds(box1 string, box2 string) int {
	chars0 := strings.Split(box1, "")
	chars1 := strings.Split(box2, "")
	diff := 0
	for i := 0; i < len(chars0); i++ {
		if chars0[i] != chars1[i] {
			diff += 1
		}
	}
	return diff
}
