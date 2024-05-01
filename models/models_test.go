package models

import (
	"testing"
)

func TestAuthor(t *testing.T) {
	author := Author{
		Name: "John Doe",
		Key:  "john-doe",
	}

	if author == (Author{}) {
		t.Errorf("Author should not be empty")
	}

	if author.Name != "John Doe" {
		t.Errorf("Expected author name to be 'John Doe'. Got %s", author.Name)
	}

	if author.Key != "john-doe" {
		t.Errorf("Expected author key to be 'john-doe'. Got %s", author.Key)
	}
}
