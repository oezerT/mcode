package main

import "strings"

func decode(input string) (output string) {
	words := strings.Split(input, "  ")
	for _, word := range words {
		letters := strings.Split(word, " ")
		for i := range letters {
			output += mapkey(letters[i])
		}
		output += " "
	}
	return output
}

func mapkey(value string) (key string) {
  for k, v := range m {
		if v == strings.TrimSpace(value) {
		  return k
    }
  }
  return
}
