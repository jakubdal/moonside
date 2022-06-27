package satellite

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/jakubdal/moonside/satellite/internal"
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
	HTTPDoer      internal.HTTPDoer
	BaseURL       url.URL
	ServiceID     string
	GameNamespace Namespace
}

func Default() *Satellite {
	return &Satellite{
		HTTPDoer:      http.DefaultClient,
		BaseURL:       baseURL,
		ServiceID:     DefaultServiceID,
		GameNamespace: "ps2:v2",
	}
}

func (s *Satellite) GameData(verb Verb, collection, queryString string) {
	_, err := GameData(s.HTTPDoer, s.BaseURL, s.ServiceID, verb, s.GameNamespace, collection, queryString)
	if err != nil {
		panic(err)
	}
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

func GameData(httpDoer internal.HTTPDoer, baseURL url.URL, serviceID string, verb Verb, gameNamespace Namespace, collection, queryString string) (string, error) {
	requestURL, err := GameDataURL(baseURL, serviceID, verb, gameNamespace, collection, queryString)
	if err != nil {
		return "", fmt.Errorf("GameDataURL: %w", err)
	}
	req, err := http.NewRequest(http.MethodGet, requestURL.String(), nil)
	if err != nil {
		return "", fmt.Errorf("http.NewRequest: %w", err)
	}
	resp, err := httpDoer.Do(req)
	if err != nil {
		return "", fmt.Errorf("httpDoer.Do: %w", err)
	}
	defer internal.CleanupHTTPResponse(resp)

	// WARNING: This place is likely to be a bottleneck at some point, together with the default unmarshaling.
	//
	// It will not be handled for now, because it's not a problem right now and the code is more readable this way.
	/*
		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return "", fmt.Errorf("ioutil.ReadAll: %w", err)
		}
	*/
	s, err := prettyJSON(resp)
	if err != nil {
		panic(err)
	}
	fmt.Printf("\n%s\n\n", s)

	return "", nil
}

func prettyJSON(resp *http.Response) (string, error) {
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
