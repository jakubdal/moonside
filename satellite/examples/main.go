package main

import (
	"github.com/jakubdal/moonside/satellite"
)

func main() {
	s := satellite.Default()
	verb := satellite.VerbGet
	// collection := ""
	collection := ""
	queryString := ""
	// queryString := "name.first_lower=depressiondew&c:limit=5"

	s.GameData(verb, collection, queryString)
}
