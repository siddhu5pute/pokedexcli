package trainer

type PokemonData struct {
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
	Height         int    `json:"height"`
	ID             int    `json:"id"`
	Weight         int    `json:"weight"`
	Stats          []struct {
		BaseStat int `json:"base_stat"`
		Stat     struct {
			Name string `json:"name"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Type struct {
			Name string `json:"name"`
		} `json:"type"`
	} `json:"types"`
}

type Trainer struct {
	Name    string                 `json:"name"`
	Pokedex map[string]PokemonData `json:"pokedex"`
	Badges  int                    `json:"badges"`
	Caught  int                    `json:"caught"`
	Rival   string                 `json:"rival"`
}

func NewTrainer(name, rival string) *Trainer {
	return &Trainer{
		Name:    name,
		Pokedex: map[string]PokemonData{},
		Rival:   rival,
	}
}
