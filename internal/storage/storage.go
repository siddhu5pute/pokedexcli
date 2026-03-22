package storage

import (
	"encoding/json"
	"os"

	"github.com/siddhu5pute/pokedexcli/internal/trainer"
)

func SaveTrainer(trainer *trainer.Trainer, name string) error {
	err := os.MkdirAll("storage", 0755)
	if err != nil {
		return err
	}
	data, err := json.Marshal(trainer)
	if err != nil {
		return err
	}
	err = os.WriteFile("storage/"+name+".json", data, 0644)
	if err != nil {
		return err
	}
	return nil
}

func LoadTrainer(name string) (*trainer.Trainer, error) {
	_, err := os.Stat("storage/" + name + ".json")
	if os.IsNotExist(err) {
		return nil, nil
	}
	data, err := os.ReadFile("storage/" + name + ".json")
	if err != nil {
		return nil, err
	}
	var trainer *trainer.Trainer
	err = json.Unmarshal(data, &trainer)
	if err != nil {
		return nil, err
	}
	return trainer, nil
}
