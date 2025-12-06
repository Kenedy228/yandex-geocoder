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

func (r *Response) Coordinates() (*Coordinates, error) {
	split := strings.Split(r.pos(), " ")

	long, err := strconv.ParseFloat(split[0], 64)

	if err != nil {
		return nil, err
	}

	lat, err := strconv.ParseFloat(split[1], 64)

	if err != nil {
		return nil, err
	}

	return &Coordinates{Latitude: lat, Longitude: long}, nil
}

func (r *Response) IsPrecised() bool {
	if r.precision() != "exact" || r.kind() != "house" {
		return false
	}

	if r.suggestion() != "" {
		return false
	}

	return true
}

func (r *Response) FormattedAddress() string {
	if r.precision() != "exact" || r.kind() != "house" {
		return ""
	}

	return r.formatted()
}

func (r *Response) precision() string {
	return r.Response.GeoObjectCollection.
		FeatureMember[0].
		GeoObject.
		MetaDataProperty.
		GeocoderMetaData.
		Precision
}

func (r *Response) kind() string {
	return r.Response.GeoObjectCollection.
		FeatureMember[0].
		GeoObject.
		MetaDataProperty.
		GeocoderMetaData.
		Kind
}

func (r *Response) suggestion() string {
	return r.Response.GeoObjectCollection.
		MetaDataProperty.
		GeocoderResponseMetaData.
		Suggest
}

func (r *Response) pos() string {
	return r.Response.GeoObjectCollection.
		FeatureMember[0].
		GeoObject.
		Point.
		Pos
}

func (r *Response) formatted() string {
	return r.Response.GeoObjectCollection.
		FeatureMember[0].
		GeoObject.
		MetaDataProperty.
		GeocoderMetaData.
		Address.
		Formatted
}
