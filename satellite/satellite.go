package satellite

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/tidwall/gjson"

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

func GameData[T any](httpDoer internal.HTTPDoer, requestURL url.URL) ([]T, error) {
	req, err := http.NewRequest(http.MethodGet, requestURL.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("http.NewRequest: %w", err)
	}
	resp, err := httpDoer.Do(req)
	if err != nil {
		return nil, fmt.Errorf("httpDoer.Do: %w", err)
	}
	defer internal.CleanupHTTPResponse(resp)

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected http.StatusCode in response: %v", resp.StatusCode)
	}
	// WARNING: This place is likely to be a bottleneck at some point, together with the default unmarshaling.
	//
	// It will not be handled for now, because it's not a problem right now and the code is more readable this way.
	var resultList []T
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("ioutil.ReadAll: %w", err)
	}
	// resultListKey format is probaby incorrect for the other verb?
	respJSON := gjson.ParseBytes(respBody)
	respJSON.ForEach(func(key, value gjson.Result) bool {
		if key.String() == "returned" {
			return true
		}
		err = json.Unmarshal([]byte(value.String()), &resultList)
		return false
	})
	if err != nil {
		return nil, fmt.Errorf("[respBody=%s] json.Unmarshal: %w", respBody, err)
	}

	return resultList, nil
}
