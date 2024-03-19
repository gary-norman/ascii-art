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
	if len(os.Args) > 3 {
		fmt.Println("Usage: go run . [STRING] [BANNER]\n\nEX: go run . something standard")
		return
	}
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
		fmt.Printf("output flag is set to  %t\n", *output)
		return
	}
	if *align {
		fmt.Printf("align flag is set to  %t\n", *align)
	} else {
		PrintAscii()
	}
}
