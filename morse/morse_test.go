package morse 


import(
	"testing"
)

//Test the MorseCodeToEnglish function
func TestMorseCodeToEnglish(t *testing.T) {

	tests := []struct{
		input string 
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

//Test the EnglishToMorse function
func TestEnglishToMorse(t *testing.T) {
	tests := []struct{
		input string 
		expected string 
	}{
		{"A B C", ".-   -...   -.-."},
		{"GOLANG", "--. --- .-.. .- -. --.",},
		{"MORSE CODE", "-- --- .-. ... .   -.-. --- -.. ."},

	}

	for _, test := range tests {
		result := EnglishToMorse(test.input)
		if result != test.expected {
			t.Errorf("englishToMorse(%q): %q; want: %q", test.input, result, test.expected)
		}
	}
}

