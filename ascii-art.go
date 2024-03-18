package main

import (
	// "flag"
	"fmt"
	"log"
	"os"
	"strings"
)

/**
* TODO reverse

* TODO fs

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
	// save argument as a variable
	text := os.Args[1]
	style := "standard"
	if len(os.Args) == 3 {
		style = os.Args[2]
	}
	textarr := strings.Split(text, "\\n")
	// import file of target ascii
	file, err := os.ReadFile("ascii_styles/" + style + ".txt")
	if err != nil {
		log.Fatal(err)
	}
	// assign this to a variable, split by endlines
	source := strings.Split(string(file), "\n")
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
	// reverse := flag.Bool("reverse", false, "Tell the program to run the reverse function")
	// color := flag.Bool("color", false, "Tell the program to run the color function")
	// output := flag.Bool("output", false, "Tell the program to run the output function")
	// align := flag.Bool("align", false, "Tell the program to run the align function")
	PrintAscii()
}
