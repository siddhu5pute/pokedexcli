package storage

import (
	"os"
	"sort"
	"strings"

	"github.com/siddhu5pute/pokedexcli/internal/trainer"
)

func GetLeaderboard() ([]*trainer.Trainer, error) {
	files, err := os.ReadDir("storage")
	if err != nil {
		return nil, err
	}

	var trainers []*trainer.Trainer
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		name := strings.TrimSuffix(file.Name(), ".json")
		t, err := LoadTrainer(name)
		if err != nil {
			continue
		}
		if t != nil {
			trainers = append(trainers, t)
		}
	}

	sort.Slice(trainers, func(i, j int) bool {
		return trainers[i].Caught > trainers[j].Caught
	})

	return trainers, nil
}
