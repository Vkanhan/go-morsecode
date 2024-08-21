package morse

import "strings"

// Morse code dictionary
var morseCodeMap = map[string]string{
	".-":    "A",
	"-...":  "B",
	"-.-.":  "C",
	"-..":   "D",
	".":     "E",
	"..-.":  "F",
	"--.":   "G",
	"....":  "H",
	"..":    "I",
	".---":  "J",
	"-.-":   "K",
	".-..":  "L",
	"--":    "M",
	"-.":    "N",
	"---":   "O",
	".--.":  "P",
	"--.-":  "Q",
	".-.":   "R",
	"...":   "S",
	"-":     "T",
	"..-":   "U",
	"...-":  "V",
	".--":   "W",
	"-..-":  "X",
	"-.--":  "Y",
	"--..":  "Z",
	"-----": "0",
	".----": "1",
	"..---": "2",
	"...--": "3",
	"....-": "4",
	".....": "5",
	"-....": "6",
	"--...": "7",
	"---..": "8",
	"----.": "9",
}

// Reverse the dictionary for converting English to Morse code
var englishToMorseMap = reverseMap(morseCodeMap)

// reverseMap function creates a reversed map for english to morse code
func reverseMap(originalMap map[string]string) map[string]string {
	reversedMap := make(map[string]string)

	// Swap key-value pairs from the original map to create the reversed map.
	for key, value := range originalMap {
		reversedMap[value] = key
	}
	return reversedMap
}

// Function to convert morse code to english
func MorseCodeToEnglish(morseCode string) string {
	var decodedMessage strings.Builder

	//words are seperated by 3 spaces
	words := strings.Split(morseCode, "   ")

	for _, word := range words {
		//letters are seperated by 1 space
		letters := strings.Split(word, " ")

		//Look up each letter in the morseCodeMap and append the corresponding
		// translated letter to the decodedMessage if the letter is found in the map.
		for _, letter := range letters {
			if translatedLetter, found := morseCodeMap[letter]; found {
				decodedMessage.WriteString(translatedLetter)
			} else {
				// if the decoded message is not recognized
				decodedMessage.WriteString("?")
			}
		}
		//Add space after each word
		decodedMessage.WriteString(" ")
	}
	return strings.TrimSpace(decodedMessage.String())
}

// Function to convert English text to Morse code
func EnglishToMorse(englishText string) string {
	var morseCode strings.Builder

	// Convert text to uppercase to match the Morse code map
	englishText = strings.ToUpper(englishText)

	for _, char := range englishText {
		switch {
		case char == ' ':
			morseCode.WriteString("  ") // Add two spaces between words as there is already one space don't add three spaces.
		default:
			if morse, ok := englishToMorseMap[string(char)]; ok {
				morseCode.WriteString(morse + " ") // Add a space after each letter
			} else {
				morseCode.WriteString("?") // If the symbol is not recognized
			}
		}
	}
	return strings.TrimSpace(morseCode.String()) 
}
