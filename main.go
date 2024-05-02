package main

import (
	"github.com/oriiyx/ol-cypher/models"
	"github.com/oriiyx/ol-cypher/reader"
)

func main() {
	// Create a new Reader instance with the path to your JSON file.
	r := reader.Reader{
		AuthorJsonLocation:  "./output/author.json",
		WorkJsonLocation:    "./output/work.json",
		EditionJsonLocation: "./output/edition.json",
	}

	authorChannel := make(chan models.Author)
	go r.ReadAuthors(authorChannel)
	for author := range authorChannel {
		_ = author
	}

	workChannel := make(chan models.Work)
	go r.ReadWorks(workChannel)
	for work := range workChannel {
		_ = work
	}

	editionChannel := make(chan models.Edition)
	go r.ReadEditions(editionChannel)
	for edition := range editionChannel {
		_ = edition
	}
}
