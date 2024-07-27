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

type PlusCode struct {
	CompoundCode string `json:"compoundCode"`
	GlobalCode   string `json:"globalCode"`
}

type LatLng struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type Viewport struct {
	Northeast LatLng `json:"northeast"`
	Southwest LatLng `json:"southwest"`
}

type Review struct {
	AuthorName              string `json:"authorName"`
	Rating                  int    `json:"rating"`
	Text                    string `json:"text"`
	Time                    int64  `json:"time"`
	RelativeTimeDescription string `json:"relativeTimeDescription"`
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

type Attribution struct {
	Source string `json:"source"`
	Url    string `json:"url"`
}

type PaymentOptions struct {
	AcceptedMethods []string `json:"acceptedMethods"`
}

type ParkingOptions struct {
	ParkingType []string `json:"parkingType"`
}

type SubDestination struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type FuelOptions struct {
	FuelTypes []string `json:"fuelTypes"`
}

type EVChargeOptions struct {
	ChargingStations []string `json:"chargingStations"`
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
	Name                         string               `json:"name"`
	ID                           string               `json:"id"`
	DisplayName                  LocalizedText        `json:"displayName"`
	Types                        []string             `json:"types"`
	PrimaryType                  string               `json:"primaryType"`
	PrimaryTypeDisplayName       LocalizedText        `json:"primaryTypeDisplayName"`
	NationalPhoneNumber          string               `json:"nationalPhoneNumber"`
	InternationalPhoneNumber     string               `json:"internationalPhoneNumber"`
	FormattedAddress             string               `json:"formattedAddress"`
	ShortFormattedAddress        string               `json:"shortFormattedAddress"`
	AddressComponents            []AddressComponent   `json:"addressComponents"`
	PlusCode                     PlusCode             `json:"plusCode"`
	Location                     LatLng               `json:"location"`
	Viewport                     Viewport             `json:"viewport"`
	Rating                       float64              `json:"rating"`
	GoogleMapsUri                string               `json:"googleMapsUri"`
	WebsiteUri                   string               `json:"websiteUri"`
	Reviews                      []Review             `json:"reviews"`
	RegularOpeningHours          OpeningHours         `json:"regularOpeningHours"`
	Photos                       []Photo              `json:"photos"`
	AdrFormatAddress             string               `json:"adrFormatAddress"`
	BusinessStatus               string               `json:"businessStatus"`
	PriceLevel                   string               `json:"priceLevel"`
	Attributions                 []Attribution        `json:"attributions"`
	IconMaskBaseUri              string               `json:"iconMaskBaseUri"`
	IconBackgroundColor          string               `json:"iconBackgroundColor"`
	CurrentOpeningHours          OpeningHours         `json:"currentOpeningHours"`
	CurrentSecondaryOpeningHours []OpeningHours       `json:"currentSecondaryOpeningHours"`
	RegularSecondaryOpeningHours []OpeningHours       `json:"regularSecondaryOpeningHours"`
	EditorialSummary             LocalizedText        `json:"editorialSummary"`
	PaymentOptions               PaymentOptions       `json:"paymentOptions"`
	ParkingOptions               ParkingOptions       `json:"parkingOptions"`
	SubDestinations              []SubDestination     `json:"subDestinations"`
	FuelOptions                  FuelOptions          `json:"fuelOptions"`
	EVChargeOptions              EVChargeOptions      `json:"evChargeOptions"`
	GenerativeSummary            GenerativeSummary    `json:"generativeSummary"`
	AreaSummary                  AreaSummary          `json:"areaSummary"`
	UtcOffsetMinutes             int                  `json:"utcOffsetMinutes"`
	UserRatingCount              int                  `json:"userRatingCount"`
	Takeout                      bool                 `json:"takeout"`
	Delivery                     bool                 `json:"delivery"`
	DineIn                       bool                 `json:"dineIn"`
	CurbsidePickup               bool                 `json:"curbsidePickup"`
	Reservable                   bool                 `json:"reservable"`
	ServesBreakfast              bool                 `json:"servesBreakfast"`
	ServesLunch                  bool                 `json:"servesLunch"`
	ServesDinner                 bool                 `json:"servesDinner"`
	ServesBeer                   bool                 `json:"servesBeer"`
	ServesWine                   bool                 `json:"servesWine"`
	ServesBrunch                 bool                 `json:"servesBrunch"`
	ServesVegetarianFood         bool                 `json:"servesVegetarianFood"`
	OutdoorSeating               bool                 `json:"outdoorSeating"`
	LiveMusic                    bool                 `json:"liveMusic"`
	MenuForChildren              bool                 `json:"menuForChildren"`
	ServesCocktails              bool                 `json:"servesCocktails"`
	ServesDessert                bool                 `json:"servesDessert"`
	ServesCoffee                 bool                 `json:"servesCoffee"`
	GoodForChildren              bool                 `json:"goodForChildren"`
	AllowsDogs                   bool                 `json:"allowsDogs"`
	Restroom                     bool                 `json:"restroom"`
	GoodForGroups                bool                 `json:"goodForGroups"`
	GoodForWatchingSports        bool                 `json:"goodForWatchingSports"`
	AccessibilityOptions         AccessibilityOptions `json:"accessibilityOptions"`
}

func (p Place) String() string {
	baseStr := fmt.Sprintf("%s\n%s\n%s\n%s\n%f %f\n",
		p.DisplayName.Text, p.FormattedAddress, p.Name, p.PrimaryType,
		p.Location.Latitude, p.Location.Longitude)

	return baseStr
}
