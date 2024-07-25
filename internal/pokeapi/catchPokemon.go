package pokeapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/JustinLi007/pokedexcli/internal/database"
)

func (c *Client) CatchPokemon(pokemonName string) (database.Pokemon, error) {
	url := baseURL + "/pokemon"

	if len(strings.Fields(pokemonName)) == 0 {
		return database.Pokemon{}, errors.New("Pokemon name not specified")
	}

	url = fmt.Sprintf("%v/%v", url, pokemonName)

	data, ok := c.cache.Get(url)
	if !ok {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return database.Pokemon{}, err
		}

		resp, err := c.httpClient.Do(req)
		if err != nil {
			return database.Pokemon{}, err
		}
		defer resp.Body.Close()

		if resp.StatusCode > 399 {
			return database.Pokemon{}, fmt.Errorf("Failed to retrieve information about %v\n", pokemonName)
		}

		data, err = io.ReadAll(resp.Body)
		if err != nil {
			return database.Pokemon{}, err
		}

		c.cache.Add(url, data)
	}

	pokemonResp := database.Pokemon{}
	if err := json.Unmarshal(data, &pokemonResp); err != nil {
		return database.Pokemon{}, err
	}

	return pokemonResp, nil
}
