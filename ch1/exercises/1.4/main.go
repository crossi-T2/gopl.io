// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 10.
//!+

// Dup2 prints the count and text of lines that appear more than once
// in the input.  It reads from stdin or from a list of named files.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	// keep track of the filenames for a certain key
	filenames := make(map[string]map[string]int)

	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, filenames)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, filenames)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("line: %s\toccurrences: %d\tfilename(s): ", line, n)

			for filename, _ := range filenames[line] {
				fmt.Printf("%v ", filename)
			}
			fmt.Printf("\n")
		}
	}
}

func countLines(f *os.File, counts map[string]int, filenames map[string]map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		var line string = input.Text()

		counts[line]++

		var filenamesForLine map[string]int

		filenamesForLine = filenames[line]

		if filenamesForLine == nil {
			filenamesForLine = make(map[string]int)
		}

		filenamesForLine[f.Name()] = 1
		filenames[line] = filenamesForLine
	}
	// NOTE: ignoring potential errors from input.Err()
}

//!-
