package reader

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/oriiyx/ol-cypher/models"
)

type Reader struct {
	AuthorJsonLocation  string
	WorkJsonLocation    string
	EditionJsonLocation string
	RatingJsonLocation  string
}

// ReadRatings streams authors from the JSON file specified in RatingJsonLocation.
// It sends each author to the provided channel and closes the channel when done.
func (r *Reader) ReadRatings(ch chan<- models.Ratings) {
	file, err := os.Open(r.RatingJsonLocation)
	if err != nil {
		close(ch)
		fmt.Println("Error opening JSON file:", err)
		return
	}
	defer file.Close()

	decoder := json.NewDecoder(file)

	// Read opening bracket of the JSON array
	if _, err := decoder.Token(); err != nil {
		close(ch)
		fmt.Println("Error reading start of JSON:", err)
		return
	}

	for decoder.More() {
		var rating models.Ratings
		err := decoder.Decode(&rating)
		if err != nil {
			fmt.Println("Error decoding JSON:", err)
			continue
		}
		ch <- rating
	}

	close(ch)
}

// ReadAuthors streams authors from the JSON file specified in AuthorJsonLocation.
// It sends each author to the provided channel and closes the channel when done.
func (r *Reader) ReadAuthors(ch chan<- models.Author) {
	file, err := os.Open(r.AuthorJsonLocation)
	if err != nil {
		close(ch)
		fmt.Println("Error opening JSON file:", err)
		return
	}
	defer file.Close()

	decoder := json.NewDecoder(file)

	// Read opening bracket of the JSON array
	if _, err := decoder.Token(); err != nil {
		close(ch)
		fmt.Println("Error reading start of JSON:", err)
		return
	}

	for decoder.More() {
		var author models.Author
		err := decoder.Decode(&author)
		if err != nil {
			fmt.Println("Error decoding JSON:", err)
			continue
		}
		ch <- author
	}

	close(ch)
}

func (r *Reader) ReadWorks(ch chan<- models.Work) {
	file, err := os.Open(r.WorkJsonLocation)
	if err != nil {
		close(ch)
		fmt.Println("Error opening JSON file:", err)
		return
	}
	defer file.Close()

	decoder := json.NewDecoder(file)

	// Read opening bracket of the JSON array
	if _, err := decoder.Token(); err != nil {
		close(ch)
		fmt.Println("Error reading start of JSON:", err)
		return
	}

	for decoder.More() {
		var work models.Work
		err := decoder.Decode(&work)
		if err != nil {
			fmt.Println("Error decoding JSON:", err)
			continue
		}
		ch <- work
	}

	close(ch)
}

func (r *Reader) ReadEditions(ch chan<- models.Edition) {
	file, err := os.Open(r.EditionJsonLocation)
	if err != nil {
		close(ch)
		fmt.Println("Error opening JSON file:", err)
		return
	}
	defer file.Close()

	decoder := json.NewDecoder(file)

	// Read opening bracket of the JSON array
	if _, err := decoder.Token(); err != nil {
		close(ch)
		fmt.Println("Error reading start of JSON:", err)
		return
	}

	for decoder.More() {
		var edition models.Edition
		err := decoder.Decode(&edition)
		if err != nil {
			fmt.Println("Error decoding JSON:", err)
			continue
		}
		ch <- edition
	}

	close(ch)
}
