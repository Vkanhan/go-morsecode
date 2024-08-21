# Morse Code Converter

A simple program to convert between Morse code and English text.

## Features

- **Morse to English:** Convert any Morse code sequence into readable English text.
- **English to Morse:** Translate English text into Morse code for communication.

## Installation

Clone the repository and navigate into the project directory:

```bash
git clone https://github.com/Vkanhan/go-morsecode.git
cd go-morsecode
```

## Example

**Convert Morse to English:**

```plaintext
Choose an option:
1. Convert Morse code to English
2. Convert English to Morse code
Enter 1 or 2: 1
Enter Morse code (use spaces to separate letters and three spaces to separate words):
.... . .-.. .-.. ---   .-- --- .-. .-.. -..
English Text: HELLO WORLD
```


**Convert English to Morse:**

```plaintext
Choose an option:
1. Convert Morse code to English
2. Convert English to Morse code
Enter 1 or 2: 2
Enter English text:
HELLO WORLD
Morse Code: .... . .-.. .-.. ---   .-- --- .-. .-.. -..
```

## File Structure

```
go-morsecode/
├── main.go
├── morse/
│ ├── morse.go
│ └── morse_test.go
├── .gitignore
└── go.mod
```

## License

This project is open-source under the MIT License. See the [LICENSE](LICENSE) file for details.


