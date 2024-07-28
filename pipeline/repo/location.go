package repo

import (
	"context"

	"github.com/ahmadnaufal/recommender-pipeline/model"
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

func (l *LocationPostgres) InsertPlace(ctx context.Context, place model.Place) error {
	q := `INSERT INTO locations (
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
		geom
	) VALUES (
		:id,
		:google_place_id,
		:place_name,
		:google_maps_uri,
		:address,
		:latitude,
		:longitude,
		:dish_type,
		:types,
		:rating,
		:user_rating_count,
		:reviews,
		:tags,
		:summary_review_food,
		:summary_review_place,
		:phone_number,
		:score,
		ST_SetSRID(ST_MakePoint(:longitude, :latitude), 4326)
	);`

	query, args, err := sqlx.Named(q, place)
	if err != nil {
		return err
	}

	dollarQuery := sqlx.Rebind(sqlx.DOLLAR, query)
	_, err = l.db.ExecContext(ctx, dollarQuery, args...)
	if err != nil {
		return err
	}

	return nil
}
