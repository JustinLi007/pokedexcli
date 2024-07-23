package pokeapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type LocationDetail struct {
	PokemonEncounters []PokemonEncounters `json:"pokemon_encounters"`
}

type PokemonEncounters struct {
	Pokemon Pokemon `json:"pokemon"`
}

type Pokemon struct {
	Name string `json:"name"`
}

func (c *Client) ListLocationDetails(locationName string) (LocationDetail, error) {
	url := baseURL + "/location-area"

	if locationName == "" {
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

		data, err = io.ReadAll(resp.Body)
		if err != nil {
			return LocationDetail{}, err
		}

		c.cache.Add(url, data)
	}

	encounter := LocationDetail{}
	if err := json.Unmarshal(data, &encounter); err != nil {
		return LocationDetail{}, err
	}

	return encounter, nil
}
