package internal

import "fmt"

type DataFetcher interface {
	FetchData() error
}

type Handler struct {
	DataFetcher DataFetcher
}

func (h Handler) fetchData() []string {
	err := h.DataFetcher.FetchData()
	if err != nil {
		panic("fudeu")
	}
	return []string{}
}

func NewHandler(df DataFetcher) Handler {
	return Handler{DataFetcher: df}
}

func (h Handler) Find(company string) {
	h.fetchData()
	fmt.Println("trying to find the company")
}
