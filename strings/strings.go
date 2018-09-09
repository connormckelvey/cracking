package strings

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
