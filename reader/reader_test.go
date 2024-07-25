package reader_test

import (
	"testing"

	"github.com/oriiyx/ol-cypher/models"
	"github.com/oriiyx/ol-cypher/reader"
)

func TestReader_ReadRatings(t *testing.T) {
	r := reader.Reader{
		RatingJsonLocation: "../output/ratings.json",
	}

	rating1, rating2, rating3, rating4, rating5 := 1, 2, 3, 4, 5

	ratingsChannel := make(chan models.Ratings)
	go r.ReadRatings(ratingsChannel)
	for rating := range ratingsChannel {
		if rating.Rating != rating1 && rating.Rating != rating2 && rating.Rating != rating3 && rating.Rating != rating4 && rating.Rating != rating5 {
			t.Error("Rating rating is empty")
		}

		if rating.BookId == "" {
			t.Error("Rating book is empty")
		}

		if rating.RatingDate == "" {
			t.Error("Rating date is empty")
		}
	}
}

func TestReader_ReadEditions(t *testing.T) {
	r := reader.Reader{
		EditionJsonLocation: "../output/test-edition.json",
	}

	editionChannel := make(chan models.Edition)
	go r.ReadEditions(editionChannel)
	for edition := range editionChannel {
		if edition.Key == "" {
			t.Error("Edition key is empty")
		}
		if edition.Title == "" {
			t.Error("Edition title is empty")
		}
		if edition.ISBN10[0] == "" {
			t.Error("Edition ISBN10 is empty")
		}
	}
}

func TestReader_ReadWorks(t *testing.T) {
	r := reader.Reader{
		WorkJsonLocation: "../output/test-work.json",
	}

	workChannel := make(chan models.Work)
	go r.ReadWorks(workChannel)
	for work := range workChannel {
		if work.Key == "" {
			t.Error("Work key is empty")
		}
		if work.Title == "" {
			t.Error("Work title is empty")
		}
	}
}

func TestReader_ReadAuthors(t *testing.T) {
	r := reader.Reader{
		AuthorJsonLocation: "../output/test-author.json",
	}

	authorChannel := make(chan models.Author)
	go r.ReadAuthors(authorChannel)
	for author := range authorChannel {
		if author.Key == "" {
			t.Error("Author key is empty")
		}
		if author.Name == "" {
			t.Error("Author name is empty")
		}
	}
}
