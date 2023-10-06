package pkg

type Scraper struct {
	url string
}

func NewScraper(url string) *Scraper {
	return &Scraper{
		url: url,
	}
}

func (s *Scraper) FetchData() {

}
