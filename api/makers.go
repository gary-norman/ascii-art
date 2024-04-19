package api

import (
	"fmt"
	"github.com/gary-norman/ascii-art/pkg"
	"strings"
)

// MakeArtAligned Transform the input text origString to the output art, line by line, with left, right, or center aligned content
func MakeArtAligned(origString string, y map[int][]string, ds []int, ws pkg.Winsize, divider int) string {
	var art string
	replaceNewline := strings.ReplaceAll(origString, "\r\n", "\\n") // correct newline formatting
	wordSlice := strings.Split(replaceNewline, "\\n")
	for i := 0; i < len(wordSlice); i++ {
		for j := 0; j < len(y[32]); j++ {
			var line string
			art += strings.Repeat(" ", (int(ws.Col)-ds[i])/divider)
			for _, letter := range wordSlice[i] {
				line = line + y[int(letter)][j]
			}
			art += line + "\n"
			line = ""
		}
	}
	art = strings.TrimRight(art, "\n")
	return art
}

// MakeArtJustified Transform the input text origString to the output art, line by line, with justified content
func MakeArtJustified(origString string, y map[int][]string, ds []int, ws pkg.Winsize) string {
	var art string
	replaceNewline := strings.ReplaceAll(origString, "\r\n", "\\n") // correct newline formatting
	wordSlice := strings.Split(replaceNewline, "\\n")
	for i := 0; i < len(wordSlice); i++ {
		for j := 0; j < len(y[32]); j++ {
			var line string
			spaces := 0
			for _, letter := range wordSlice[i] {
				if letter == 32 {
					spaces += 1
				}
			}
			for _, letter := range wordSlice[i] {
				if spaces > 0 {
					line = line + y[int(letter)][j]
					if letter == 32 {
						line = line + strings.Repeat(" ", (int(ws.Col)-ds[i])/spaces)
					}
				} else {
					line = line + y[int(letter)][j]
					sliceLen := len(wordSlice[i])
					if sliceLen >= 2 {
						line = line + strings.Repeat(" ", (int(ws.Col)-ds[i])/(sliceLen-1))
					}
				}
			}
			line = strings.TrimRight(line, " ")
			art += line + "\n"
			line = ""
		}
	}
	art = strings.TrimRight(art, "\n")
	return art
}

// MakeArtColorized Transform the input text origString to the output art, line by line, colorizing specified text
func MakeArtColorized(origString string, y map[int][]string, letters []rune, color string, colorAll bool) string {
	var specifiedColor string
	reset := "\033[0m"
	switch color {
	case "red":
		specifiedColor = "\033[31m"
	case "green":
		specifiedColor = "\033[32m"
	case "yellow":
		specifiedColor = "\033[33m"
	case "blue":
		specifiedColor = "\033[34m"
	case "orange":
		specifiedColor = "\033[38;5;208m"
	default:
		fmt.Print("\nAvailable colors are " + "\033[31m" + "red" + reset + ", " +
			"\033[32m" + "green" + reset + "," + "\033[33m" + "yellow" + reset + ", " +
			"\033[38;5;208m" + "orange" + reset + ", and " + "\033[34m" + "blue" + reset + ".\n")
	}
	var art string
	replaceNewline := strings.ReplaceAll(origString, "\r\n", "\\n") // correct newline formatting
	wordSlice := strings.Split(replaceNewline, "\\n")
	for _, word := range wordSlice {
		for j := 0; j < len(y[32]); j++ {
			var line string
			if colorAll {
				for _, letter := range word {
					line = specifiedColor + line + y[int(letter)][j]
				}
			} else {
				for _, letter := range word {
					if pkg.SliceContainsItem(letters, letter) {
						line = line + specifiedColor + y[int(letter)][j] + reset
					} else {
						line = line + y[int(letter)][j]
					}
				}
			}
			art += line + "\n" + reset
			line = ""
		}
	}
	art = strings.TrimRight(art, "\n")
	return art
}

// MakeArt Transform the input text origString to the output art, line by line
func MakeArt(origString string, y map[int][]string) string {
	var art string
	replaceNewline := strings.ReplaceAll(origString, "\r\n", "\\n") // correct newline formatting
	wordSlice := strings.Split(replaceNewline, "\\n")               // split the input into slices
	for _, word := range wordSlice {                                // loop over the word to get the characters
		for j := 0; j < len(y[32]); j++ { // loop over each vertical line of the word
			var line string
			for _, letter := range word { // loop over each character
				line = line + y[int(letter)][j] // add each line of the character to the line string
			}
			art += line + "\n" // add each line string (followed by a line break) to the final output
			line = ""
		}
	}
	art = strings.TrimRight(art, "\n") // remove the final line break
	return art
}
