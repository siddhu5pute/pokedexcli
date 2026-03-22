package main

import (
	"github.com/siddhu5pute/pokedexcli/internal/pokecache"
	"github.com/siddhu5pute/pokedexcli/internal/trainer"
)

type locationAreaResponse struct {
	Count    int            `json:"count"`
	Next     *string        `json:"next"`
	Previous *string        `json:"previous"`
	Results  []locationArea `json:"results"`
}

type locationArea struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type config struct {
	next     *string
	previous *string
	cache    *pokecache.Cache
	trainer  *trainer.Trainer
}

type pokemonEncounters struct {
	Encounter []struct {
		PokeDetail struct {
			Name string `json:"name"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

const UrlLocation = "https://pokeapi.co/api/v2/location-area/"
const UrlPokemon = "https://pokeapi.co/api/v2/pokemon/"
