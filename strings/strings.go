package strings

import (
	"strings"
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
func rotateSquareMatrix(matrix [][]int) [][]int {
	n := len(matrix)
	blank := make([][]int, n)
	for i := 0; i < n; i++ {
		blank[i] = make([]int, n)
	}
	for row := range matrix {
		for col := range matrix[row] {
			blank[row][col] = matrix[n-1-col][row]
		}
	}
	return blank
}

func rotateSquareMatrixInPlace(matrix [][]int) [][]int {
	n := len(matrix)
	for layer := 0; layer < n; layer++ {
		first := layer
		last := n - 1 - layer
		for i := first; i < last; i++ {
			offset := i - first
			top := matrix[first][i]                                // save top
			matrix[first][i] = matrix[last-offset][first]          // left -> top
			matrix[last-offset][first] = matrix[last][last-offset] // bottom -> left
			matrix[last][last-offset] = matrix[i][last]            // right -> bottom
			matrix[i][last] = top                                  // top -> right
		}
	}
	return matrix
}

// Assume you have a method isSubstring which checks if one word is a substring
// of another. Given two strings, s1 and s2, write code to check if s2 is a
// rotation of s1 using only one call to isSubstring (i.e., “waterbottle” is a
// rotation of “erbottlewat”).
func isRotatation(str1 string, str2 string) bool {
	if len(str1) == len(str2) {
		return strings.Contains(str1+str1, str2)
	}
	return false
}
