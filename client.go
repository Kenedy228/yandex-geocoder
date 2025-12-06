package geocoder

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type Client struct {
}

func NewClient() *Client {
	return &Client{}
}

func (client *Client) Search(params *SearchParams) (*Response, error) {
	preparedURL := fmt.Sprintf("%s?apikey=%s&geocode=%s&format=%s&lang=%s&results=%s",
		baseURL,
		url.QueryEscape(params.ApiKey),
		url.QueryEscape(params.Geocode),
		url.QueryEscape(format),
		url.QueryEscape(language),
		url.QueryEscape(count))

	resp, err := http.Get(preparedURL)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, err
	}

	decoder := json.NewDecoder(resp.Body)
	var data Response

	if err := decoder.Decode(&data); err != nil {
		return nil, err
	}

	return &data, nil
}
