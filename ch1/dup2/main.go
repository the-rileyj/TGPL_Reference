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
	counts := make(map[string]map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, "stdin")
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, arg)
			f.Close()
		}
	}
	for file, ma := range counts {
		for _, n := range ma {
			if n > 1 {
				fmt.Printf("%s\n", file)
				break
			}
		}
	}
}

func countLines(f *os.File, counts map[string]map[string]int, fname string) {
	input := bufio.NewScanner(f)
	counts[fname] = make(map[string]int)
	for input.Scan() {
		counts[fname][input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Err()
}

//!-
