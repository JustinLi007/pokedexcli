package pokeapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func (c *Client) GetLocationDetails(locationName string) (LocationDetail, error) {
	url := baseURL + "/location-area"

	if len(strings.Fields(locationName)) == 0 {
		return LocationDetail{}, errors.New("Location name not specified")
	}

	url = fmt.Sprintf("%v/%v", url, locationName)

	data, ok := c.cache.Get(url)
	if !ok {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return LocationDetail{}, err
		}

		resp, err := c.httpClient.Do(req)
		if err != nil {
			return LocationDetail{}, nil
		}
		defer resp.Body.Close()

		if resp.StatusCode > 399 {
			return LocationDetail{}, fmt.Errorf("Failed to retrieve details about %v\n", locationName)
		}

		data, err = io.ReadAll(resp.Body)
		if err != nil {
			return LocationDetail{}, err
		}

		c.cache.Add(url, data)
	}

	locationResp := LocationDetail{}
	if err := json.Unmarshal(data, &locationResp); err != nil {
		return LocationDetail{}, err
	}

	return locationResp, nil
}
