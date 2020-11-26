package librarything

import (
	"encoding/xml"
)

// Response was generated 2019-11-27 06:55:35 by box on redkeep.
type Response struct {
	XMLName xml.Name `xml:"response"`
	Text    string   `xml:",chardata"`
	Stat    string   `xml:"stat,attr"`
	Ltml    struct {
		Text    string `xml:",chardata"`
		Xmlns   string `xml:"xmlns,attr"`
		Version string `xml:"version,attr"`
		Item    struct {
			Text   string `xml:",chardata"`
			ID     string `xml:"id,attr"`
			Type   string `xml:"type,attr"`
			Author struct {
				Text       string `xml:",chardata"` // Philip Kerr
				ID         string `xml:"id,attr"`
				Authorcode string `xml:"authorcode,attr"`
			} `xml:"author"`
			Title           string `xml:"title"`  // March Violets
			Rating          string `xml:"rating"` // 7.2
			URL             string `xml:"url"`    // http://www.librarything.c...
			Commonknowledge struct {
				Text      string `xml:",chardata"`
				FieldList struct {
					Text  string `xml:",chardata"`
					Field []struct {
						Text        string `xml:",chardata"`
						Type        string `xml:"type,attr"`
						Name        string `xml:"name,attr"`
						DisplayName string `xml:"displayName,attr"`
						VersionList struct {
							Text    string `xml:",chardata"`
							Version struct {
								Text     string `xml:",chardata"`
								ID       string `xml:"id,attr"`
								Archived string `xml:"archived,attr"`
								Lang     string `xml:"lang,attr"`
								Date     struct {
									Text      string `xml:",chardata"` // Fri, 02 Mar 2018 17:38:43...
									Timestamp string `xml:"timestamp,attr"`
								} `xml:"date"`
								Person struct {
									Text string `xml:",chardata"`
									ID   string `xml:"id,attr"`
									Name string `xml:"name"` // elkiedee, elkiedee, elkie...
									URL  string `xml:"url"`  // http://www.librarything.c...
								} `xml:"person"`
								FactList struct {
									Text string   `xml:",chardata"`
									Fact []string `xml:"fact"` // March Violets, ![CDATA[ <...
								} `xml:"factList"`
							} `xml:"version"`
						} `xml:"versionList"`
					} `xml:"field"`
				} `xml:"fieldList"`
			} `xml:"commonknowledge"`
		} `xml:"item"`
		Legal string `xml:"legal"` // By using this data you ag...
	} `xml:"ltml"`
}
