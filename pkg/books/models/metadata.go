package models

import "strings"

// Metadata -- book metadata
type Metadata struct {
	ID             string
	Title          string
	SubTitle       string
	Author         string
	Authors        []Author
	Categories     []string
	Series         string
	Tags           string
	Ratings        string
	RatingsPercent string
	ReviewsCount   int
	RatingsSum     int
	RatingsCount   int
	URL            map[string]string
	Cover          string
	Covers         []string
	Identifiers    Identifiers
	PublishDate    string
	Publishers     []string
	PublishCountry string
	Description    string
	Subjects       string
	NumberOfPages  int
	Preview        string
	PhysicalFormat string
	IsEbook        string
	LanguageCode   string
	Legal          string

	ISBN string

	RAW interface{}
}

// Author -- Author name and id
type Author struct {
	Name string
	Key  string
	ID   string
}

// Identifiers -- book identifiers
type Identifiers struct {
	ISBN13       []string
	ISBN10       []string
	Amazon       []string
	ASIN         string
	KindleASIN   string
	Google       []string
	Gutenberg    []string
	Goodreads    []string
	Librarything []string
}

// GetAuthors -- return the author(s)
func (m Metadata) GetAuthors() string {
	authors := make([]string, len(m.Authors))
	for k, author := range m.Authors {
		authors[k] = author.Name
	}

	return strings.Join(authors, ", ")
}
