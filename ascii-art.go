package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

/*

strings:
--------
hello
HELLO
HeLlo HuMaN
1Hello 2There
Hello\nThere
Hello\n\nThere
{Hello & There #}
hello There 1 to 2!
MaD3IrA&LiSboN
1a\"#FdwHywR&/()=
{|}~
[\]^_ 'a
RGB
:;<=>?@
\!" #$%&'"'"'()*+,-./
ABCDEFGHIJKLMNOPQRSTUVWXYZ
abcdefghijklmnopqrstuvwxyz
<a random string> with at least four lower case letters and three upper case letters.
<a random string> with at least five lower case letters, a space and two numbers.
<a random string> with at least one upper case letters and 3 special characters.
<a random string> with at least two lower case letters, two spaces, one number, two special characters and three upper case letters.

+Does the project run quickly and effectively? (Favoring recursive, no unnecessary data requests, etc)
+Does the code obey the good practices?
+Is there a test file for this code?
+Are the tests checking each possible case?
+Is the output of the program well structured? Are the characters displayed correctly in line?
https://github.com/01-edu/public/blob/master/subjects/ascii-art/audit/README.md

*/

func main() {
	// save argument as a variable
	text := os.Args[1]
	// import file of target ascii
	file, err := os.ReadFile("standard.txt")
	if err != nil {
		log.Fatal(err)
	}
	// zero the ascii code to provide a baseline
	baseline := 32
	// assign this to a variable, split by endlines
	source := strings.Split(string(file), "\n")

	for i := 1; i < 10; i++ {
		for _, char := range text {
			fmt.Print(source[(int(char)-(baseline))*9+(i)])
		}
		fmt.Println("")
	}
}
