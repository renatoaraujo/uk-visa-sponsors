package internal

import (
	"fmt"
	"log"
	"strings"
)

type DataFetcher interface {
	FetchData() ([]map[string]string, error)
}

type Handler struct {
	DataFetcher DataFetcher
}

type Organisation struct {
	name          string
	townCity      string
	country       string
	typeAndRating string
	route         string
}

var cachedOrgs []Organisation

func (h *Handler) loadData() error {
	if cachedOrgs == nil {
		data, err := h.DataFetcher.FetchData()
		if err != nil {
			return err
		}

		for _, entry := range data {
			org := Organisation{
				name:          entry["OrganisationName"],
				townCity:      entry["TownCity"],
				country:       entry["County"],
				typeAndRating: entry["TypeAndRating"],
				route:         entry["Route"],
			}
			cachedOrgs = append(cachedOrgs, org)
		}
	}

	return nil
}

func NewHandler(df DataFetcher) *Handler {
	return &Handler{DataFetcher: df}
}

func (h *Handler) searchInOrganisations(name string) []Organisation {
	var found []Organisation

	for _, org := range cachedOrgs {
		if strings.Contains(strings.ToLower(org.name), strings.ToLower(name)) {
			found = append(found, org)
		}
	}

	if len(found) > 1 {
		fmt.Println("multiple organisations with this name")
	}

	return found
}

func (h *Handler) Find(company string) {
	if err := h.loadData(); err != nil {
		log.Fatalf("failed to load data; %w", err)
	}

	orgs := h.searchInOrganisations(company)

	for _, org := range orgs {
		fmt.Println(fmt.Sprintf("company %s found, and it is ranked as %s and can provide the %s", org.name, org.typeAndRating, org.route))
	}
}
