package main

import "strings"

func decode(input string) (output string) {
	words := strings.Split(input, "  ")
	for _, word := range words {
		letters := strings.Split(word, " ")
		for i := range letters {
			output += revertKeyValue(letters[i])
		}
		output += " "
	}
	return output
}

func revertKeyValue(value string) (key string) {
  for k, v := range m {
		if v == strings.TrimSpace(value) {
		  return k
    }
  }
  return
}
