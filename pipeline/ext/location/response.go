package location

import "fmt"

type LocalizedText struct {
	LanguageCode string `json:"languageCode"`
	Text         string `json:"text"`
}

type AddressComponent struct {
	LongName  string   `json:"longName"`
	ShortName string   `json:"shortName"`
	Types     []string `json:"types"`
}

type LatLng struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type Review struct {
	AuthorName              string        `json:"authorName"`
	Rating                  int           `json:"rating"`
	Text                    LocalizedText `json:"originalText"`
	Time                    int64         `json:"time"`
	RelativeTimeDescription string        `json:"relativeTimeDescription"`
}

type OpeningHours struct {
	OpenNow     bool     `json:"openNow"`
	WeekdayText []string `json:"weekdayText"`
}

type Photo struct {
	Height           int      `json:"height"`
	HtmlAttributions []string `json:"htmlAttributions"`
	PhotoReference   string   `json:"photoReference"`
	Width            int      `json:"width"`
}

type PaymentOptions struct {
	AcceptedMethods []string `json:"acceptedMethods"`
}

type GenerativeSummary struct {
	Summary string `json:"summary"`
}

type AreaSummary struct {
	Summary string `json:"summary"`
}

type AccessibilityOptions struct {
	WheelchairAccessible bool `json:"wheelchairAccessible"`
}

type Place struct {
	Name                     string               `json:"name"`
	ID                       string               `json:"id"`
	DisplayName              LocalizedText        `json:"displayName"`
	Types                    []string             `json:"types"`
	PrimaryType              string               `json:"primaryType"`
	InternationalPhoneNumber string               `json:"internationalPhoneNumber"`
	FormattedAddress         string               `json:"formattedAddress"`
	AddressComponents        []AddressComponent   `json:"addressComponents"`
	Location                 LatLng               `json:"location"`
	Rating                   float64              `json:"rating"`
	GoogleMapsUri            string               `json:"googleMapsUri"`
	RegularOpeningHours      OpeningHours         `json:"regularOpeningHours"`
	Photos                   []Photo              `json:"photos"`
	PriceLevel               string               `json:"priceLevel"`
	Reviews                  []Review             `json:"reviews"`
	CurrentOpeningHours      OpeningHours         `json:"currentOpeningHours"`
	EditorialSummary         LocalizedText        `json:"editorialSummary"`
	GenerativeSummary        GenerativeSummary    `json:"generativeSummary"`
	AreaSummary              AreaSummary          `json:"areaSummary"`
	UserRatingCount          int                  `json:"userRatingCount"`
	Takeout                  bool                 `json:"takeout"`
	Delivery                 bool                 `json:"delivery"`
	DineIn                   bool                 `json:"dineIn"`
	CurbsidePickup           bool                 `json:"curbsidePickup"`
	Reservable               bool                 `json:"reservable"`
	ServesBreakfast          bool                 `json:"servesBreakfast"`
	ServesLunch              bool                 `json:"servesLunch"`
	ServesDinner             bool                 `json:"servesDinner"`
	ServesBeer               bool                 `json:"servesBeer"`
	ServesWine               bool                 `json:"servesWine"`
	ServesBrunch             bool                 `json:"servesBrunch"`
	ServesVegetarianFood     bool                 `json:"servesVegetarianFood"`
	OutdoorSeating           bool                 `json:"outdoorSeating"`
	LiveMusic                bool                 `json:"liveMusic"`
	MenuForChildren          bool                 `json:"menuForChildren"`
	ServesCocktails          bool                 `json:"servesCocktails"`
	ServesDessert            bool                 `json:"servesDessert"`
	ServesCoffee             bool                 `json:"servesCoffee"`
	GoodForChildren          bool                 `json:"goodForChildren"`
	AllowsDogs               bool                 `json:"allowsDogs"`
	Restroom                 bool                 `json:"restroom"`
	GoodForGroups            bool                 `json:"goodForGroups"`
	GoodForWatchingSports    bool                 `json:"goodForWatchingSports"`
	AccessibilityOptions     AccessibilityOptions `json:"accessibilityOptions"`
	// Reviews                      []Review             `json:"reviews"`
}

func (p Place) String() string {
	baseStr := fmt.Sprintf("%s\n%s\n%s\n%s\n%f %f\n",
		p.DisplayName.Text, p.FormattedAddress, p.Name, p.PrimaryType,
		p.Location.Latitude, p.Location.Longitude)

	return baseStr
}
