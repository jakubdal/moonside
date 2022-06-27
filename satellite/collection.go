package satellite

// Collection is a dynamic, game-dependent object that can be queried through Census API.
//
// Source: https://census.daybreakgames.com/#collection
// It is implemented inside Satellite, because it's a game-agnostic structure.
type Collection struct {
	Name        string   `json:"name"`
	ResolveList []string `json:"resolve_list"`

	// Hidden is sometimes bool, sometimes string. For now I will ignore it, maybe I'll handle it later.
	// Hidden      string   `json:"hidden"`
	// Count is sometimes int64, sometimes string. For now I will ignore it, maybe I'll handle it later.
	// Count       string   `json:"count"`
}
