package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {

	byteCountPtr := flag.Bool("c", false, "Counts the bytes of a file.")
	lineCountPtr := flag.Bool("l", false, "Counts the lines of a file.")
	wordCountPtr := flag.Bool("w", false, "Counts the words of a file.")
	charCountPtr := flag.Bool("m", false, "Counts the characters of a file.")

	flag.Parse()

	filePath := flag.Args()[0]
	areFlagsSet := flag.NFlag() > 0

	var input io.Reader

	if filePath != "" {
		file, err := os.Open(filePath)
		if err != nil {
			fmt.Println("rwc: error opening file: ", err)
			os.Exit(1)
		}

		defer file.Close()

		input = file
	} else {
		fmt.Println("fatal: file path is required")
		os.Exit(1)
	}

	lineCount, wordCount, charCount, byteCount := count(input)

	if *lineCountPtr || !areFlagsSet {
		fmt.Printf("%v\t", lineCount)
	}

	if *wordCountPtr || !areFlagsSet {
		fmt.Printf("%v\t", wordCount)
	}

	if *byteCountPtr || !areFlagsSet {
		fmt.Printf("%v\t", byteCount)
	}

	if *charCountPtr {
		fmt.Printf("%v\t", charCount)
	}

	fmt.Printf("%v\n", filePath)
}

func count(input io.Reader) (int, int, int, int) {
	scanner := bufio.NewScanner(input)

	var lineCount, wordCount, charCount, byteCount int

	for scanner.Scan() {
		lineCount++
		byteCount += len(scanner.Bytes()) + 2
		charCount += len(scanner.Text()) + 2
		wordCount += len(strings.Split(scanner.Text(), " "))
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error reading input:", err)
	}

	return lineCount, wordCount, charCount, byteCount
}
