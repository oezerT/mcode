package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
)

func main() {
	for {
		fmt.Printf("Please enter morsecode/plaintext (or ! to exit): ")
		input, _ := readFromConsole()
		input = strings.ToLower(input)

		if strings.TrimSpace(input) == "!" {
			break
		}

		if isMorseCode(input) {
			fmt.Println(decode(input))
		} else {
			fmt.Println(encode(input))
		}
	}
}

func readFromConsole() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	return reader.ReadString('\n')
}

func isMorseCode(input string) bool {
	firstLetter := strings.TrimSpace(string(input[0]))
	return firstLetter == "." || firstLetter == "-"
}
