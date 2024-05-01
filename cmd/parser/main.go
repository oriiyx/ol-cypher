package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/oriiyx/ol-cypher/parser"
)

func main() {
	filePathLocation := "./data/dump.txt"
	fileOutputLocation := "./output"

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("\033[32mOpen Library Cypher .txt Dump Parser\033[0m")
	fmt.Println("If you don't have a .txt dump, you can download it from https://openlibrary.org/data/ol_dump_latest.txt.gz")
	fmt.Println("\033[31mInput your dump.txt file location.\033[0m (e.g. './data/dump.txt')")
	fmt.Println("Default location is './data/dump.txt")
	fmt.Println("Press return to skip.")
	fmt.Println("---------------------")

	fmt.Print(": ")
	filePathLocationRawInput, _ := reader.ReadString('\n')
	// convert CRLF to LF
	filePathLocationInput := strings.Replace(filePathLocationRawInput, "\n", "", -1)

	if len(filePathLocationInput) > 0 {
		filePathLocation = filePathLocationInput
	}
	fmt.Println("File path location: \033[31m", filePathLocation, "\033[0m")
	fmt.Println("---------------------")

	fmt.Println("\033[31mInput your output folder location.\033[0m (e.g. './output')")
	fmt.Println("Default location is './output'")
	fmt.Println("Press return to skip.")
	fmt.Println("---------------------")

	fmt.Print(": ")
	fileOutputLocationRawInput, _ := reader.ReadString('\n')
	// convert CRLF to LF
	fileOutputLocationInput := strings.Replace(fileOutputLocationRawInput, "\n", "", -1)

	if len(fileOutputLocationInput) > 0 {
		fileOutputLocation = fileOutputLocationInput
	}
	fmt.Println("File output location: \033[31m", fileOutputLocation, "\033[0m")

	hideCursor()
	defer showCursor() // Ensure the cursor is shown again after the program exits

	err := parser.Parse(filePathLocation, fileOutputLocation)
	if err != nil {
		panic(err)
	}
}

func hideCursor() {
	fmt.Print("\033[?25l") // Hide cursor
}

func showCursor() {
	fmt.Print("\033[?25h") // Show cursor
}
