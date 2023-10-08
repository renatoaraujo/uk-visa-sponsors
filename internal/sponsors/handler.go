package sponsors

type Handler struct {
	Scraper          Scraper
	Processor        Processor
	OrganisationList *OrganisationList
}

type Scraper interface {
	FindDataSourceURL(string, string) (string, error)
	GetContent(string) ([]byte, error)
}

type Processor interface {
	ProcessRawData([]byte, interface{}) error
}

func NewHandler(s Scraper, p Processor) Handler {
	return Handler{
		Scraper:          s,
		Processor:        p,
		OrganisationList: &OrganisationList{Organisations: []Organisation{}},
	}
}

func (h *Handler) Load(dataSource string) error {
	var err error

	if dataSource == "" {
		dataSource, err = h.Scraper.FindDataSourceURL("csv", "/government/publications/register-of-licensed-sponsors-workers")
		if err != nil {
			return err
		}
	}

	rawData, err := h.Scraper.GetContent(dataSource)
	if err != nil {
		return err
	}

	err = h.Processor.ProcessRawData(rawData, &h.OrganisationList.Organisations)
	if err != nil {
		return err
	}

	return nil
}
