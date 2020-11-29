package google

import (
	"encoding/json"
	"fmt"

	"github.com/rs/zerolog/log"

	"github.com/jeanmarcboite/librarytruc/pkg/books/models"
	"github.com/jeanmarcboite/librarytruc/pkg/books/online/net"
)

// LookUpISBN -- lookup a work on google, with isbn
func LookUpISBN(isbn string) (models.Metadata, error) {
	return get(isbn, net.Koanf.String("google.url.isbn"))
}

func get(isbn string, where string) (models.Metadata, error) {
	url := fmt.Sprintf(where, isbn)
	resp, err := net.HTTPGet(url)
	if err != nil {
		log.Error().Str(url, "url").Msg(err.Error())
		return models.Metadata{}, err
	}

	var response Response
	json.Unmarshal([]byte(resp), &response)

	meta, err := getMeta(response)
	meta.ISBN = isbn
	return meta, err
}

func getMeta(response Response) (models.Metadata, error) {
	if response.TotalItems < 1 {
		return models.Metadata{}, nil
	}

	item := response.Items[0]

	authors := make([]models.Author, 0)

	for _, author := range item.VolumeInfo.Authors {
		authors = append(authors, models.Author{
			Name: author,
		})
	}

	is, _ := json.MarshalIndent(item, "", "  ")
	fmt.Printf(string(is))

	return models.Metadata{
		ID:            item.ID,
		Title:         item.VolumeInfo.Title,
		Authors:       authors,
		Publishers:    []string{item.VolumeInfo.Publisher},
		Description:   item.VolumeInfo.Description,
		NumberOfPages: item.VolumeInfo.PageCount,
		Categories:    item.VolumeInfo.Categories,
		Cover:         item.VolumeInfo.ImageLinks.Thumbnail,
		RAW:           item,
	}, nil
}
