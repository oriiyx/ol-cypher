package cypher

import (
	"fmt"

	"github.com/oriiyx/ol-cypher/models"
)

type Cypher struct {
	key int
}

func GetAuthors() models.Author {
	fmt.Println("Authors: John Doe, Jane Doe")
	return models.Author{
		Name: "John Doe",
		Key:  "john-doe",
	}
}
