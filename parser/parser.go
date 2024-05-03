package parser

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

const (
	AuthorTypeKey  string = "/type/author"
	WorkTypeKey    string = "/type/work"
	EditionTypeKey string = "/type/edition"
)

func Parse(filePath string, fileOutput string) error {
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

	dataMap := make(map[string][]map[string]interface{})
	batchSize := 100000 // Define batch size
	// batchSize := 2 // Define batch size

	fmt.Print("Parsing ")
	i := 0

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "\t")

		if len(parts) != 5 {
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

		dataType := parts[0]
		jsonString := parts[len(parts)-1]

		err := insertJson(dataType, dataMap, jsonString)
		if err != nil {
			return err
		}

		i++
		if i%batchSize == 0 {
			for key, value := range dataMap {
				if err := saveJSON(key, value, fileOutput); err != nil {
					return err
				}
			}
			dataMap = make(map[string][]map[string]interface{}) // Reset the map after saving
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	// Handle any remaining data
	if len(dataMap) > 0 {
		for key, value := range dataMap {
			if err := saveJSON(key, value, fileOutput); err != nil {
				return err
			}
		}
	}

	fmt.Println("")
	// paint the fmt green
	fmt.Println("\033[32mParsing finished\033[0m")
	return nil
}

func insertJson(dataType string, dataMap map[string][]map[string]interface{}, jsonString string) error {
	var tempMap map[string]interface{}

	if err := json.Unmarshal([]byte(jsonString), &tempMap); err != nil {
		return err
	}

	switch dataType {
	case AuthorTypeKey:
		dataMap["author"] = append(dataMap["author"], tempMap)
	case WorkTypeKey:
		dataMap["work"] = append(dataMap["work"], tempMap)
	case EditionTypeKey:
		dataMap["edition"] = append(dataMap["edition"], tempMap)
	default:
		// fmt.Println("Unknown data type:", dataType)
	}

	return nil
}

func openFile(filePath string, file **os.File) error {
	var err error
	*file, err = os.Open(filePath)
	if err != nil {
		return err
	}

	return nil
}

func saveJSON(dataType string, data []map[string]interface{}, filePath string) error {
	fileName := fmt.Sprintf("%s.json", dataType)
	fullPath := fmt.Sprintf("%s/%s", filePath, fileName)

	file, err := os.OpenFile(fullPath, os.O_RDWR|os.O_CREATE, 0o644)
	if err != nil {
		return err
	}
	defer file.Close()

	// Determine if file is new or already has data
	stat, err := file.Stat()
	if err != nil {
		return err
	}
	startNewFile := stat.Size() == 0

	var dataBytes []byte
	if startNewFile {
		// Initialize a new JSON array
		dataBytes, err = json.Marshal(data)
		if err != nil {
			return err
		}
		_, err = file.Write([]byte("[" + string(dataBytes)[1:len(dataBytes)-1]))
	} else {
		// Seek to the end and adjust for appending
		_, err = file.Seek(-1, io.SeekEnd) // Move back before the last ']'
		if err != nil {
			return err
		}
		dataBytes, err = json.Marshal(data)
		if err != nil {
			return err
		}
		_, err = file.Write([]byte("," + string(dataBytes)[1:len(dataBytes)-1]))
	}
	if err != nil {
		return err
	}

	// Close the JSON array correctly
	if _, err = file.WriteString("]"); err != nil {
		return err
	}

	return nil
}
