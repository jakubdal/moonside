package satellite

// Verb informs the REST interface about the of request that is being made.
//
// Source: http://census.daybreakgames.com/#verb
type Verb string

const (
	// VerbGet retrieves the game data that matches the request criteria.
	VerbGet Verb = "get"
	// VerbCount retrieve the number of game data objects that match the request criteria.
	VerbCount Verb = "count"
)

// Namespace identifies the game from which data will be queried.
//
// Source: https://census.daybreakgames.com/#namespace
type Namespace string

const (
	// NamespaceEveryQuest2 is namespace for EverQuest II.
	NamespaceEveryQuest2 Namespace = "eq2"
	// NamespacePlanetSide2PCV1 is V1 namespace for PlanetSide 2 (PC version).
	//
	// DEPRECATED: Use NamespacePlanetSide2PCV2 or NamespacePlanetSide2PC instead.
	NamespacePlanetSide2PCV1 Namespace = "ps2:v1"
	// NamespacePlanetSide2PCV2 is V2 namespace for PlanetSide 2 (PC version).
	NamespacePlanetSide2PCV2 Namespace = "ps2:v2"
	// NamespacePlanetSide2PCV1 is namespace for PlanetSide 2 (PC version).
	NamespacePlanetSide2PC Namespace = NamespacePlanetSide2PCV2

	// NamespacePlanetSide2PS4US is namespace for US PlanetSide 2 (Playstation 4).
	NamespacePlanetSide2PS4US Namespace = "ps2ps4us:v2"
	// NamespacePlanetSide2PS4EU is namespace for EU PlanetSide 2 (Playstation 4).
	NamespacePlanetSide2PS4EU Namespace = "ps2ps4eu:v2"

	// NamespaceDCUniverseOnline is namespace for DC Univese Online (PC and Playstation 3).
	NamespaceDCUniverseOnline Namespace = "dcuo:v1"
)
