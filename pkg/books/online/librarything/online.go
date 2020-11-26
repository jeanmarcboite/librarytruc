package librarything

import (
	"encoding/xml"
	"fmt"
	"strings"

	"github.com/rs/zerolog/log"

	xml2json "github.com/basgys/goxml2json"
	"github.com/jeanmarcboite/truc/pkg/books/models"
	"github.com/jeanmarcboite/truc/pkg/books/online/net"
)

// LookUpISBN -- lookup a Book on goodreads, with isbn
func LookUpISBN(isbn string) (models.Metadata, error) {
	return get(isbn, net.Koanf.String("librarything.url.isbn"))
}

func getMeta(response Response) (models.Metadata, error) {
	author := models.Author{
		Name: response.Ltml.Item.Author.Text,
		ID:   response.Ltml.Item.Author.ID,
		Key:  response.Ltml.Item.Author.Authorcode,
	}
	return models.Metadata{
		ID:      response.Ltml.Item.ID,
		Title:   response.Ltml.Item.Title,
		Authors: []models.Author{author},
		RAW:     response,
	}, nil
}

func get(what string, where string) (models.Metadata, error) {
	url := fmt.Sprintf(where, what)

	resp, err := net.HTTPGetWithKey(url,
		net.Koanf.String("librarything.keyname"),
		net.Koanf.String("librarything.key"))
	if err != nil {
		log.Error().Str("url", url).Msg(err.Error())
		return models.Metadata{}, err
	}

	var response Response

	/* Book could be: <error>Page not found</error> */
	xml.Unmarshal(resp, &response)

	if response.XMLName.Local == "response" {
		if response.Stat == "fail" {
			xml := strings.NewReader(string(resp))
			json, _ := xml2json.Convert(xml)

			err := fmt.Errorf("%v", json)
			log.Error().Str(url, "url").Msg(err.Error())
			return models.Metadata{}, err
		}

		return getMeta(response)
	}

	return models.Metadata{}, fmt.Errorf("LibraryThing for '%v': %v", what, response)
}
