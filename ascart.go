package main

import (
	"flag"
	"fmt"
	"github.com/gary-norman/ascii-art/api"
	"github.com/gary-norman/ascii-art/pkg"
	"os"
	"os/exec"
)

func main() {
	//  flag definitions
	reverse := flag.String("reverse", "default",
		"Convert ascii art from a specified file into a string of characters.")
	skipsFlagReverse := pkg.CheckFlagSkipsEquals("--reverse=")
	color := flag.String("color", "default",
		"Format the output into a specified colour, either the entire text, or limited to specified characters.")
	skipsFlagColor := pkg.CheckFlagSkipsEquals("--color=")
	output := flag.String("output", "default",
		"Save the output to the specified filename")
	skipsFlagOutput := pkg.CheckFlagSkipsEquals("--output=")
	align := flag.String("align", "default",
		"Align the output to a specified alignment.")
	skipsFlagAlign := pkg.CheckFlagSkipsEquals("--align=")
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
		if skipsFlagColor {
			fmt.Println("Usage: go run . [OPTION] [STRING]\n\nEX: go run . --color=<color> <letters to be colored> \"something\"")
		}
		if len(additionalArgs) == 2 {
			colorAll = false
			colored = additionalArgs[0]
			colSLice = []rune(colored)
			input = additionalArgs[1]
		}
		fmt.Println(api.MakeArtColorized(input, api.GetChars(pkg.PrepareBanner("")), colSLice, *color, colorAll))
		return
	}
	if *output != "default" {
		if skipsFlagOutput {
			fmt.Println("Usage: go run . [OPTION] [STRING] [STYLE]\n\nEX: go run . --output=<filename> \"something\" standard")
			return
		}
		err := os.WriteFile(*output, []byte(api.MakeArt(input, api.GetChars(pkg.PrepareBanner(bannerStyle)))+"\n"), 0644)
		if err != nil {
			fmt.Println("Error writing to the file:", err)
			return // Exit the program on error
		}
		fmt.Println(api.MakeArt(input, api.GetChars(pkg.PrepareBanner(bannerStyle))))
		fmt.Printf("Output has been saved to %v\n", *output)
		return
	}
	if *align == "left" {
		if skipsFlagAlign {
			fmt.Println("Usage: go run .  [OPTION] [STRING] [BANNER]\n\nExample: go run . --align=right  something  standard")
			return
		}
		fmt.Println(api.MakeArt(input, api.GetChars(pkg.PrepareBanner(bannerStyle))))
		return
	}
	if *align == "right" {
		if skipsFlagAlign {
			fmt.Println("Usage: go run .  [OPTION] [STRING] [BANNER]\n\nExample: go run . --align=right  something  standard")
			return
		}
		ws := pkg.GetWinSize()
		ds := api.GetArtWidth(input, api.GetCharsWidth(pkg.PrepareBanner(bannerStyle)))
		fmt.Println(api.MakeArtAligned(input, api.GetChars(pkg.PrepareBanner(bannerStyle)), ds, ws, 1))
		return
	}
	if *align == "center" {
		if skipsFlagAlign {
			fmt.Println("Usage: go run .  [OPTION] [STRING] [BANNER]\n\nExample: go run . --align=right  something  standard")
			return
		}
		ws := pkg.GetWinSize()
		ds := api.GetArtWidth(input, api.GetCharsWidth(pkg.PrepareBanner(bannerStyle)))
		fmt.Println(api.MakeArtAligned(input, api.GetChars(pkg.PrepareBanner(bannerStyle)), ds, ws, 2))
		return
	}
	if *align == "justify" {
		if skipsFlagAlign {
			fmt.Println("Usage: go run .  [OPTION] [STRING] [BANNER]\n\nExample: go run . --align=right  something  standard")
			return
		}
		ws := pkg.GetWinSize()
		ds := api.GetArtWidth(input, api.GetCharsWidth(pkg.PrepareBanner(bannerStyle)))
		fmt.Println(api.MakeArtJustified(input, api.GetChars(pkg.PrepareBanner(bannerStyle)), ds, ws))
		return
	}
	if *reverse != "default" {
		if skipsFlagReverse {
			fmt.Println("Usage: go run . [OPTION]\n\nEX: go run . --reverse=<fileName>")
			return
		}
		file, err := os.Open(*reverse)
		if err != nil {
			fmt.Println("Error opening the file:", err)
			return // Exit the program on error
		}
		defer func(file *os.File) {
			err := file.Close()
			if err != nil {
				fmt.Println("Error closing the file:", err)
			}
		}(file)
		source := pkg.FileToVariable(file)
		emptyCols := api.RemoveValidSPaceIndex(api.GetEmptyCols(source))
		charMap := api.GetInputChars(api.ArtToSingleLine(source), emptyCols)
		mapStandard := api.GetChars(pkg.PrepareBanner("standard"))
		mapShadow := api.GetChars(pkg.PrepareBanner("shadow"))
		mapThinkertoy := api.GetChars(pkg.PrepareBanner("thinkertoy"))
		api.AsciiToChars(charMap, mapStandard, mapShadow, mapThinkertoy)
	}
	// test is for testing and debugging
	if *test {
		fmt.Println("Reserved for testing and debugging.")
	} else {
		// default output
		if len(additionalArgs) > 2 {
			fmt.Println("Usage: go run . [STRING] [STYLE] (optional)\n\nEX: go run . \"something\" thinkertoy.\nAvailable styles are standard, shadow, and thinkertoy.")
			return
		}
		fmt.Println(api.MakeArt(input, api.GetChars(pkg.PrepareBanner(bannerStyle))))
	}
}
