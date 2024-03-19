package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"

	// "regexp"
	"strings"
)

/**
* TODO reverse

* ^^ fs
	* ^^ SOLVED functionality works with shadow and standard but not with thinkertoy

* TODO color

* ^^ output

* TODO align

*/

// import standard.txt as the default ascii style, with ability to change it
// using 2nd argument *
// ^^ thinkertoy does not work
func PrepareBan(bannerStyle string) []string {
	if bannerStyle == "" {
		bannerStyle = "standard"
	}
	style := bannerStyle
	file, err := os.Open("ascii_styles/" + style + ".txt")
	if err != nil {
		log.Fatal(err)
	}
	scanned := bufio.NewScanner(file)
	scanned.Split(bufio.ScanLines)
	var source []string
	for scanned.Scan() {
		source = append(source, scanned.Text())
	}
	file.Close()
	return source
}

func getChars(source []string) map[int][]string {
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

// transform the input text origString to the output art, line by line
func makeArt(origString string, y map[int][]string) string {
	var art string
	replaceNewline := strings.ReplaceAll(origString, "\r\n", "\\n") // correct newline formatting
	wordSlice := strings.Split(replaceNewline, "\\n")
	for _, word := range wordSlice {
		for j := 0; j < len(y[32]); j++ {
			var line string
			for _, letter := range word {
				line = line + string((y[int(letter)][j]))
			}
			art += line + "\n"
			line = ""
		}
	}
	return art
}

// Contains checks if a slice contains a specific element.
func Contains(slice []rune, item rune) bool {
	for _, v := range slice {
		if v == item {
			return true // Found the item
		}
	}
	return false // Item not found
}

// transform the input text origString to the output art, line by line, colorizing specified text
func makeArtColorized(origString string, y map[int][]string, letters []rune, color string, colorAll bool) string {
	var specifiedColor string
	switch color {
	case "red":
		specifiedColor = "\033[31m"
	case "green":
		specifiedColor = "\033[32m"
	case "yellow":
		specifiedColor = "\033[33m"
	case "blue":
		specifiedColor = "\033[34m"
	}
	reset := "\033[0m"
	var art string
	replaceNewline := strings.ReplaceAll(origString, "\r\n", "\\n") // correct newline formatting
	wordSlice := strings.Split(replaceNewline, "\\n")
	for _, word := range wordSlice {
		for j := 0; j < len(y[32]); j++ {
			var line string
			if colorAll {
				for _, letter := range word {
					line = specifiedColor + line + string((y[int(letter)][j]))
				}
			} else {
				for _, letter := range word {
					if Contains(letters, letter) {
						line = line + specifiedColor + string((y[int(letter)][j])) + reset
					} else {
						line = line + string((y[int(letter)][j]))
					}
				}
			}
			art += line + "\n" + reset
			line = ""
		}
	}
	return art
}

func main() {

	// ? flag definitions
	reverse := flag.String("reverse", "default", "Tell the program to run the reverse function")
	color := flag.String("color", "default", "Tell the program to run the color function")
	output := flag.String("output", "default", "Save the output to the specified filename")
	align := flag.String("align", "default", "Tell the program to run the align function")
	flag.Parse()
	additionalArgs := flag.Args()
	input := additionalArgs[0]
	var bannerStyle string
	if len(additionalArgs) == 2 {
		bannerStyle = additionalArgs[1]
	}
	if *reverse != "default" {
		fmt.Printf("Reverse flag is set to  %v\n", *reverse)
		return
	}
	if *color != "default" {
		var colored string
		var colorAll bool
		var colSLice []rune
		colorAll = true
		if len(additionalArgs) == 2 {
			colorAll = false
			colored = additionalArgs[0]
			colSLice = []rune(colored)
			input = additionalArgs[1]
		}
		fmt.Println(makeArtColorized(input, getChars(PrepareBan("")), colSLice, *color, colorAll))
		fmt.Printf("color flag is set to  %v\n", *color)
		return
	}
	if *output != "default" {
		err := os.WriteFile(*output, []byte(makeArt(input, getChars(PrepareBan(bannerStyle)))), 0644)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(makeArt(input, getChars(PrepareBan(bannerStyle))))
		fmt.Printf("Output has been saved to %v\n", *output)
		return
	}
	if *align != "default" {
		fmt.Printf("align flag is set to  %v\n", *align)
		return
	} else {
		fmt.Println(makeArt(input, getChars(PrepareBan(bannerStyle))))
	}
}
