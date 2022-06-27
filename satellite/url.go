package satellite

import (
	"fmt"
	"net/url"

	"github.com/jakubdal/moonside/satellite/mt"
)

const (
	BaseURL          = "https://census.daybreakgames.com"
	DefaultServiceID = "s:example"
)

var baseURL = *mt.Must(url.Parse(BaseURL))

func GameDataURL(baseURL url.URL, serviceID string, verb Verb, gameNamespace Namespace, collection, queryString string) (url.URL, error) {
	baseURL.Path = fmt.Sprintf("%s/%s/%s/%s/", serviceID, verb, gameNamespace, collection)
	baseURL.RawQuery = queryString
	return baseURL, nil
}

func GameImageURL(baseURL url.URL, gameNamespace Namespace, imageType, imageID string) (url.URL, error) {
	baseURL.Path = fmt.Sprintf("files/%s/images/%s/%s.png", gameNamespace, imageType, imageID)
	return baseURL, nil
}

// Pathfinder groups API interfaces for a game namespace.
//
// The goal of this structure is to make writing function calls faster.
type Pathfinder struct {
	BaseURL       url.URL
	ServiceID     string
	GameNamespace Namespace
}

func DefaultPathfinder() *Pathfinder {
	return &Pathfinder{
		BaseURL:       baseURL,
		ServiceID:     DefaultServiceID,
		GameNamespace: "ps2:v2",
	}
}

func (s *Pathfinder) GameDataURL(verb Verb, collection, queryString string) (url.URL, error) {
	return GameDataURL(s.BaseURL, s.ServiceID, verb, s.GameNamespace, collection, queryString)
}

func (s *Pathfinder) GameImageURL(imageType, imageID string) (url.URL, error) {
	return GameImageURL(s.BaseURL, s.GameNamespace, imageType, imageID)
}
