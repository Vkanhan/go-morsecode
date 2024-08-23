package morse

import (
	"strings"
	"testing"
)

// Test the MorseCodeToEnglish function
func TestMorseCodeToEnglish(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{".-   -...   -.-.", "A B C"},
		{"--. --- .-.. .- -. --.", "GOLANG"},
		{"-- --- .-. ... .   -.-. --- -.. .", "MORSE CODE"},
	}

	for _, test := range tests {
		result := MorseCodeToEnglish(test.input)
		if result != test.expected {
			t.Errorf("morseToEnglish(%q): %q; want: %q", test.input, result, test.expected)
		}
	}
}

// BenchMark test for MorseCodeToEnglish function
func BenchmarkMorseCodeToEnglish(b *testing.B) {
	longInput := strings.Repeat(". ", 10000)
	for i := 0; i < b.N; i++ {
		MorseCodeToEnglish(longInput)
	}
}

// Test the EnglishToMorse function
func TestEnglishToMorse(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"A B C", ".-   -...   -.-."},
		{"GOLANG", "--. --- .-.. .- -. --."},
		{"MORSE CODE", "-- --- .-. ... .   -.-. --- -.. ."},
	}

	for _, test := range tests {
		result := EnglishToMorse(test.input)
		if result != test.expected {
			t.Errorf("englishToMorse(%q): %q; want: %q", test.input, result, test.expected)
		}
	}
}

// BenchMark test for EnglishToMorse function
func BenchMarkEnglishToMorse(b *testing.B) {
	englishText := strings.Repeat("E", 10000)
	for i := 0; i < b.N; i++ {
		EnglishToMorse(englishText)
	}
}
