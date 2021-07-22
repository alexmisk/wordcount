// match tool checks a string against a pattern.
// If it matches - prints the string, otherwise prints nothing.
package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	// src := readInput()
	src := strings.Join(os.Args[1:], " ")
	words := wc(src)
	fmt.Println(words)
}

// match returns true if src matches pattern,
// false otherwise.
func wc(src string) int {
	re := regexp.MustCompile(`[\S]+`)

	results := re.FindAllString(src, -1)
	return len(results)
}

// readInput reads pattern and source string
// from command line arguments and returns them.
func readInput() (src string) {
	src = strings.Join(flag.Args(), "")
	return src
}
