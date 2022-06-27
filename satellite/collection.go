package satellite

// Collection is a dynamic, game-dependent object that can be queried through Census API.
//
// Source: https://census.daybreakgames.com/#collection
type Collection struct {
	Name        string   `json:"name"`
	Count       int64    `json:"count"`
	Hidden      bool     `json:"hidden"`
	ResolveList []string `json:"resolve_list"`
}
