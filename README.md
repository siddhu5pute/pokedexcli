# Pokedexcli

A Pokedex CLI app built from scratch in Go. Explore locations, catch Pokemon, battle wild ones, and climb the leaderboard — all from your terminal.

## Requirements

- [Go](https://go.dev/dl/) 1.21 or higher

## Installation
```bash
git clone https://github.com/siddhu5pute/pokedexcli.git
cd pokedexcli
```

## Usage

**Option 1 — Run directly:**
```bash
go run .
```

**Option 2 — Build and run:**
```bash
go build .
./pokedexcli
```

## Commands

| Command | Description |
|---|---|
| `help` | List all available commands |
| `exit` | Exit the Pokedex |
| `map` | Show next page of location areas |
| `mapb` | Go back to previous page of locations |
| `explore <location>` | List Pokemon found in a location |
| `catch <pokemon>` | Try to catch a Pokemon |
| `inspect <pokemon>` | View stats of a caught Pokemon |
| `pokedex` | List all your caught Pokemon |
| `battle <pokemon>` | Battle a wild Pokemon with your first caught one |
| `leaderboard` | View the top trainers |

## Features

- Location browsing with pagination
- Pokemon catching with randomized catch rate based on base experience
- Badge system — earn a badge every 5 catches
- Battle simulator
- Persistent trainer data with leaderboard
- Response caching to avoid redundant API calls

## Author

Built from scratch by [siddhu5pute](https://github.com/siddhu5pute)
