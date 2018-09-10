package strings

import (
	"fmt"
	"unicode"
)

// 1.1
// Implement an algorithm to determine if a string has all unique characters.
func allCharsUnique(str string) bool {
	seen := make(map[rune]bool)
	for _, c := range str {
		if seen[c] {
			return false
		} else {
			seen[c] = true
		}
	}
	return true
}

// 1.1.b
// What if you can not use additional data structures?
func allCharsUniqueNoStructures(str string) bool {
	for i := 0; i < len(str)-1; i++ {
		for j := (i + 1); j < len(str); j++ {
			if str[i] == str[j] {
				return false
			}
		}
	}
	return true
}

// 1.2
// Write code to reverse a C-Style String. (C-String means that “abcd” is
// represented as five characters, including the null character.)
func reverseCString(str string) string {
	// Strings in Go are immutable, []rune is more similar to C's char[]
	runes := []rune(str)

	// -2 to ignore the null char
	for i, j := 0, len(runes)-2; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// 1.3
// Design an algorithm and write code to remove the duplicate characters in a
// string without using any additional buffer. NOTE: One or two additional
// variables are fine. An extra copy of the array is not.
func removeDuplicateChars(str string) string {
	// Strings in Go are immutable, []rune is more similar to C's char[]
	runes := []rune(str)

	for i := 0; i < len(runes); i++ {
		for j := i + 1; j < len(runes); {
			if runes[i] == runes[j] {
				runes = runes[:j+copy(runes[j:], runes[j+1:])]
			} else {
				j++
			}
		}
	}
	return string(runes)
}

// 1.4
// Write a method to decide if two strings are anagrams or not.
func areAnagrams(str1 string, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}

	runeCount := make(map[rune]int)
	for _, char := range str1 {
		runeCount[char]++
	}
	for _, char := range str2 {
		runeCount[char]--
	}
	for _, count := range runeCount {
		if count != 0 {
			return false
		}
	}
	return true
}

// 1.4.b
// After running benchmarks, doesn't seem like this solution is the most optimal // overall, it requires less memory, but takes longer, without a doubt, I would
// choose the solution above as it is easier to read, and faster
func areAnagramsOptimized(str1 string, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}

	runeCount := make(map[rune]int)
	var numUniqueChars int
	var numCompleteChecks int

	for _, char := range str1 {
		runeCount[char]++
		if runeCount[char] == 0 {
			numUniqueChars++
		}
	}

	for i, char := range str2 {
		if runeCount[char] == 0 {
			return false
		}
		runeCount[char]--

		if runeCount[char] == 0 {
			numCompleteChecks++
			if numCompleteChecks == numUniqueChars {
				return i == len(str2)-1
			}
		}
	}

	return false
}

// 1.5
// Write a method to replace all spaces in a string with ‘%20’.
func encodeSpaces(str string) string {
	var runes []rune
	for _, char := range str {
		if unicode.IsSpace(char) {
			runes = append(runes, '%', '2', '0')
		} else {
			runes = append(runes, char)
		}
	}
	return string(runes)
}

// 1.5.b encode spaces in place
func encodeSpacesInPlace(str string) string {
	// Golang strings are immutable, modify the []rune in place
	runes := []rune(str)

	for i := 0; i < len(runes); i++ {
		if unicode.IsSpace(runes[i]) {
			runes = append(runes[:i], append([]rune("%20"), runes[i+1:]...)...)
		}
	}
	return string(runes)
}

// Given an image represented by an NxN matrix, where each pixel in the image is
// 4 bytes, write a method to rotate the image by 90 degrees. Can you do this
// in place?
// Example (just ints, its easier):
// 1 2 3     7 4 1
// 4 5 6 ->  8 5 2
// 7 8 9     9 6 3
//
//     i, j  ->  j, n - i
// 1: (0, 0) -> (0, 2)
// 2: (0, 1) -> (1, 2)
// 3: (0, 2) -> (2, 2)

// 4: (1, 0) -> (0, 1)
// 5: (1, 1) -> (1, 1)
// 6: (1, 2) -> (2, 1)

// 7: (2, 0) -> (0, 0)
// 8: (2, 1) -> (1, 0)
// 9: (2, 2) -> (2, 0)

// 1 2 ->  3 1  got 2 4
// 3 4     4 2      1 3

func rotateSquareMatrix(matrix [][]int) [][]int {
	fmt.Println(matrix)
	n := len(matrix)
	blank := make([][]int, n)
	for i := 0; i < n; i++ {
		blank[i] = make([]int, n)
	}

	fmt.Println(blank)
	for row := range matrix {
		for col := range matrix[row] {
			fmt.Printf("row: %v, col: %v \n", row, col)
			blank[row][col] = matrix[n-1-col][row]
		}
	}
	return blank
}
