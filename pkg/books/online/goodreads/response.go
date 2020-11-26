package goodreads

import (
	"encoding/xml"
)

// Response -- GoodreadsResponse
type Response struct {
	XMLName xml.Name `xml:"GoodreadsResponse"`
	Books   []Book   `xml:"book"`
}

// Book - one book i the response
type Book struct {
	ID               string `xml:"id"`
	Title            string `xml:"title"`
	ISBN             string `xml:"isbn"`
	ISBN13           string `xml:"isbn13"`
	ASIN             string `xml:"asin"`
	KindleASIN       string `xml:"kindle_asin"`
	MarketplaceID    string `xml:"marketplace_id"`
	CountryCode      string `xml:"county_code"`
	ImageURL         string `xml:"image_url"`
	SmallImageURL    string `xml:"small_image_url"`
	PublicationYear  string `xml:"publication_year"`
	PublicationMonth string `xml:"publication_month"`
	PublicationDay   string `xml:"publication_day"`
	Publisher        string `xml:"publisher"`
	LanguageCode     string `xml:"language_code"`
	IsEbook          string `xml:"is_ebook"`
	Description      string `xml:"description"`
	Work             Work   `xml:"work"`
}

// Work -- info about the original work
type Work struct {
	ID                             int    `xml:"id"`
	BooksCount                     int    `xml:"books_count"`
	BestBookID                     int    `xml:"best_book_id"`
	ReviewsCount                   int    `xml:"reviews_count"`
	RatingsSum                     int    `xml:"ratings_sum"`
	RatingsCount                   int    `xml:"ratings_count"`
	TextReviewsCount               int    `xml:"text_reviews_count"`
	OriginalPublicationYear        int    `xml:"original_publication_year"`
	OriginalPublicationMonth       int    `xml:"original_publication_month"`
	OriginalPublicationDay         int    `xml:"original_publication_day"`
	OriginalTitle                  string `xml:"original_title"`
	OriginalLanguageID             int    `xml:"original_language_id"`
	MediaType                      string `xml:"media_type"`
	RatingDist                     string `xml:"rating_dist"`
	DescUserID                     int    `xml:"desc_user_id"`
	DefaultChapteringBookID        int    `xml:"default_chaptering_book_id"`
	DefaultDescriptionLanguageCode string `xml:"default_description_language_code"`
	WorkURI                        string `xml:"work_uri"`
}
