package main

import (
	"fmt"
	"net/http"

	"github.com/jakubdal/moonside/satellite"
)

func main() {
	s := satellite.Default()
	verb := satellite.VerbGet
	// collection := ""
	collection := ""
	queryString := ""
	// queryString := "name.first_lower=depressiondew&c:limit=5"

	requestURL, err := s.GameDataURL(verb, collection, queryString)
	if err != nil {
		panic(err)
	}
	collections, err := satellite.GameData[satellite.Collection](http.DefaultClient, requestURL)
	if err != nil {
		panic(err)
	}

	for _, collection := range collections {
		fmt.Println(collection.Name)
	}
}
