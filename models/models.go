package models

import "encoding/json"

type Author struct {
	Created struct {
		Type  string `json:"type"`
		Value string `json:"value"`
	} `json:"created"`
	Key          string `json:"key"`
	LastModified struct {
		Type  string `json:"type"`
		Value string `json:"value"`
	} `json:"last_modified"`
	LatestRevision int      `json:"latest_revision"`
	Name           string   `json:"name"`
	Revision       int      `json:"revision"`
	SourceRecords  []string `json:"source_records"`
	Type           struct {
		Key string `json:"key"`
	} `json:"type"`

	// Optional fields
	AlternateNames []string `json:"alternate_names,omitempty"`
	// Bio            *struct {
	// 	Type  string `json:"type"`
	// 	Value string `json:"value"`
	// } `json:"bio,omitempty"`
	Bio          json.RawMessage `json:"bio,omitempty"`
	BirthDate    *string         `json:"birth_date,omitempty"`
	DeathDate    *string         `json:"death_date,omitempty"`
	PersonalName *string         `json:"personal_name,omitempty"`
	Photos       []int           `json:"photos,omitempty"`
	RemoteIDs    *struct {
		ISNI     string `json:"isni,omitempty"`
		VIAF     string `json:"viaf,omitempty"`
		Wikidata string `json:"wikidata,omitempty"`
	} `json:"remote_ids,omitempty"`
}

type Work struct {
	Authors []struct {
		Author struct {
			Key string `json:"key"`
		} `json:"author"`
		Type json.RawMessage `json:"type"`
	} `json:"authors"`
	Covers  []int `json:"covers"`
	Created struct {
		Type  string `json:"type"`
		Value string `json:"value"`
	} `json:"created"`
	Key          string `json:"key"`
	LastModified struct {
		Type  string `json:"type"`
		Value string `json:"value"`
	} `json:"last_modified"`
	LatestRevision int    `json:"latest_revision"`
	Revision       int    `json:"revision"`
	Title          string `json:"title"`
	Type           struct {
		Key string `json:"key"`
	} `json:"type"`

	// Optional fields
	Description      json.RawMessage `json:"description,omitempty"`
	FirstPublishDate *string         `json:"first_publish_date,omitempty"`
	FirstSentence    *struct {
		Type  string `json:"type"`
		Value string `json:"value"`
	} `json:"first_sentence,omitempty"`
	Subjects []string `json:"subjects,omitempty"`
}

type Edition struct {
	Authors []*struct {
		Key string `json:"key"`
	} `json:"authors,omitempty"`
	ByStatement *string `json:"by_statement,omitempty"`
	Created     struct {
		Type  string `json:"type"`
		Value string `json:"value"`
	} `json:"created"`
	Identifiers struct {
		Goodreads []string `json:"goodreads,omitempty"`
	} `json:"identifiers,omitempty"`
	ISBN10    []string `json:"isbn_10,omitempty"`
	ISBN13    []string `json:"isbn_13,omitempty"`
	Key       string   `json:"key"`
	Languages []struct {
		Key string `json:"key"`
	} `json:"languages,omitempty"`
	LastModified struct {
		Type  string `json:"type"`
		Value string `json:"value"`
	} `json:"last_modified"`
	LatestRevision     int      `json:"latest_revision"`
	LCClassifications  []string `json:"lc_classifications,omitempty"`
	LCCN               []string `json:"lccn,omitempty"`
	Location           []string `json:"location,omitempty"`
	NumberOfPages      *int     `json:"number_of_pages,omitempty"`
	Pagination         *string  `json:"pagination,omitempty"`
	PhysicalDimensions *string  `json:"physical_dimensions,omitempty"`
	PhysicalFormat     string   `json:"physical_format,omitempty"`
	PublishCountry     *string  `json:"publish_country,omitempty"`
	PublishDate        string   `json:"publish_date"`
	PublishPlaces      []string `json:"publish_places,omitempty"`
	Publishers         []string `json:"publishers,omitempty"`
	Revision           int      `json:"revision"`
	ScanRecords        []struct {
		Key string `json:"key"`
	} `json:"scan_records,omitempty"`
	SourceRecords []string `json:"source_records,omitempty"`
	Subjects      []string `json:"subjects,omitempty"`
	Subtitle      *string  `json:"subtitle,omitempty"`
	Title         string   `json:"title"`
	Type          struct {
		Key string `json:"key"`
	} `json:"type"`
	Weight *string `json:"weight,omitempty"`
	Works  []struct {
		Key string `json:"key"`
	} `json:"works,omitempty"`
}
