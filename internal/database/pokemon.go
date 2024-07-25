package database

type Pokemon struct {
	BaseExperience int `json:"base_experience"`
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
