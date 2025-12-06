package geocoder

import (
	"strconv"
	"strings"
)

type Response struct {
	Response struct {
		GeoObjectCollection struct {
			MetaDataProperty struct {
				GeocoderResponseMetaData struct {
					Request string `json:"request"`
					Found   string `json:"found"`
					Results string `json:"results"`
					Fix     string `json:"fix,omitempty"`
					Suggest string `json:"suggest,omitempty"`
					Skip    string `json:"skip,omitempty"`
				} `json:"GeocoderResponseMetaData"`
			} `json:"metaDataProperty"`
			FeatureMember []struct {
				GeoObject struct {
					MetaDataProperty struct {
						GeocoderMetaData struct {
							Precision string `json:"precision"`
							Kind      string `json:"kind"`
							Text      string `json:"text"`
							Address   struct {
								CountryCode string `json:"country_code"`
								Formatted   string `json:"formatted"`
								PostalCode  string `json:"postal_code,omitempty"`
								Components  []struct {
									Kind string `json:"kind"`
									Name string `json:"name"`
								} `json:"Components"`
							} `json:"Address"`
						} `json:"GeocoderMetaData"`
					} `json:"metaDataProperty"`
					Name        string `json:"name,omitempty"`
					Description string `json:"description,omitempty"`
					BoundedBy   struct {
						Envelope struct {
							LowerCorner string `json:"lowerCorner"`
							UpperCorner string `json:"upperCorner"`
						} `json:"Envelope"`
					} `json:"boundedBy"`
					URI   string `json:"uri,omitempty"`
					Point struct {
						Pos string `json:"pos"`
					} `json:"Point"`
				} `json:"GeoObject"`
			} `json:"featureMember"`
		} `json:"GeoObjectCollection"`
	} `json:"response"`
}

type Coordinates struct {
	Longitude float64
	Latitude  float64
}

type Precision int

const (
	PrecisionExact Precision = iota
	PrecisionNumber
	PrecisionNear
	PrecisionRange
	PrecisionStreet
	PrecisionOther
)

type DataType int

const (
	TypeHouse DataType = iota
	TypeStreet
	TypeMetro
	TypeDistrict
	TypeLocality
	TypeArea
	TypeProvince
	TypeCountry
	TypeRegion
	TypeHydro
	TypeRailwayStation
	TypeStation
	TypeRoute
	TypeVegetation
	TypeAirport
	TypeEntrance
	TypeOther
)

func (r *Response) Coordinates() (*Coordinates, error) {
	pos := r.Response.GeoObjectCollection.
		FeatureMember[0].
		GeoObject.
		Point.
		Pos

	splitted := strings.Split(pos, " ")

	long, err := strconv.ParseFloat(splitted[0], 64)

	if err != nil {
		return nil, err
	}

	lat, err := strconv.ParseFloat(splitted[1], 64)

	if err != nil {
		return nil, err
	}

	return &Coordinates{Latitude: lat, Longitude: long}, nil
}

func (r *Response) Precision() Precision {
	precision := r.Response.GeoObjectCollection.
		FeatureMember[0].
		GeoObject.
		MetaDataProperty.
		GeocoderMetaData.
		Precision

	switch precision {
	case "exact":
		return PrecisionExact
	case "number":
		return PrecisionNumber
	case "near":
		return PrecisionNear
	case "range":
		return PrecisionRange
	case "street":
		return PrecisionStreet
	default:
		return PrecisionOther
	}
}

func (r *Response) DataType() DataType {
	kind := r.Response.GeoObjectCollection.
		FeatureMember[0].
		GeoObject.
		MetaDataProperty.
		GeocoderMetaData.
		Kind

	switch kind {
	case "house":
		return TypeHouse
	case "street":
		return TypeStreet
	case "metro":
		return TypeMetro
	case "district":
		return TypeDistrict
	case "locality":
		return TypeLocality
	case "area":
		return TypeArea
	case "province":
		return TypeProvince
	case "country":
		return TypeCountry
	case "region":
		return TypeRegion
	case "hydro":
		return TypeHydro
	case "railway_station":
		return TypeRailwayStation
	case "station":
		return TypeStation
	case "route":
		return TypeRoute
	case "vegetation":
		return TypeVegetation
	case "airport":
		return TypeAirport
	case "entrance":
		return TypeEntrance
	default:
		return TypeOther
	}
}
