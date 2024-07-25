package database

import (
	"encoding/json"
	"errors"
	"os"
	"sync"
)

type PokedexDB struct {
	path string
	mtx  *sync.RWMutex
}

type PokedexStruct struct {
	Pokedex map[string]Pokemon `json:"pokedex"`
}

func NewPokedexDB(filepath string) (*PokedexDB, error) {
	db := &PokedexDB{
		path: filepath,
		mtx:  &sync.RWMutex{},
	}

	err := db.ensureDB()

	return db, err
}

func (db *PokedexDB) ensureDB() error {
	_, err := os.Open(db.path)
	if errors.Is(err, os.ErrNotExist) {
		return db.createDB()
	}

	return err
}

func (db *PokedexDB) createDB() error {
	pokedex := PokedexStruct{
		Pokedex: make(map[string]Pokemon),
	}

	return db.WriteDB(pokedex)
}

func (db *PokedexDB) LoadDB() (PokedexStruct, error) {
	db.mtx.RLock()
	defer db.mtx.RUnlock()

	pokedexStruct := PokedexStruct{
		Pokedex: make(map[string]Pokemon),
	}

	file, err := os.Open(db.path)
	defer file.Close()

	if err != nil {
		return pokedexStruct, err
	}

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&pokedexStruct); err != nil {
		return pokedexStruct, err
	}

	return pokedexStruct, nil
}

func (db *PokedexDB) WriteDB(pokedexStruct PokedexStruct) error {
	db.mtx.Lock()
	defer db.mtx.Unlock()

	data, err := json.MarshalIndent(pokedexStruct, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(db.path, data, 0666)
}
