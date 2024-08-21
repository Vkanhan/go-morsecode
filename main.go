package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

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
func morseCodeToEnglish(morseCode string) string {
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
func englishToMorse(englishText string) string {
	var morseCode strings.Builder

	// Convert text to uppercase to match the Morse code map
	englishText = strings.ToUpper(englishText)

	for _, char := range englishText {
		switch {
		case char == ' ':
			morseCode.WriteString("   ") // Add three spaces between words
		default:
			if morse, ok := englishToMorseMap[string(char)]; ok {
				morseCode.WriteString(morse + " ") // Add a space after each letter
			} else {
				morseCode.WriteString("?") // If the symbol is not recognized
			}
		}
	}

	return strings.TrimSpace(morseCode.String()) // Remove the trailing space
}

// Function to handle user input and output
func handleUserInput() {
	// Create a new scanner to read user input
	scanner := bufio.NewScanner(os.Stdin)

	// Prompt the user to choose an option
	fmt.Println("Choose an option:")
	fmt.Println("1. Convert Morse code to English")
	fmt.Println("2. Convert English to Morse code")
	fmt.Print("Enter 1 or 2: ")
	scanner.Scan()
	option := scanner.Text()

	switch option {
	case "1":
		// Morse code to English
		fmt.Println("Enter Morse code (use spaces to separate letters and three spaces to separate words):")
		scanner.Scan()
		morseCode := scanner.Text()

		// Convert Morse code to English
		englishText := morseCodeToEnglish(morseCode)

		// Output the result
		fmt.Println("English Text:", englishText)

	case "2":
		// English to Morse code
		fmt.Println("Enter English text:")
		scanner.Scan()
		englishText := scanner.Text()

		// Convert English to Morse code
		morseCode := englishToMorse(englishText)

		// Output the result
		fmt.Println("Morse Code:", morseCode)

	default:
		fmt.Println("Invalid option")
	}
}

func main() {
	handleUserInput()
}
