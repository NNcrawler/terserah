package repo

import (
	"context"

	"github.com/ahmadnaufal/recommender-worker/model"
	"github.com/jmoiron/sqlx"
)

type LocationPostgres struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) LocationPostgres {
	return LocationPostgres{
		db: db,
	}
}

func (l *LocationPostgres) ListByClosestDistance(ctx context.Context, latitude, longitude float64, numOfRecommendations int) ([]model.Place, error) {
	var places []model.Place
	q := `
		SELECT 
			id,
			google_place_id,
			place_name,
			google_maps_uri,
			address,
			latitude,
			longitude,
			dish_type,
			types,
			rating,
			user_rating_count,
			reviews,
			tags,
			summary_review_food,
			summary_review_place,
			phone_number,
			score,
			ST_Distance(geom::geography, ST_SetSRID(ST_MakePoint($1, $2), 4326)::geography) AS distance
		FROM 
			locations
		WHERE
			ST_DWithin(
				geom,
				ST_SetSRID(ST_MakePoint($1, $2), 4326),
				$3	
			)
		ORDER BY geom <-> ST_SetSRID(ST_MakePoint($4, $5), 4326)
		LIMIT $6
	`

	err := l.db.SelectContext(ctx, &places, q, longitude, latitude, 500.0, longitude, latitude, numOfRecommendations)
	if err != nil {
		return places, err
	}

	return places, nil
}
