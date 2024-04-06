package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {

	byteCountPtr := flag.Bool("c", false, "Counts the bytes of a file.")
	lineCountPtr := flag.Bool("l", false, "Counts the lines of a file.")
	wordCountPtr := flag.Bool("w", false, "Counts the words of a file.")

	flag.Parse()

	filePath := flag.Args()[0]

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

	if *byteCountPtr {
		byteCount := countBySplit(input, bufio.ScanBytes)
		fmt.Println(fmt.Sprintf("%v\t%v", byteCount, filePath))
		return
	}

	if *lineCountPtr {
		lineCount := countBySplit(input, bufio.ScanLines)
		fmt.Println(fmt.Sprintf("%v\t%v", lineCount, filePath))
		return
	}

	if *wordCountPtr {
		wordCount := countBySplit(input, bufio.ScanWords)
		fmt.Println(fmt.Sprintf("%v\t%v", wordCount, filePath))
		return
	}
}

func countBySplit(input io.Reader, split bufio.SplitFunc) int {
	scanner := bufio.NewScanner(input)
	scanner.Split(split)

	counter := 0

	for scanner.Scan() {
		counter++
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error reading input:", err)
	}

	return counter
}
