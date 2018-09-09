package strings

import (
	"testing"
)

func TestAllCharsUnique(t *testing.T) {
	tests := []struct {
		in       string
		expected bool
	}{
		{"abcdefg", true},
		{"abcdefgfedcba", false},
	}

	for _, test := range tests {
		actual := allCharsUnique(test.in)
		if actual != test.expected {
			t.Errorf("Expected: %v, got: %v", test.expected, actual)
		}
	}
}

func TestAllCharsUniqueNoStructures(t *testing.T) {
	tests := []struct {
		in       string
		expected bool
	}{
		{"abcdefg", true},
		{"abcdefgfedcba", false},
	}

	for _, test := range tests {
		actual := allCharsUniqueNoStructures(test.in)
		if actual != test.expected {
			t.Errorf("Expected: %v, got: %v", test.expected, actual)
		}
	}
}

func TestReverseCString(t *testing.T) {
	tests := []struct {
		in       string
		expected string
	}{
		{"abcd\000", "dcba\000"},
		{"1234\000", "4321\000"},
		{"∂∑ƒå\000", "åƒ∑∂\000"},
		{"b∂∑ƒå\000", "åƒ∑∂b\000"},
	}

	for _, test := range tests {
		actual := reverseCString(test.in)
		if actual != test.expected {
			t.Errorf("Expected: %v, got: %v", test.expected, actual)
		}
	}
}

func TestRemoveDuplicateChars(t *testing.T) {
	tests := []struct {
		in       string
		expected string
	}{
		{"aabbcc", "abc"},
		{"abccbaab", "abc"},
		{"abccbaaaabddggeba", "abcdge"},
	}

	for _, test := range tests {
		actual := removeDuplicateChars(test.in)
		if actual != test.expected {
			t.Errorf("Expected: %v, got: %v", test.expected, actual)
		}
	}
}

func TestAreAnagrams(t *testing.T) {
	tests := []struct {
		in1      string
		in2      string
		expected bool
	}{
		{"b∂∑ƒå", "åƒ∑∂b", true},
		{"tar", "rat", true},
		{"elbow", "below", true},
		{"state", "taste", true},
		{"rar", "rat", false},
		{"elbow", "bellow", false},
		{"cider", "cries", false},
	}

	for _, test := range tests {
		actual := areAnagrams(test.in1, test.in2)
		actual2 := areAnagramsOptimized(test.in1, test.in2)
		// t.Logf("Str1: %v, Str2: %v", test.in1, test.in2)
		if actual != test.expected != actual2 {
			t.Errorf("Expected: %v, got: %v", test.expected, actual)
			t.Errorf("Expected: %v, got: %v", test.expected, actual2)
		}
	}
}

func BenchmarkAeAnagrams(b *testing.B) {
	str1 := "hydroxydeoxycorticosterones"
	str2 := "hydroxydesoxycorticosterone"
	for n := 0; n < b.N; n++ {
		areAnagrams(str1, str2)
	}
}

func BenchmarkAeAnagramsOptimized(b *testing.B) {
	str1 := "hydroxydeoxycorticosterones"
	str2 := "hydroxydesoxycorticosterone"
	for n := 0; n < b.N; n++ {
		areAnagramsOptimized(str1, str2)
	}
}
