package main

import (
	"fmt"
	"sort"
)

func main() {
	letterCombinations("23")
}

func letterCombinations(digits string) []string {
	keyboardMapping := map[string]string{
		"2": "abc",
		"3": "def",
		"4": "ghi",
		"5": "jkl",
		"6": "mno",
		"7": "pqrs",
		"8": "tuv",
		"9": "wxyz",
	}

	selectedLetters := []string{}
	for i := range digits {
		selectedLetters = append(selectedLetters, keyboardMapping[string(digits[i])])
	}
	fmt.Printf("selectedLetters = %v\n", selectedLetters)

	letterPointers := []int{}
	// Initialise pointers at sizes of strings i.e. starting right to left
	for i := range selectedLetters {
		letterPointers = append(letterPointers, len(selectedLetters[i])-1)
	}
	originalPointers := make([]int, len(letterPointers))
	_ = copy(originalPointers, letterPointers)
	fmt.Printf("letterPointers = %v\n", letterPointers)

	letterCombos := []string{}
	// Calculate current set of letters based on pointers
	letterCombos = findCombos(letterPointers, selectedLetters, letterCombos, originalPointers)
	fmt.Println("final combos:")
	sort.Strings(letterCombos)
	fmt.Println(letterCombos)

	return letterCombos
}

func findCombos(pointers []int, letters, combos []string, originalPointers []int) []string {
	// If all pointer values are at 0, this is the final iteration
	lastPointer := false
	if pointersAreDone(pointers) {
		lastPointer = true
	}

	combo := ""
	for i := range pointers {
		combo += string(letters[i][pointers[i]])
	}
	fmt.Printf("combo: %v\n", combo)

	combos = append(combos, combo)

	// iterate pointers
	if lastPointer {
		return combos
	}
	pointers = countDownPointers(pointers, originalPointers)
	return findCombos(pointers, letters, combos, originalPointers)
}

func pointersAreDone(pointers []int) bool {
	for _, v := range pointers {
		if v != 0 {
			return false
		}
	}
	return true
}

// counts pointers down by 1 e.g. [2,3,4] -> [2,3,3]
// Requres original values to know what to reset a counter to
// e.g. if originals are [3,3,3], then counter [2,3,0,] -> [2,2,3]
func countDownPointers(pointers []int, originals []int) []int {
	for i := len(pointers) - 1; i >= 0; i-- {
		if pointers[i] != 0 {
			pointers[i] = pointers[i] - 1
			// Reset previous pointer values
			for j := i + 1; j <= len(pointers)-1; j++ {

				pointers[j] = originals[j]
			}
			break
		}
	}
	return pointers
}
