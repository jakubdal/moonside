package satellite

import (
	"testing"

	"github.com/jakubdal/moonside/satellite/mt"
)

// Source: https://census.daybreakgames.com/#url-pattern
func Test_GameDataURL(t *testing.T) {
	gameDataURL, err := GameDataURL(baseURL, DefaultServiceID, "get", "ps2:v2", "character", "character_id=5428018587875812257&c:show=name")
	mt.FailOnError(t, "", err)
	expectedDataURL := "https://census.daybreakgames.com/s:example/get/ps2:v2/character/?character_id=5428018587875812257&c:show=name"
	mt.ErrorIfNotEqual(t, expectedDataURL, gameDataURL.String())
}

// Source: https://census.daybreakgames.com/#url-pattern
func Test_GameImageURL(t *testing.T) {
	imageURL, err := GameImageURL(baseURL, "ps2", "static", "5391")
	mt.FailOnError(t, "", err)
	expectedImageURL := "https://census.daybreakgames.com/files/ps2/images/static/5391.png"
	mt.ErrorIfNotEqual(t, expectedImageURL, imageURL.String())
}
