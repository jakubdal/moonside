package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/jakubdal/moonside/satellite"
)

func main() {
	s := satellite.Default()
	verb := satellite.VerbGet
	collection := ""
	queryString := ""

	gameDataURL, err := s.GameDataURL(verb, collection, queryString)
	if err != nil {
		panic(err)
	}
	resp, err := http.Get(gameDataURL.String())
	if err != nil {
		panic(err)
	}
	prettyResp, err := prettyJSON(resp)
	if err != nil {
		panic(err)
	}
	fmt.Println(prettyResp)
}

func prettyJSON(resp *http.Response) (string, error) {
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("ioutil.ReadAll: %w", err)
	}
	var m map[string]interface{}
	err = json.Unmarshal(respBody, &m)
	if err != nil {
		return "", fmt.Errorf("[respBody=%s] json.Unmarshal: %w", respBody, err)
	}
	prettyB, err := json.MarshalIndent(m, "", "\t")
	if err != nil {
		return "", fmt.Errorf("[respBody=%s] json.MarshalIndent: %w", respBody, err)
	}
	return string(prettyB), nil
}
