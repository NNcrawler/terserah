package location

type Location struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type Circle struct {
	Center Location `json:"center"`
	Radius float64  `json:"radius"`
}

type LocationRestriction struct {
	Circle Circle `json:"circle"`
}

type Request struct {
	IncludedTypes       []string            `json:"includedTypes"`
	MaxResultCount      int                 `json:"maxResultCount"`
	LocationRestriction LocationRestriction `json:"locationRestriction"`
	LanguageCode        string              `json:"languageCode"`
}
