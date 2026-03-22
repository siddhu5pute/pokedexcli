package battle

import (
	"fmt"
	"math/rand"

	"github.com/siddhu5pute/pokedexcli/internal/trainer"
)

type BattleResult struct {
	Won     bool
	Message string
}

func SimulateBattle(playerPoke, wildPoke trainer.PokemonData) BattleResult {
	playerHP := getstat(playerPoke, "hp")
	wildHP := getstat(wildPoke, "hp")
	playerAttack := getstat(playerPoke, "attack")
	wildAttack := getstat(wildPoke, "attack")
	playerDefense := getstat(playerPoke, "defense")
	wildDefense := getstat(wildPoke, "defense")

	for playerHP > 0 && wildHP > 0 {
		playerDamage := max(1, playerAttack-wildDefense+rand.Intn(10))
		wildHP -= playerDamage
		fmt.Printf("Your Pokemon deals %v damage! Wild HP: %v\n", playerDamage, max(0, wildHP))

		if wildHP <= 0 {
			break
		}

		wildDamage := max(1, wildAttack-playerDefense+rand.Intn(10))
		playerHP -= wildDamage
		fmt.Printf("Wild Pokemon deals %v damage! Your HP: %v\n", wildDamage, max(0, playerHP))
	}

	if playerHP > 0 {
		return BattleResult{Won: true, Message: "You won the battle!"}
	}
	return BattleResult{Won: false, Message: "You lost the battle!"}
}

func getstat(p trainer.PokemonData, statName string) int {
	for _, s := range p.Stats {
		if s.Stat.Name == statName {
			return s.BaseStat
		}
	}
	return 0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
