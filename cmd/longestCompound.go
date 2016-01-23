package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/NodePrime/quiz"
)

func main() {
	var stdin, verbose bool
	var filename string
	flag.BoolVar(&stdin, "stdin", false, "Take user input from stdin")
	flag.StringVar(&filename, "infile", "", "Input file to read user input from")
	flag.BoolVar(&verbose, "verbose", false, "Verbose output including decomposition (false by default)")
	flag.Parse()
	if (!stdin && filename == "") || (filename != "" && stdin) {
		flag.Usage()
		fmt.Println("Please choose either -stdin or -filename {non empy filename}")
		return
	}
	// open input
	input := os.Stdin
	if filename != "" {
		file, err := os.Open("word.list")
		dieOnError(err)
		input = file
	}
	// compute
	words, err := quiz.ToWords(input)
	dieOnError(err)
	longestCompound, subwords, err := quiz.LongestCompoundWord(words)
	dieOnError(err)
	if verbose {
		fmt.Printf("'%s' -> %v (%d)\n", longestCompound, subwords, len(subwords))
	} else {
		fmt.Println(longestCompound)
	}
}

// dieOnError displays error and ends the program
func dieOnError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
