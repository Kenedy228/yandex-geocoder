package geocoder

import (
	"encoding/json"
	"net/http"
)

type Client struct {
}

func NewClient() *Client {
	return &Client{}
}

func (client *Client) Search(params *SearchParams) (*Response, error) {
	url := params.encode() 

	resp, err := http.Get(url)

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
