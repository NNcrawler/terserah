package location

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
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
	"places.plusCode",
	"places.primaryType",
	"places.primaryTypeDisplayName",
	"places.shortFormattedAddress",
	"places.location",
	"places.rating",
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

func (g *GoogleLocationAPI) GetNearby(ctx context.Context, latitude, longitude, radius float64, numOfRecommendation uint) ([]Place, error) {
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

	return placesResp.Places, nil
}
