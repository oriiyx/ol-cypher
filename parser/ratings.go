package parser

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type RatingsJson struct {
	Work       string
	Edition    string
	Rating     int
	RatingDate string
}

func ParseRatings(filePath string, fileOutput string) error {
	fmt.Println("Parsing the .txt dump file inside ", filePath, " folder")
	var file *os.File
	err := openFile(filePath, &file)
	if err != nil {
		return err
	}
	defer file.Close()

	fmt.Println("Successfully opened file:", file.Name())

	scanner := bufio.NewScanner(file)
	buf := make([]byte, 0, 64*1024)
	scanner.Buffer(buf, 16384*1024)

	fmt.Print("Parsing ")
	i := 0
	var ratingData []RatingsJson

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "\t")

		if len(parts) != 4 {
			continue
		}

		// Reset the dots after three to keep the length consistent
		if i%150000 == 0 {
			fmt.Print("\rParsing        ") // Ensure enough spaces to clear previous output
			fmt.Print("\rParsing ")        // Reset cursor position for dots
		}

		if i%50000 == 0 {
			fmt.Print(".")
		}

		rating, err := strconv.Atoi(parts[2])
		if err != nil {
			continue
		}

		ratingJson := RatingsJson{
			Work:       parts[0],
			Edition:    parts[1],
			Rating:     rating,
			RatingDate: parts[3],
		}

		ratingData = append(ratingData, ratingJson)
		i++
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	fileName := "ratings.json"
	fullPath := fmt.Sprintf("%s/%s", fileOutput, fileName)

	marshalJson, err := json.Marshal(ratingData)
	err = os.WriteFile(fullPath, marshalJson, 0o644)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("")
	// paint the fmt green
	fmt.Println("\033[32mParsing finished\033[0m")
	return nil
}
