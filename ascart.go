package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"

	// "regexp"
	"strings"
)

/**
* TODO reverse

* ^^ fs
	* ^^ SOLVED functionality works with shadow and standard but not with thinkertoy

* ^^ color

* ^^ output

* TODO align

*/

// PrepareBan import standard.txt as the default ascii style, with ability to change it
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
	errClose := file.Close()
	if errClose != nil {
		return nil
	}
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
				line = line + y[int(letter)][j]
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
	default:
		fmt.Print("\nAvailable colors are " + "\033[31m" + "red" + reset + ", " + "\033[32m" + "green" + reset + ", " + "\033[33m" + "yellow" + reset + ", and " + "\033[34m" + "blue" + reset + ".\n")
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
					if Contains(letters, letter) {
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
	return art
}

func main() {
	// ? flag definitions
	reverse := flag.String("reverse", "default", "Convert ascii art from a specified file into a string of characters.")
	color := flag.String("color", "default", "Format the output into a specified colour, either the entire text, or limited to specified characters.")
	output := flag.String("output", "default", "Save the output to the specified filename")
	align := flag.String("align", "default", "Align the output to a specified alignment.")
	help := flag.Bool("help", false, "Provide the user with a help file.")
	flag.Parse()                  // parse the flags so that they can be used
	additionalArgs := flag.Args() // tell the program to treat every argument following the flag from arg[o]
	var input string
	if len(additionalArgs) > 0 {
		input = additionalArgs[0]
	}
	var bannerStyle string
	if len(additionalArgs) == 2 {
		bannerStyle = additionalArgs[1]
	}
	// TODO complete reverse project
	if *reverse != "default" {
		fmt.Printf("Reverse flag is set to  %v\n", *reverse)
		return
	}

	if *help {
		cmd := exec.Command("cat", "help.txt")
		cmd.Stdout = os.Stdout
		if err := cmd.Run(); err != nil {
			fmt.Println("could not run command: ", err)
		}
	}

	if *color != "default" {
		var colored string
		var colorAll bool
		var colSLice []rune
		colorAll = true
		if len(additionalArgs) == 3 {
			fmt.Println("Usage: go run . [OPTION] [STRING]\n\nEX: go run . --color=<color> <letters to be colored> \"something\"")
		}
		if len(additionalArgs) == 2 {
			colorAll = false
			colored = additionalArgs[0]
			colSLice = []rune(colored)
			input = additionalArgs[1]
		}
		fmt.Println(makeArtColorized(input, getChars(PrepareBan("")), colSLice, *color, colorAll))
		return
	}
	if *output != "default" {
		// TODO add error/instruction message
		err := os.WriteFile(*output, []byte(makeArt(input, getChars(PrepareBan(bannerStyle)))), 0644)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(makeArt(input, getChars(PrepareBan(bannerStyle))))
		fmt.Printf("Output has been saved to %v\n", *output)
		return
	}
	// TODO complete align project
	if *align != "default" {
		fmt.Printf("align flag is set to  %v\n", *align)
		return
	} else {
		// TODO add error/instruction message
		fmt.Println(makeArt(input, getChars(PrepareBan(bannerStyle))))
	}
}
