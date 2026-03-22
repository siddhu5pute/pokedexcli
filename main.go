package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/siddhu5pute/pokedexcli/internal/pokecache"
	"github.com/siddhu5pute/pokedexcli/internal/storage"
	"github.com/siddhu5pute/pokedexcli/internal/trainer"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter your trainer name and rival: ")
	scanner.Scan()
	text := scanner.Text()
	words := strings.Fields(text)
	name := words[0]
	rival := words[1]

	conf := config{
		cache: pokecache.NewCache(5 * time.Second),
	}

	existingTrainer, err := storage.LoadTrainer(name)
	if err != nil {
		fmt.Println("Error loading trainer:", err)
		os.Exit(1)
	}

	if existingTrainer != nil {
		fmt.Printf("Welcome back %v!\n", name)
		conf.trainer = existingTrainer
	} else {
		fmt.Printf("Welcome new trainer %v!\n", name)
		conf.trainer = trainer.NewTrainer(name, rival)
		if err := storage.SaveTrainer(conf.trainer, name); err != nil {
			fmt.Println("Error saving trainer:", err)
			os.Exit(1)
		}
	}

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		text := scanner.Text()
		words := cleanInput(text)
		if len(words) == 0 {
			continue
		}
		commandName := words[0]
		cmd, exists := commands[commandName]
		if !exists {
			fmt.Println("Unknown command")
			continue
		}
		if err := cmd.callback(&conf, words[1:]); err != nil {
			fmt.Println(err)
		}
	}
}
