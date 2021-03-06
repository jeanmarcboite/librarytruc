package online

import (
	"github.com/jeanmarcboite/librarytruc/pkg/books/models"
	"github.com/jeanmarcboite/librarytruc/pkg/books/online/goodreads"
	"github.com/jeanmarcboite/librarytruc/pkg/books/online/google"
	"github.com/jeanmarcboite/librarytruc/pkg/books/online/openlibrary"
)

// LookUpISBN -- lookup a work on goodreads and openlibrary, with isbn
func LookUpISBN(isbn string) (map[string]models.Metadata, error) {
	metadata := make(map[string]models.Metadata)

	/** LibraryThing returns "APIs Temporarily disabled."
	l, err := librarything.LookUpISBN(isbn)

	if err != nil {
		return metadata, err
	}
	metadata["librarything"] = l
	*/
	o, err := openlibrary.LookUpISBN(isbn)
	if err == nil {
		metadata["openlibrary"] = o
	}

	g, err := goodreads.LookUpISBN(isbn)
	if err == nil {
		metadata["goodreads"] = g
	}

	goog, err := google.LookUpISBN(isbn)
	if err == nil {
		metadata["google"] = goog
	}

	return metadata, nil
}

// SearchTitle --
func LookUpTitle(title string) ([]map[string]models.Metadata, error) {
	docs, err := openlibrary.LookUpTitle(title)
	if err != nil {
		return nil, err
	}

	books := make([]map[string]models.Metadata, len(docs))
	for k, doc := range docs {
		metadata := make(map[string]models.Metadata)
		if err == nil {
			metadata["openlibrary"] = doc
			if doc.ISBN != "" {
				g, err := goodreads.LookUpISBN(doc.ISBN)
				if err == nil {
					metadata["goodreads"] = g
				}

			}
			books[k] = metadata
		}
	}

	return books, err
}
