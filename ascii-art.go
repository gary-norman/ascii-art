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

// // save the first argument as an array equal in size to the wordcount
// func prepareArg(text string) []string {
// 	textarr := strings.Split(text, "\\n")
// 	textarr = removeEmptyStrings(textarr)
// 	return textarr
// }

// func PrintAscii(bannerStyle string) {
// 	if len(os.Args) > 3 {
// 		fmt.Println("Usage: go run . [STRING] [BANNER]\n\nEX: go run . something standard")
// 		return
// 	}
// 	source, textarr := PrepareBan(bannerStyle), prepareArg()
// 	for i := 0; i < len(textarr); i++ {
// 		for j := 1; j < 10; j++ {
// 			for _, char := range textarr[i] {
// 				fmt.Print(source[(int(char)-(32))*9+(j)])
// 			}
// 			fmt.Println("")
// 		}
// 	}
// }

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

	// ? flag definitions
	reverse := flag.Bool("reverse", false, "Tell the program to run the reverse function")
	color := flag.Bool("color", false, "Tell the program to run the color function")
	output := flag.String("output", "default", "Save the output to the specified filename")
	align := flag.Bool("align", false, "Tell the program to run the align function")
	flag.Parse()
	additionalArgs := flag.Args()
	input := additionalArgs[0]
	var bannerStyle string
	if len(additionalArgs) == 2 {
		bannerStyle = additionalArgs[1]
	}
	if *reverse {
		fmt.Printf("Reverse flag is set to  %t\n", *reverse)
		return
	}
	if *color {
		fmt.Printf("color flag is set to  %t\n", *color)
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
	if *align {
		fmt.Printf("align flag is set to  %t\n", *align)
	} else {
		fmt.Println(makeArt(input, getChars(PrepareBan(bannerStyle))))
	}
}
