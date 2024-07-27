package location

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/ahmadnaufal/recommender-worker/model"
)

type PlaceResponse struct {
	Places []Place `json:"places"`
}

var expectedResponseFields = []string{
	"places.displayName",
	"places.formattedAddress",
	"places.googleMapsUri",
	"places.id",
	"places.location",
	"places.name",
	"places.photos",
	"places.types",
	"places.primaryType",
	"places.primaryTypeDisplayName",
	"places.internationalPhoneNumber",
	"places.location",
	"places.rating",
	"places.userRatingCount",
	"places.currentOpeningHours",
	"places.servesBreakfast",
	"places.servesLunch",
	"places.servesDinner",
	"places.servesVegetarianFood",
	"places.servesBrunch",
	"places.servesCoffee",
	"places.priceLevel",
	"places.reviews",
}

var includedFnbTypes = []string{"restaurant", "cafe", "bakery", "bar"}

type GoogleLocationAPI struct {
	client *http.Client
	host   string
	apiKey string
}

func New(host, apiKey string) GoogleLocationAPI {
	return GoogleLocationAPI{
		client: &http.Client{},
		host:   host,
		apiKey: apiKey,
	}
}

func (g *GoogleLocationAPI) GetNearby(ctx context.Context, latitude, longitude, radius float64, numOfRecommendation uint) ([]model.Place, error) {
	requestURL := fmt.Sprintf("%s/v1/places:searchNearby", g.host)
	requestNearby := Request{
		IncludedTypes:  includedFnbTypes,
		MaxResultCount: int(numOfRecommendation),
		LocationRestriction: LocationRestriction{
			Circle: Circle{
				Center: Location{
					Latitude:  latitude,
					Longitude: longitude,
				},
				Radius: radius,
			},
		},
		LanguageCode: "id",
	}
	requestBody, _ := json.Marshal(requestNearby)
	req, _ := http.NewRequest(http.MethodPost, requestURL, bytes.NewBuffer(requestBody))
	// provide headers
	req.Header.Add("X-Goog-Api-Key", g.apiKey)
	req.Header.Add("X-Goog-FieldMask", strings.Join(expectedResponseFields, ","))

	resp, err := g.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var placesResp PlaceResponse
	err = json.NewDecoder(resp.Body).Decode(&placesResp)
	if err != nil {
		return nil, err
	}

	modelPlaces := []model.Place{}
	for _, lpc := range placesResp.Places {
		modelPlaces = append(modelPlaces, placeResponseToPlaceModel(lpc))
	}

	return modelPlaces, nil
}

func placeResponseToPlaceModel(place Place) model.Place {
	var res model.Place

	res.GooglePlaceID = place.ID
	res.PlaceName = place.DisplayName.Text
	res.GoogleMapsURI = place.GoogleMapsUri
	res.Address = place.FormattedAddress
	res.Latitude = place.Location.Latitude
	res.Longitude = place.Location.Longitude
	res.Types = place.Types
	res.PrimaryType = place.PrimaryType
	res.PhoneNumber = place.InternationalPhoneNumber
	res.Rating = place.Rating
	res.UserRatingCount = place.UserRatingCount
	res.PriceLevel = place.PriceLevel

	res.Tags = []string{}
	if place.ServesCoffee {
		res.Tags = append(res.Tags, "coffee")
	}
	if place.ServesBreakfast {
		res.Tags = append(res.Tags, "breakfast")
	}
	if place.ServesLunch {
		res.Tags = append(res.Tags, "lunch")
	}
	if place.ServesDinner {
		res.Tags = append(res.Tags, "dinner")
	}
	if place.ServesBrunch {
		res.Tags = append(res.Tags, "brunch")
	}
	if place.ServesVegetarianFood {
		res.Tags = append(res.Tags, "vegetarian")
	}
	if place.ServesDessert {
		res.Tags = append(res.Tags, "dessert")
	}

	for _, r := range place.Reviews {
		res.Reviews = append(res.Reviews, r.Text.Text)
	}

	return res
}
