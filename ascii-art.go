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

* TODO fs
	* ? functionality works with shadow and standard but not with thinkertoy

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

func PrepareInputs() ([]string, []string) {
	// import standard.txt as the default ascii style, with ability to change it
	// using 2nd argument *
	// ^^ thinkertoy does not work
	style := "standard"
	if len(os.Args) == 3 {
		style = os.Args[2]
	}
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
	// save the first argument as an array equal in size to the wordcount
	text := os.Args[1]
	textarr := strings.Split(text, "\\n")
	textarr = removeEmptyStrings(textarr)
	return source, textarr
}

func PrintAscii() {
	source, textarr := PrepareInputs()
	for i := 0; i < len(textarr); i++ {
		for j := 1; j < 10; j++ {
			for _, char := range textarr[i] {
				fmt.Print(source[(int(char)-(32))*9+(j)])
			}
			fmt.Println("")
		}
	}
}

func main() {
	// ? flag definitions
	var reverse bool
	flag.BoolVar(&reverse, "reverse", false, "Tell the program to run the reverse function")
	var color bool
	flag.BoolVar(&color, "color", false, "Tell the program to run the color function")
	var output bool
	flag.BoolVar(&output, "output", false, "Tell the program to run the output function")
	var align bool
	flag.BoolVar(&align, "align", false, "Tell the program to run the align function")
	flag.Parse()
	if reverse {
		print("reverse")
	}
	if color {
		print("color")
	}
	if output {
		print("output")
	}
	if align {
		print("align")
	} else {
		PrintAscii()
	}
}
