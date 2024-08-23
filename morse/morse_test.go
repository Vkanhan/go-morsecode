package morse

import (
	"strings"
	"testing"
)

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

func BenchmarkMorseCodeToEnglish(b *testing.B) {
	longInput := strings.Repeat(". ", 10000)
	for i := 0; i < b.N; i++ {
		MorseCodeToEnglish(longInput)
	}
}

func FuzzMorseToEnglish(f *testing.F) {
	f.Add(".... . .-.. .-.. ---   .-- --- .-. .-.. -..")
	f.Add("... --- ...")
	f.Add("-- --- .-. ... .   -.-. --- -.. .")
	f.Add("----.----..----.----.-----")
	f.Add("   ")
	f.Add("..--..--..")
	f.Add("")

	f.Fuzz(func(t *testing.T, morseCode string) {
		_ = MorseCodeToEnglish(morseCode)
	})
}

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

func BenchMarkEnglishToMorse(b *testing.B) {
	englishText := strings.Repeat("E", 10000)
	for i := 0; i < b.N; i++ {
		EnglishToMorse(englishText)
	}
}

func FuzzEnglishToMorse(f *testing.F) {
	f.Add("HELLO WORLD")
	f.Add("SOS")
	f.Add("MORSE CODE")
	f.Add("12345")
	f.Add(" ")
	f.Add("!@#$%")
	f.Add("")

	f.Fuzz(func(t *testing.T, englishText string) {
		_ = EnglishToMorse(englishText)
	})
}
