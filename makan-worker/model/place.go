package model

import "github.com/lib/pq"

type Place struct {
	ID                 string         `db:"id"`
	GooglePlaceID      string         `db:"google_place_id"`
	PlaceName          string         `db:"place_name"`
	GoogleMapsURI      string         `db:"google_maps_uri"`
	Address            string         `db:"address"`
	Latitude           float64        `db:"latitude"`
	Longitude          float64        `db:"longitude"`
	Types              pq.StringArray `db:"types"`
	PrimaryType        string         `db:"primary_type"`
	PhoneNumber        string         `db:"phone_number"`
	Rating             float64        `db:"rating"`
	UserRatingCount    int            `db:"user_rating_count"`
	PriceLevel         string         `db:"price_level"`
	Reviews            pq.StringArray `db:"reviews"`
	DishType           pq.StringArray `db:"dish_type"`
	SummaryReviewFood  string         `db:"summary_review_food"`
	SummaryReviewPlace string         `db:"summary_review_place"`
	Score              int            `db:"score"`
	Distance           float64        `db:"distance"`

	Tags       pq.StringArray `db:"tags"`
	TotalScore float64
}
