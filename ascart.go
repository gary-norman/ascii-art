package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"reflect"
	"strings"
	"syscall"
	"unsafe"
)

/**
* TODO reverse

* ^^ fs
	* ^^ SOLVED functionality works with shadow and standard but not with thinkertoy

* ^^ color

* ^^ output

* ^^ align

*/

// Winsize * struct that stores the height and width of the terminal.
type Winsize struct {
	Row    uint16
	Col    uint16
	Xpixel uint16 // unused
	Ypixel uint16 // unused
}

// GetWinSize * populates the Winsize structure
func GetWinSize() Winsize {
	// Get the file descriptor for stdout
	fd := syscall.Stdout

	// Create an instance of Winsize
	var ws Winsize

	// Use the TIOCGWINSZ ioctl system call to get the window size
	_, _, err := syscall.Syscall(syscall.SYS_IOCTL, uintptr(fd), uintptr(syscall.TIOCGWINSZ), uintptr(unsafe.Pointer(&ws)))
	if err != 0 {
		fmt.Println("Error getting terminal size:", err)
	}
	return ws
}

// PrepareBan import standard.txt as the default ascii style, with ability to change it using 2nd argument *
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

// getChars * map the ascii characters provided in the style.txt file, indexed by ascii code
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

// FileToVariable * takes a file as input and returns it as a slice of strings
func FileToVariable(file *os.File) []string {
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

// slicesEqual * checks if two slices of strings are identical.
func slicesEqual(slice1, slice2 []string) bool {
	return reflect.DeepEqual(slice1, slice2)
}

// artToSingleLine * TODO places each line of characters from FileToVariable on a single line

// getEmptyCols * get the index of the final space of each character in the reverse flag
func getEmptyCols(source []string) []int {
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

// removeValidSPaceIndex * remove indices for valid spaces, before the end space
func removeValidSPaceIndex(indices []int) []int {
	for i := 0; i < len(indices)-1; i++ {
		if len(indices)-i > 6 {
			if indices[i] == (indices[i+6])-6 {
				indices = append(indices[:i+1], indices[i+6:]...)
			}
		}
	}
	return indices
}

// getInputChars * map the ascii characters provided in the reverse flag, zero indexed
func getInputChars(source []string, indices []int) map[int][]string {
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

// AsciiToChars * compares getChar and getInputChar and prints the string to the terminal
func AsciiToChars(input, standard, shadow, thinkertoy map[int][]string) {
	output := make(map[int][]int)
	styles := []map[int][]string{standard, shadow, thinkertoy}
	for _, style := range styles {
		for key1, slice1 := range input {
			for key2, slice2 := range style {
				if slicesEqual(slice1, slice2) {
					output[key1] = append(output[key1], key2)
				}
			}
		}
	}
	for i := 0; i < len(output); i++ {
		fmt.Printf("%c", output[i][0])
	}
}

// getCharsWidth * determine the width of each individual ascii art character
func getCharsWidth(source []string) map[int]int {
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

// GetArtWidth * determine the width of each line that gets printed to the terminal (without EOL)
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

// Contains * checks if a slice contains a specific element.
func Contains(slice []rune, item rune) bool {
	for _, v := range slice {
		if v == item {
			return true // Found the item
		}
	}
	return false // Item not found
}

// makeArt * transform the input text origString to the output art, line by line
func makeArt(origString string, y map[int][]string) string {
	var art string
	replaceNewline := strings.ReplaceAll(origString, "\r\n", "\\n") // correct newline formatting
	wordSlice := strings.Split(replaceNewline, "\\n")               //split the input into slices
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

// makArtAligned * transform the input text origString to the output art, line by line, with left, right, or center aligned content
func makeArtAligned(origString string, y map[int][]string, ds []int, ws Winsize, divider int) string {
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

// makArtJustified * transform the input text origString to the output art, line by line, with justified content
func makeArtJustified(origString string, y map[int][]string, ds []int, ws Winsize) string {
	var art string
	replaceNewline := strings.ReplaceAll(origString, "\r\n", "\\n") // correct newline formatting
	wordSlice := strings.Split(replaceNewline, "\\n")
	for i := 0; i < len(wordSlice); i++ {
		for j := 0; j < len(y[32]); j++ {
			var line string
			for _, letter := range wordSlice[i] {
				line = line + y[int(letter)][j]
				sliceLen := len(wordSlice[i])
				if sliceLen >= 2 {
					line = line + strings.Repeat(" ", (int(ws.Col)-ds[i])/(sliceLen-1))
				}
			}
			art += line + "\n"
			line = ""
		}
	}
	art = strings.TrimRight(art, "\n")
	return art
}

// makArtColorized * transform the input text origString to the output art, line by line, colorizing specified text
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
	art = strings.TrimRight(art, "\n")
	return art
}

func main() {
	// ? flag definitions
	reverse := flag.String("reverse", "default", "Convert ascii art from a specified file into a string of characters.")
	color := flag.String("color", "default", "Format the output into a specified colour, either the entire text, or limited to specified characters.")
	output := flag.String("output", "default", "Save the output to the specified filename")
	align := flag.String("align", "default", "Align the output to a specified alignment.")
	help := flag.Bool("help", false, "Provide the user with a help file.")
	test := flag.Bool("test", false, "testing")
	flag.Parse() // parse the flags so that they can be used
	// define args to be non-flag arguments
	additionalArgs := flag.Args() // tell the program to treat every argument following the flag as arg[o...]
	var input string
	if len(additionalArgs) > 0 {
		input = additionalArgs[0]
	}
	var bannerStyle string
	if len(additionalArgs) == 2 {
		bannerStyle = additionalArgs[1]
	}
	// call the functions depending on the flag
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
		if len(additionalArgs) > 2 {
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
		if len(additionalArgs) > 2 {
			fmt.Println("Usage: go run . [OPTION] [STRING] [STYLE]\n\nEX: go run . --output=<filename> \"something\" shadow.")
		}
		err := os.WriteFile(*output, []byte(makeArt(input, getChars(PrepareBan(bannerStyle)))), 0644)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(makeArt(input, getChars(PrepareBan(bannerStyle))))
		fmt.Printf("Output has been saved to %v\n", *output)
		return
	}
	// TODO * error message
	if *align == "right" {
		ws := GetWinSize()
		ds := GetArtWidth(input, getCharsWidth(PrepareBan(bannerStyle)))
		fmt.Println(makeArtAligned(input, getChars(PrepareBan(bannerStyle)), ds, ws, 1))
		return
	}
	if *align == "center" {
		ws := GetWinSize()
		ds := GetArtWidth(input, getCharsWidth(PrepareBan(bannerStyle)))
		fmt.Println(makeArtAligned(input, getChars(PrepareBan(bannerStyle)), ds, ws, 2))
		return
	}
	if *align == "justify" {
		ws := GetWinSize()
		ds := GetArtWidth(input, getCharsWidth(PrepareBan(bannerStyle)))
		fmt.Println(makeArtJustified(input, getChars(PrepareBan(bannerStyle)), ds, ws))
		return
	}
	// TODO complete reverse project
	if *reverse != "default" {
		file, err := os.Open(*reverse)
		if err != nil {
			log.Fatal(err)
		}
		source := FileToVariable(file)
		emptyCols := removeValidSPaceIndex(getEmptyCols(source))
		charMap := getInputChars(source, emptyCols)
		AsciiToChars(charMap, getChars(PrepareBan("standard")), getChars(PrepareBan("shadow")), getChars(PrepareBan("thinkertoy")))
		//asciiMap := getChars(PrepareBan("thinkertoy"))
		//fmt.Println("")
		//for k, v := range charMap {
		//	fmt.Println(k, ":", v)
		//}
		//fmt.Printf("104 %v\n", asciiMap[104])
		//fmt.Printf("101 %v\n", asciiMap[101])
		//fmt.Printf("121 %v\n", asciiMap[121])
	}
	// test is for testing and debugging
	if *test {
		fmt.Println(getChars(PrepareBan(bannerStyle)))
	} else {
		// default output
		if len(additionalArgs) > 2 {
			fmt.Println("Usage: go run . [STRING] [STYLE] (optional)\n\nEX: go run . \"something\" thinkertoy.\nAvailable styles are standard, shadow, and thinkertoy.")
		}
		fmt.Println(makeArt(input, getChars(PrepareBan(bannerStyle))))
	}
}
