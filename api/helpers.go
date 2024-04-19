package api

import (
	"fmt"
	"github.com/gary-norman/ascii-art/pkg"
	"strings"
)

// ArtToSingleLine Places each line of characters from FileToVariable on a single line, delineated by "** "
func ArtToSingleLine(source []string) []string {
	var output []string
	if len(source) == 8 {
		return source
	}
	for i := 0; i < 8; i++ {
		output = append(output, source[i])
	}
	if len(source) > 8 {
		source = source[8:]
	}
	x := 0
	for len(source) > 0 {
		for i := 0; i < 8; i++ {
			output = append(output, output[i+x]+"* "+"# "+source[i])
		}
		x += 8
		if len(source) > 8 {
			source = source[8:]
		} else {
			source = nil
		}
	}
	for len(output) > 8 {
		output = output[8:]
	}
	return output
}

// GetEmptyCols Get the index of the final space of each character in the reverse flag
func GetEmptyCols(source []string) []int {
	source = ArtToSingleLine(source)
	var emptyCols []int
	for i := 0; i < len(source[0]); i++ {
		empty := true
		for j := 0; j < len(source); j++ {
			if source[j][i] != 32 {
				empty = false
			}
		}
		if empty == true {
			emptyCols = append(emptyCols, i)
		}
	}
	return emptyCols
}

// RemoveValidSPaceIndex Remove indices for valid spaces, before the end space
func RemoveValidSPaceIndex(indices []int) []int {
	for i := 0; i < len(indices)-1; i++ {
		if len(indices)-i > 6 {
			if indices[i] == (indices[i+6])-6 {
				indices = append(indices[:i+1], indices[i+6:]...)
			}
		}
	}
	return indices
}

// GetChars Map the ascii characters provided in the style.txt file, indexed by ascii code
func GetChars(source []string) map[int][]string {
	charMap := make(map[int][]string)
	id := 31
	for _, line := range source {
		if string(line) == "" {
			id++
		} else {
			charMap[id] = append(charMap[id], line)
		}
	}
	return charMap
}

// GetInputChars Map the ascii characters provided in the reverse flag, zero indexed
func GetInputChars(source []string, indices []int) map[int][]string {
	charMap := make(map[int][]string)
	startIndex := 0
	for id := range indices {
		for _, line := range source {
			charMap[id] = append(charMap[id], line[startIndex:indices[id]]+" ")
		}
		startIndex = indices[id] + 1
	}
	return charMap
}

// AsciiToChars Compares getChar and getInputChar and prints the string to the terminal
func AsciiToChars(input, standard, shadow, thinkertoy map[int][]string) {
	output := make(map[int][]int)
	var newLine1 []string
	for i := 0; i < 8; i++ {
		newLine1 = append(newLine1, "* ")
	}
	var newLine2 []string
	for i := 0; i < 8; i++ {
		newLine2 = append(newLine2, "# ")
	}
	slash := 92
	n := 110
	styles := []map[int][]string{standard, shadow, thinkertoy}
	for _, style := range styles {
		for key1, slice1 := range input {
			for key2, slice2 := range style {
				if pkg.CompareSlices(slice1, newLine1) {
					output[key1] = append(output[key1], slash)
				}
				if pkg.CompareSlices(slice1, newLine2) {
					output[key1] = append(output[key1], n)
				}
				if pkg.CompareSlices(slice1, slice2) {
					output[key1] = append(output[key1], key2)
				}
			}
		}
	}
	for i := 0; i < len(output); i++ {
		fmt.Printf("%c", output[i][0])
	}
}

// GetCharsWidth Determine the width of each individual ascii art character
func GetCharsWidth(source []string) map[int]int {
	charWidthMap := make(map[int]int)
	id := 31
	for _, line := range source {
		if string(line) == "" {
			id++
		} else {
			charWidthMap[id] = len(line)
		}
	}
	return charWidthMap
}

// GetArtWidth Determine the width of each line that gets printed to the terminal (without EOL)
func GetArtWidth(origString string, y map[int]int) []int {
	var width []int
	replaceNewline := strings.ReplaceAll(origString, "\r\n", "\\n") // correct newline formatting
	wordSlice := strings.Split(replaceNewline, "\\n")
	for i := 0; i < len(wordSlice); i++ {
		sum := 0
		for _, char := range wordSlice[i] {
			sum += y[int(char)]
			//for _, num := range y[int(char)] {
			//	sum += num
			//}
		}
		width = append(width, sum)
	}
	return width
}
