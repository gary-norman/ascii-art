package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

/**
* TODO reverse

* ^^ fs
	* ^^ SOLVED functionality works with shadow and standard but not with thinkertoy

* TODO color

* TODO output

* TODO align

*/

// removeEmptyStrings - Use this to remove empty string values inside an array.
func removeEmptyStrings(s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}

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

// save the first argument as an array equal in size to the wordcount
func prepareArg() []string {
	text := os.Args[1]
	textarr := strings.Split(text, "\\n")
	textarr = removeEmptyStrings(textarr)
	return textarr
}

func PrintAscii(bannerStyle string) {
	if len(os.Args) > 3 {
		fmt.Println("Usage: go run . [STRING] [BANNER]\n\nEX: go run . something standard")
		return
	}
	source, textarr := PrepareBan(bannerStyle), prepareArg()
	for i := 0; i < len(textarr); i++ {
		for j := 1; j < 10; j++ {
			for _, char := range textarr[i] {
				fmt.Print(source[(int(char)-(32))*9+(j)])
			}
			fmt.Println("")
		}
	}
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

func main() {
	fileOutput := "filename.txt"
	input := os.Args[1]
	var bannerStyle string
	if len(os.Args) == 3 {
		bannerStyle = os.Args[2]
	}
	// ? flag definitions
	reverse := flag.Bool("reverse", false, "Tell the program to run the reverse function")
	color := flag.Bool("color", false, "Tell the program to run the color function")
	output := flag.Bool("output", false, "Tell the program to run the output function")
	align := flag.Bool("align", false, "Tell the program to run the align function")
	flag.Parse()
	if *reverse {
		fmt.Printf("Reverse flag is set to  %t\n", *reverse)
		return
	}
	if *color {
		fmt.Printf("color flag is set to  %t\n", *color)
		return
	}
	if *output {
		err := os.WriteFile(fileOutput, string(makeArt(input, getChars(PrepareBan(bannerStyle)))), 0644)
		fmt.Println(makeArt(input, getChars(PrepareBan(bannerStyle))))
		return
	}
	if *align {
		fmt.Printf("align flag is set to  %t\n", *align)
	} else {
		fmt.Println(makeArt(input, getChars(PrepareBan(bannerStyle))))
	}
}
