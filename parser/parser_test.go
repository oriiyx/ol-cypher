package parser

import (
	"os"
	"testing"
)

// Expected nonexistent.txt to not exist
func TestOpenFile(t *testing.T) {
	var file *os.File
	err := openFile("nonexistent.txt", &file)
	if err == nil {
		t.Errorf("Expected error opening nonexistent file")
	}
}

// Expected test.txt to exist
func TestOpenFileExistingFile(t *testing.T) {
	var file *os.File
	err := openFile("../data/test.txt", &file)
	if err != nil {
		t.Errorf("Expected test.txt to exist")
	}
}

// Test insertJson function	with author data
func TestInsertJsonAuthor(t *testing.T) {
	dataMap := make(map[string][]map[string]interface{})
	jsonString := `{"key": "/authors/OL1A", "type": {"key": "/type/author"}, "name": "George Orwell"}`
	err := insertJson(AuthorTypeKey, dataMap, jsonString)
	if err != nil {
		t.Errorf("Expected no error")
	}
	if len(dataMap["author"]) != 1 {
		t.Errorf("Expected 1 author")
	}
}
