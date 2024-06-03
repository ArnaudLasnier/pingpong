package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

const Endpoint = "https://api.sampleapis.com/beers/ale"

type Beer struct {
	Price  string `json:"price"`
	Name   string `json:"name"`
	Rating struct {
		Average float64 `json:"average"`
		Reviews int64   `json:"reviews"`
	}
	Image string `json:"image"`
	ID    int    `json:"id"`
}

func main() {
	response, err := http.Get(Endpoint)
	if err != nil {
		panic(err)
	}
	body, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	apiBeers := make([]*Beer, 0)
	err = json.Unmarshal(body, &apiBeers)
	if err != nil {
		panic(err)
	}
	csvBeers := make([][]string, 0)
	for _, beer := range apiBeers {
		csvBeers = append(csvBeers, []string{
			fmt.Sprint(beer.ID),
			beer.Name,
			beer.Price,
		})
	}
	file, err := os.Create("beers.csv")
	if err != nil {
		panic(err)
	}
	csvWriter := csv.NewWriter(file)
	err = csvWriter.WriteAll(csvBeers)
	if err != nil {
		panic(err)
	}
}
