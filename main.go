package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func main() {

	byteCountPtr := flag.Bool("c", false, "Count the bytes of a file.")

	flag.Parse()

	filePath := flag.Args()[0]

	if filePath == "" {
		fmt.Println("fatal: file path is required")
		os.Exit(1)
	}

	if *byteCountPtr {
		byteCount := countBytes(filePath)
		fmt.Println(fmt.Sprintf("%v\t%v", byteCount, filePath))
		return
	}

	fmt.Println("Value of count option:", *byteCountPtr)
}

func countBytes(filePath string) int {

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("error opening file:", filePath, " err: ", err)
		os.Exit(1)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanBytes)

	counter := 0

	for scanner.Scan() {
		counter += len(scanner.Bytes())
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error reading input:", err)
	}

	return counter
}
