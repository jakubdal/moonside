package satellite

import (
	"fmt"
	"net/url"
)

const (
	BaseURL          = "https://census.daybreakgames.com"
	DefaultServiceID = "s:example"
)

// TODO(jakubdal): Generic "must" function potential?
var baseURL = func() url.URL {
	baseURL, err := url.Parse(BaseURL)
	if err != nil {
		panic(err)
	}
	return *baseURL
}()

// Satellite groups raw contact interfaces for a game namespace.
type Satellite struct {
	BaseURL       url.URL
	ServiceID     string
	GameNamespace Namespace
}

func Default() *Satellite {
	return &Satellite{
		BaseURL:       baseURL,
		ServiceID:     DefaultServiceID,
		GameNamespace: "ps2:v2",
	}
}

// Collections returns a list of Collection available for GameNamespace set in constructor.
//
// It is a shortcut to calling GameData TODO: FINISH COMMENT
func (s *Satellite) Collections() ([]Collection, error) {
	type CollectionResponse struct {
		Collections []Collection `json:"datatype_list`
	}
	requestURL, err := s.GameDataURL(VerbGet, string(s.GameNamespace), "")
	_, _ = requestURL, err
	return nil, nil
}

func (s *Satellite) GameData() {

}

func (s *Satellite) GameDataURL(verb Verb, collection, queryString string) (url.URL, error) {
	return GameDataURL(s.BaseURL, s.ServiceID, verb, s.GameNamespace, collection, queryString)
}

func (s *Satellite) GameImageURL(imageType, imageID string) (url.URL, error) {
	return GameImageURL(s.BaseURL, s.GameNamespace, imageType, imageID)
}

func GameDataURL(baseURL url.URL, serviceID string, verb Verb, gameNamespace Namespace, collection, queryString string) (url.URL, error) {
	baseURL.Path = fmt.Sprintf("%s/%s/%s/%s/", serviceID, verb, gameNamespace, collection)
	baseURL.RawQuery = queryString
	return baseURL, nil
}

func GameImageURL(baseURL url.URL, gameNamespace Namespace, imageType, imageID string) (url.URL, error) {
	baseURL.Path = fmt.Sprintf("files/%s/images/%s/%s.png", gameNamespace, imageType, imageID)
	return baseURL, nil
}
