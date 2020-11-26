package openlibrary

// Response -- struct returned by http://openlibrary.org/search.json?q=the+lord+of+the+rings
type Response struct {
	NumFound int   `json:"numFound"`
	Docs     []Doc `json:"docs"`
}

// Doc -- one item of Response
type Doc struct {
	TitleSuggest     string   `json:"title_suggest"`
	AuthorKey        []string `json:"author_key"`
	AuthorName       []string `json:"author_name"`
	Contributor      []string `json:"contributor"`
	CoverI           string   `json:"cover_i"`
	EbookCount       int      `json:"ebook_count"`
	EditionKey       []string `json:"edition_key"`
	FirstPublishYear int      `json:"first_publish_year"`
	HasFulltext      bool     `json:"has_full_text"`
	ISBN             []string `json:"isbn"`
	Key              string   `json:"key"`
	Language         []string `json:"language"`
	OCLC             []string `json:"oclc"`
	Place            []string `json:"place"`
	PublishDate      []string `json:"publish_date"`
	PublishYear      []int    `json:"publish_year"`
	Seed             []string `json:"seed"`
	Subject          []string `json:"subject"`
	Text             []string `json:"text"`
	Title            string   `json:"title"`
	Type             string   `json:"type"`
}

// BookResponse -- response from openlibrary
type BookResponse struct {
	Data Book `json:"data"`
}

// Book -- struct return for a given work (e.g. isbn)
type Book struct {
	InfoURL      string `json:"info_url"`
	BibKey       string `json:"bib_key"`
	PreviewURL   string `json:"preview_url"`
	ThumbnailURL string `json:"thumbnail_url"`
	Preview      string `json:"preview"`
	Cover        string
	Details      struct {
		NumberOfPages   int `json:"number_of_pages"`
		TableOfContents []struct {
			Title string `json:"title"`
			Type  struct {
				Key string `json:"key"`
			}
			Level int `json:"level"`
		}
		Weight            string   `json:"weight"`
		Covers            []string `json:"covers"`
		LCClassifications []string `json:"lc_classifications"`
		LatestRevision    int      `json:"latest_revision"`
		SourceRecords     []string `json:"source_records"`
		Title             string   `json:"title"`
		Languages         []struct {
			Key string `json:"key"`
		} `json:"languages"`
		Subjects       []string `json:"subjects"`
		PublishCountry string   `json:"publish_country"`
		ByStatement    string   `json:"by_statement"`
		OCLCNumbers    []string `json:"oclc_numbers"`
		Type           struct {
			Key string `json:"key"`
		} `json:"type"`
		PhysicalDimensions string   `json:"physical_dimensions"`
		Revision           string   `json:"revision"`
		Publishers         []string `json:"publishers"`
		Description        string   `json:"description"`
		PhysicalFormat     string   `json:"physical_format"`
		LastModified       struct {
			Type  string `json:"type"`
			Value string `json:"value"`
		} `json:"last_modified"`
		Key     string `json:"key"`
		Authors []struct {
			Name string `json:"name"`
			Key  string `json:"key"`
		} `json:"authors"`
		PublishPlaces   string `json:"publish_places"`
		Pagination      string `json:"pagination"`
		Classifications string `json:"classifications"`
		Created         struct {
			Type  string `json:"type"`
			Value string `json:"value"`
		} `json:"created"`
		LCCN        []string `json:"lccn"`
		Notes       string   `json:"notes"`
		Identifiers struct {
			Amazon       []string `json:"amazon"`
			Google       []string `json:"google"`
			Gutenberg    []string `json:"project_gutenberg"`
			Goodreads    []string `json:"goodreads"`
			Librarything []string `json:"librarything"`
		} `json:"identifiers"`
		ISBN13            []string `json:"isbn_13"`
		DeweyDecimalClass []string `json:"dewey_decimal_class"`
		ISBN10            []string `json:"isbn_10"`
		PublishDate       string   `json:"publish_date"`
		Works             []struct {
			Key string `json:"key"`
		} `json:"works"`
	} `json:"details"`
}

// AuthorsName -- return the names of the author(s)
func (b Book) AuthorsName() []string {
	names := make([]string, len(b.Details.Authors))
	for i, v := range b.Details.Authors {
		names[i] = v.Name
	}
	return names
}

/*
 'https://openlibrary.org/api/books?bibkeys=ISBN:9780980200447&jscmd=details&format=json'
 'https://openlibrary.org/api/books?bibkeys=ISBN:0201558025&format=json'
 **/
