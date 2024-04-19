package pkg

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"syscall"
	"unsafe"
)

// Winsize is a struct that stores the height and width of the terminal.
type Winsize struct {
	Row    uint16
	Col    uint16
	Xpixel uint16 // unused
	Ypixel uint16 // unused
}

// GetWinSize populates the Winsize structure
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

// PrepareBanner import standard.txt as the default ascii style, with ability to change it using 2nd argument *
func PrepareBanner(bannerStyle string) []string {
	if bannerStyle == "" {
		bannerStyle = "standard"
	}
	style := bannerStyle
	file, err := os.Open("ascii_styles/" + style + ".txt")
	if err != nil {
		fmt.Println("Error opening the file:", err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("Error closing the file:", err)
		}
	}(file)
	scanned := bufio.NewScanner(file)
	scanned.Split(bufio.ScanLines)
	var source []string
	for scanned.Scan() {
		source = append(source, scanned.Text())
	}
	return source
}

// FileToVariable takes a file as input and returns it as a slice of strings
func FileToVariable(file *os.File) []string {
	scanned := bufio.NewScanner(file)
	scanned.Split(bufio.ScanLines)
	var source []string
	for scanned.Scan() {
		source = append(source, scanned.Text())
	}
	return source
}

// CompareSlices compares two slices for equality.
func CompareSlices(slice1, slice2 []string) bool {
	if len(slice1) != len(slice2) {
		return false // Slices are of different lengths
	}
	for i, v := range slice1 {
		if v != slice2[i] {
			return false // Elements at the same position are different
		}
	}
	return true // Slices are equal
}

// SliceContainsItem checks if a slice contains a specific element.
func SliceContainsItem(slice []rune, item rune) bool {
	for _, v := range slice {
		if v == item {
			return true // Found the item
		}
	}
	return false // Item not found
}

// CheckFlagSkipsEquals checks if the value passed to the flag uses <flag>=<argument> format
func CheckFlagSkipsEquals(flag string) bool {
	usedEqualsSyntax := true
	for _, arg := range os.Args[1:] {
		// Check if the argument starts with the flag name followed by "="
		if strings.HasPrefix(arg, flag) {
			usedEqualsSyntax = false
		}
	}
	return usedEqualsSyntax
}
