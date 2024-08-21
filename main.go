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

// Function to convert morse code to english
func morseCodeToEnglish(morseCode string) string {
	var decodedMessage string

	//words are seperated by 3 spaces
	words := strings.Split(morseCode, "   ")

	for _, word := range words {
		//letters are seperated by 1 space
		letters := strings.Split(word, " ")
		for _, letter := range letters {
			if translatedLetter, found := morseCodeMap[letter]; found {
				decodedMessage += translatedLetter
			} else {
				// if the decoded message is not recognized
				decodedMessage += "?"
			}
		}
		//Add space after each word
		decodedMessage += " "
	}
	return strings.TrimSpace(decodedMessage)
}


func main() {
	//take the user input
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter the Morse Code: ")

	//scan the input
	scanner.Scan()
	morseCode := scanner.Text()

	englishText := morseCodeToEnglish(morseCode)

	fmt.Println(englishText)
}


