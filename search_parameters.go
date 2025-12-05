package geocoder

import (
	"fmt"
	"net/url"
)

type SearchParams struct {
	ApiKey   string
	Language string
	Geocode  string
	Format   string
	Count    string
}

func NewParamsWithDefaults(apiKey string, geocode string) *SearchParams {
	return &SearchParams{
		ApiKey: apiKey,
		Language: "ru_RU",
		Geocode:  geocode,
		Format:   "json",
		Count:    "1",
	}
}

func (p *SearchParams) encode() string {
	return fmt.Sprintf(
		"%s?apikey=%s&lang=%s&geocode=%s&format=%s&results=%s",
		baseURL,
		url.QueryEscape(p.ApiKey),
		url.QueryEscape(p.Language),
		url.QueryEscape(p.Geocode),
		url.QueryEscape(p.Format),
		url.QueryEscape(p.Count))
}
