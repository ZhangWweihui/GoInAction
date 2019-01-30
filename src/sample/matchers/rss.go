package matchers

import (
	"../search"
	"encoding/xml"
	"errors"
	"fmt"
	"log"
	"net/http"
	"regexp"
)

type(
	item struct {
		XMLName		   xml.Name `xml:"item"`
		Title          string	`xml:"title"`
		Description    string	`xml:"description"`
		PubDate        string	`xml:"pubDate"`
		Link           string	`xml:"link"`
		Guid           string	`xml:"guid"`
		ContentEncoded string	`xml:"content:encoded"`
		DcCreator      string	`xml:"dc:creator"`
		GeoRssPoint    string   `xml:"georss:point"`
	}

	image struct {
		XMLName	xml.Name `xml:"image"`
		URL		string	 `xml:"url"`
		Title	string	 `xml:"title"`
		Link	string	 `xml:"link"`
	}

	channel struct {
		XMLName			xml.Name `xml:"channel"`
		Title			string	 `xml:"title"`
		Link			string	 `xml:"link"`
		Description		string	 `xml:"description"`
		Language		string	 `xml:"language"`
		Copyright		string	 `xml:"copyright"`
		Generator		string	 `xml:"generator"`
		LastBuildDate	string	 `xml:"lastBuildDate"`
		TTL             string   `xml:"ttl"`
		ManagingEditor  string   `xml:"managingEditor"`
		WebMaster       string   `xml:"webMaster"`
		Image			image	 `xml:"image"`
		Item			[]item	 `xml:"item"`
	}

	rssDocument struct {
		XMLName	xml.Name `xml:"rss"`
		channel	channel	 `xml:"channel"`
	}
)

type rssMatcher struct {}

func init(){
	var rssMatcher rssMatcher
	search.Register("rss",rssMatcher)
}

func (m rssMatcher) Search(feed *search.Feed, searchTerm string) ([]*search.Result, error) {
	var results []*search.Result
	log.Printf("Search Feed Type[%s] Site[%s] For URI[%s]\n", feed.Type, feed.Name, feed.URI)

	doc,err := m.retrieve(feed)
	if err != nil {
		return nil, err
	}

	for _,item := range doc.channel.Item {
		match, err := regexp.MatchString(searchTerm, item.Title)
		if err!=nil {
			return nil,err
		}

		if match {
			results = append(results, &search.Result{
				Field : "Field",
				Content : item.Title,
			})
		}

		match, err = regexp.MatchString(searchTerm, item.Description)
		if err!=nil {
			return nil,err
		}

		if match {
			results = append(results,&search.Result{
				Field : "Field",
				Content : item.Description,
			})
		}
	}

	return results, nil
}

func (m rssMatcher) retrieve(feed *search.Feed) (*rssDocument, error) {
	if feed.URI == "" {
		return nil,errors.New("No rss feed uri  provided")
	}

	resp, err := http.Get(feed.URI)
	if err != nil {
		return nil,err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Http response error code %d\n", resp.StatusCode)
	}

	var doc rssDocument
	err = xml.NewDecoder(resp.Body).Decode(&doc)
	return &doc, err
}
