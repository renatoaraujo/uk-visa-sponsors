package internal

import "fmt"

type DataFetcher interface {
	FetchData()
}

type Handler struct {
	DataFetcher DataFetcher
}

func (h Handler) fetchData() []string {
	return []string{}
}

func NewHandler(df DataFetcher) Handler {
	return Handler{DataFetcher: df}
}

func (h Handler) Find(company string) {
	fmt.Println("trying to find the company")
}
