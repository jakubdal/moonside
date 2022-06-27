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
