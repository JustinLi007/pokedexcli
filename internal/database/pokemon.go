package database

import "errors"

var ErrNotExist = errors.New("Resource does not exist")

type Pokemon struct {
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
	Height         int    `json:"height"`
	Stats          []struct {
		BaseStat int `json:"base_stat"`
		Effort   int `json:"effort"`
		Stat     struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"type"`
	} `json:"types"`
	Weight int `json:"weight"`
}

func (db *PokedexDB) AddPokemon(pokemonName string, pokemon Pokemon) error {
	pokedexStruct, err := db.LoadDB()
	if err != nil {
		return err
	}

	// WARN: shallow copy, care Pokemon struct Field types
	newPokemon := pokemon

	// WARN: overwrite dups
	pokedexStruct.Pokedex[pokemonName] = newPokemon

	return db.WriteDB(pokedexStruct)
}

func (db *PokedexDB) GetPokemon(pokemonName string) (Pokemon, error) {
	pokedexStruct, err := db.LoadDB()
	if err != nil {
		return Pokemon{}, err
	}

	pokemon, ok := pokedexStruct.Pokedex[pokemonName]
	if !ok {
		return Pokemon{}, ErrNotExist
	}

	return pokemon, nil
}
