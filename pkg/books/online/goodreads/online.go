package goodreads

import (
	"encoding/xml"
	"fmt"
	"strings"

	"github.com/jeanmarcboite/librarytruc/pkg/books/models"
	"github.com/jeanmarcboite/librarytruc/pkg/books/online/net"
	"github.com/rs/zerolog/log"
)

// LookUpISBN -- lookup a work on goodreads, with isbn
func LookUpISBN(isbn string) (models.Metadata, error) {
	return get(isbn, net.Koanf.String("goodreads.url.isbn"))
}

// SearchTitle -- search for a work with a title
func SearchTitle(title string) (models.Metadata, error) {
	return get(strings.Join(strings.Fields(title), "+"),
		net.Koanf.String("goodreads.url.title"))
}

func get(what string, where string) (models.Metadata, error) {
	url := fmt.Sprintf(where, what)

	response, err := net.HTTPGetWithKey(url,
		net.Koanf.String("goodreads.keyname"),
		net.Koanf.String("goodreads.key"))
	if err != nil {
		log.Error().Str("url", url).Msg(err.Error())
		return models.Metadata{}, err
	}

	var goodreads Response

	/* response could be: <error>Page not found</error> */
	xml.Unmarshal(response, &goodreads)

	if goodreads.XMLName.Local == "GoodreadsResponse" {
		return getMeta(goodreads.Books[0])
	}

	return models.Metadata{}, fmt.Errorf("Nothing found on goodreads for '%v'", what)
}

func getMeta(goodreads Book) (models.Metadata, error) {
	meta := models.Metadata{
		ID:      goodreads.ID,
		Title:   goodreads.Title,
		Authors: []models.Author{},
		Identifiers: models.Identifiers{
			ISBN10:     []string{goodreads.ISBN},
			ISBN13:     []string{goodreads.ISBN13},
			ASIN:       goodreads.ASIN,
			KindleASIN: goodreads.KindleASIN,
		},
		PublishCountry: goodreads.CountryCode,
		Publishers:     []string{goodreads.Publisher},
		Description:    goodreads.Description,
		Cover:          goodreads.ImageURL,
		IsEbook:        goodreads.IsEbook,
		ReviewsCount:   goodreads.Work.ReviewsCount,
		RatingsSum:     goodreads.Work.RatingsSum,
		RatingsCount:   goodreads.Work.RatingsCount,
		Ratings:        goodreads.Work.RatingDist,

		RAW: goodreads,
	}

	return meta, nil
}
