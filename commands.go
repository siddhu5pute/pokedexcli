package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"

	"github.com/siddhu5pute/pokedexcli/internal/battle"
	"github.com/siddhu5pute/pokedexcli/internal/storage"
	"github.com/siddhu5pute/pokedexcli/internal/trainer"
)

func commandExit(conf *config, location []string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(conf *config, location []string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	for _, cmd := range commands {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}

func fetchAndDecode(url string, conf *config) (locationAreaResponse, error) {
	if data, ok := conf.cache.Get(url); ok {
		fmt.Println("Cache hit!")
		var locaArea locationAreaResponse
		err := json.Unmarshal(data, &locaArea)
		return locaArea, err
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return locationAreaResponse{}, fmt.Errorf("error creating request: %v", err)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return locationAreaResponse{}, fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return locationAreaResponse{}, fmt.Errorf("error reading body: %v", err)
	}

	conf.cache.Add(url, data)

	var locaArea locationAreaResponse
	err = json.Unmarshal(data, &locaArea)
	return locaArea, err
}

func commandMap(conf *config, location []string) error {
	urlForCurrentPage := UrlLocation
	if conf.next != nil {
		urlForCurrentPage = *conf.next
	}
	locaArea, err := fetchAndDecode(urlForCurrentPage, conf)
	if err != nil {
		return err
	}
	for _, locationNames := range locaArea.Results {
		fmt.Println(locationNames.Name)
	}
	conf.next = locaArea.Next
	conf.previous = locaArea.Previous
	return nil
}

func commandMapb(conf *config, location []string) error {
	if conf.previous == nil {
		fmt.Println("you're on the first page")
		return nil
	}
	locaArea, err := fetchAndDecode(*conf.previous, conf)
	if err != nil {
		return err
	}
	for _, locationNames := range locaArea.Results {
		fmt.Println(locationNames.Name)
	}
	conf.next = locaArea.Next
	conf.previous = locaArea.Previous
	return nil
}

func commandExplore(conf *config, location []string) error {
	urlPoke := UrlLocation + location[0] + "/"
	var data []byte
	if cachedData, ok := conf.cache.Get(urlPoke); ok {
		data = cachedData
		fmt.Println("Cache hit!")
	} else {
		req, err := http.NewRequest("GET", urlPoke, nil)
		if err != nil {
			return fmt.Errorf("error creating request: %v", err)
		}
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			return fmt.Errorf("error making request: %v", err)
		}
		defer resp.Body.Close()

		data, err = io.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("error reading body: %v", err)
		}

		conf.cache.Add(urlPoke, data)
	}

	var locaAreaPokes pokemonEncounters
	err := json.Unmarshal(data, &locaAreaPokes)
	if err != nil {
		return err
	}
	for _, PokemonNames := range locaAreaPokes.Encounter {
		fmt.Println(PokemonNames.PokeDetail.Name)
	}
	return nil
}

func commandCatch(conf *config, pokeName []string) error {
	if len(pokeName) == 0 {
		return fmt.Errorf("No Pokemon to Catch...")
	}
	urlCatchedPoke := UrlPokemon + pokeName[0] + "/"
	var data []byte
	if cachedPoke, ok := conf.cache.Get(urlCatchedPoke); ok {
		data = cachedPoke
		fmt.Println("Cache hit!")
	} else {
		req, err := http.NewRequest("GET", urlCatchedPoke, nil)
		if err != nil {
			return fmt.Errorf("error creating request: %v", err)
		}
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			return fmt.Errorf("error making request: %v", err)
		}
		defer resp.Body.Close()
		data, err = io.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("error reading body: %v", err)
		}
		conf.cache.Add(urlCatchedPoke, data)
	}

	var pokemon trainer.PokemonData
	err := json.Unmarshal(data, &pokemon)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %v...\n", pokeName[0])

	if pokemon.BaseExperience == 0 {
		pokemon.BaseExperience = 100
	}

	randNum := rand.Intn(pokemon.BaseExperience)
	if randNum < 50 {
		conf.trainer.Pokedex[pokeName[0]] = pokemon
		conf.trainer.Caught++
		if conf.trainer.Caught%5 == 0 {
			conf.trainer.Badges++
			fmt.Printf("🎉 You earned a new badge! Total badges: %v\n", conf.trainer.Badges)
		}
		fmt.Printf("%v was caught!\n", pokeName[0])
		if err := storage.SaveTrainer(conf.trainer, conf.trainer.Name); err != nil {
			return fmt.Errorf("error saving trainer: %v", err)
		}
	} else {
		fmt.Printf("%v escaped!\n", pokeName[0])
	}
	return nil
}

func commandInspect(conf *config, pokeName []string) error {
	if len(pokeName) == 0 {
		return fmt.Errorf("No Pokemon to Catch...")
	}
	if pokeData, ok := conf.trainer.Pokedex[pokeName[0]]; ok {
		fmt.Println("Name: ", pokeData.Name)
		fmt.Println("Height: ", pokeData.Height)
		fmt.Println("Weight: ", pokeData.Weight)
		fmt.Println("Stats:")
		for _, stat := range pokeData.Stats {
			fmt.Printf("  -%v: %v\n", stat.Stat.Name, stat.BaseStat)
		}
		fmt.Println("Types:")
		for _, t := range pokeData.Types {
			fmt.Printf("  - %v\n", t.Type.Name)
		}
	} else {
		fmt.Println("you have not caught that pokemon")
	}
	return nil
}

func commandPokedex(conf *config, _ []string) error {
	if len(conf.trainer.Pokedex) == 0 {
		fmt.Println("Your Pokedex is empty! Go catch some Pokemon!")
		return nil
	}
	fmt.Println("Your Pokedex:")
	for key := range conf.trainer.Pokedex {
		fmt.Printf(" - %v\n", key)
	}
	return nil
}

func commandLeaderboard(conf *config, _ []string) error {
	trainers, err := storage.GetLeaderboard()
	if err != nil {
		return fmt.Errorf("error getting leaderboard: %v", err)
	}
	fmt.Println("🏆 Top Trainers:")
	for i, t := range trainers {
		fmt.Printf("%v. %v - %v Pokemon | Badges: %v\n", i+1, t.Name, t.Caught, t.Badges)
	}
	return nil
}

func commandBattle(conf *config, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("provide a pokemon name to battle")
	}
	if len(conf.trainer.Pokedex) == 0 {
		return fmt.Errorf("you have no pokemon to battle with! catch some first")
	}

	// get your first pokemon
	var playerPoke trainer.PokemonData
	for _, p := range conf.trainer.Pokedex {
		playerPoke = p
		break
	}

	// fetch wild pokemon
	urlWildPoke := UrlPokemon + args[0] + "/"
	var data []byte
	if cachedPoke, ok := conf.cache.Get(urlWildPoke); ok {
		data = cachedPoke
	} else {
		req, err := http.NewRequest("GET", urlWildPoke, nil)
		if err != nil {
			return fmt.Errorf("error creating request: %v", err)
		}
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			return fmt.Errorf("error making request: %v", err)
		}
		defer resp.Body.Close()
		data, err = io.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("error reading body: %v", err)
		}
		conf.cache.Add(urlWildPoke, data)
	}

	var wildPoke trainer.PokemonData
	if err := json.Unmarshal(data, &wildPoke); err != nil {
		return err
	}

	fmt.Printf("Battle starting: %v vs %v!\n", playerPoke.Name, wildPoke.Name)
	result := battle.SimulateBattle(playerPoke, wildPoke)
	fmt.Println(result.Message)
	return nil
}
