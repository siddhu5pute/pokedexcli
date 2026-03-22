package main

import "strings"

func cleanInput(text string) []string {
	text = strings.ToLower(text)
	text = strings.TrimSpace(text)
	return strings.Fields(text)
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, []string) error
}

// 🔴 global variable (DECLARED only)
var commands map[string]cliCommand

// 🔴 init runs automatically before main()
func init() {
	commands = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Display the location nearby",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Display the location nearby backwards",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "It provides the pokemons nearby loaction/Area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "It will try to catch the Pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "It will provide the relavant data abt ur Pokemon!",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "It will display the Pokemons that u posses!",
			callback:    commandPokedex,
		},
		"leaderboard": {
			name:        "leaderboard",
			description: "It provides u the current ranking of Trainers...",
			callback:    commandLeaderboard,
		},
		"battle": {
			name:        "battle",
			description: "Battle a wild Pokemon",
			callback:    commandBattle,
		},
	}
}
