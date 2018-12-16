package main

import (
	"fmt"
	"strings"
)

func main() {
	arrays := []*string{}
	str := "This is a sample text but a complicated problem to be solved, " +
		"so we are adding more text to see that it actually works thoughhhh."
	wordformat(str, arrays)
	//printArrays(arrays)
}

func wordformat(s string, arrays []*string) {
	const space string = " "
	start := 0
	end := 0
	var subs string
	length, newSpaces, currentSpaces := 0, 0, 0
	tokens := []string{}
	const pageWidth int = 20

	getSpace := map[int]string{0: "", 1: " ", 2: "  ", 3: "   ", 4: "    ", 5: "     ", 6: "      "}

	// Loops through the input paragraph and break it down into a number of substrings
	// that are each less than or equal to 20 in length without breaking any words
	for len(s) > 0 {
		if (start + 21) > len(s) { // Handles the last substring of paragraph
			end = len(s)
			subs = s[start:end]
			arrays = append(arrays, &subs)
			break

		} else { // Grab the first chunk and the next through the loop
			end = start + 21
		}
		subs = s[start:end]
		si := strings.LastIndex(subs, space)
		st := subs[:si]
		arrays = append(arrays, &st) // Add each substring to the array
		start += si + 1
	}
	// Loops through each substring and get each one justified.
	// Each substring must begin with a word and end with a word.
	for a, line := range arrays {
		length = len(*line)
		newSpaces = pageWidth - length
		tokens = strings.Fields(*line)
		if len(tokens) == 1 {
			arrays[a] = &tokens[0]
			fmt.Printf("Array [%v] = \"%v\"\n", a+1, *arrays[a])
			break
		}
		currentSpaces = len(tokens) - 1

		if newSpaces > currentSpaces {
			quotient := newSpaces / currentSpaces
			var finalStr string
			for i := 0; i < len(tokens); i++ {
				if i == len(tokens)-1 {
					// The last word can only have space added to the left, if any.
					// Makes sure the final string is exactly 20 in length.
					moreSpace := pageWidth - len(*arrays[a]) - len(tokens[i])
					finalStr += getSpace[moreSpace] + tokens[i]
					arrays[a] = &finalStr
					fmt.Printf("Array [%v] = \"%v\"\n", a+1, *arrays[a])
					break
				}
				finalStr += tokens[i] + getSpace[quotient]
				arrays[a] = &finalStr
			}

		} else {
			var finalStr string
			// this covers newSpaces < currentSpaces and
			// newSpaces == currentSpaces and possibly anything else.
			for i := 0; i < len(tokens); i++ {
				if i == len(tokens)-1 {
					// The last word can only have space added to the left, if any.
					// Makes sure the final string is exactly 20 in length.
					moreSpace := pageWidth - len(*arrays[a]) - len(tokens[i])
					finalStr += getSpace[moreSpace] + tokens[i]
					arrays[a] = &finalStr
					fmt.Printf("Array [%v] = \"%v\"\n", a+1, *arrays[a])
					break
				}
				finalStr += tokens[i] + getSpace[1]
				arrays[a] = &finalStr
			}
		}
	}

}
