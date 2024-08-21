package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Vkanhan/go-morsecode/morse"
)

// handleConversion prompts the user for input, performs the conversion, and displays the result
func handleConversion(prompt string, converter func(string)string) {
	// Prompt the user for input
	fmt.Println(prompt)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	// Perform the conversion using the provided converter function 
	result := converter(input)

	// Display the result
	fmt.Println("Result: ", result)
}

func main() {
	//Display the menu options to the user
	fmt.Println("Choose an option:\n1. Convert Morse code to English\n2. Convert English to Morse code\nEnter 1 or 2: ")

	// Create a scanner to read user input
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	option := scanner.Text()

	// Handle the user's choice
	switch option {
	case "1":
		handleConversion("Enter Morse code (use spaces to separate letters and three spaces to separate words):", morse.MorseCodeToEnglish)
	case "2":
		handleConversion("Enter English text:", morse.EnglishToMorse)
	default:
		fmt.Println("Invalid option")
	}
}


