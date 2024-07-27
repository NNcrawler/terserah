package model

type Place struct {
	ID              string
	PlaceName       string
	GoogleMapsURI   string
	Address         string
	Latitude        float64
	Longitude       float64
	Types           []string
	PrimaryType     string
	PhoneNumber     string
	IsOpen          bool
	Rating          float64
	UserRatingCount int
	PriceLevel      string
	Reviews         []string
	DishType        []string

	ReviewsSummary ReviewSummary

	Tags []string
}

type ReviewSummary struct {
	Food  string
	Place string
}
