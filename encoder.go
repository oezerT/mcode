package main

import "strings"

func encode(input string) (output string) {
	result := strings.SplitAfter(input, "")

	for i := range result {
		output += m[result[i]] + " "
	}

	return output
}
