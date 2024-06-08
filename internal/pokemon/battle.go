package pokemon

import (
	"errors"
	"net/http"

	types "github.com/abroudoux/pokemon-battle-simulator/internal/types"
	"github.com/gin-gonic/gin"
)

const INVALID_ID string = "Invalid ID"
const POKEMON_NOT_FOUND string = "Pokemon not found"
const FILE_PATH string = "../../internal/data/pokedex.json"

func findHighestSpeed(pokemon1 types.Pokemon, pokemon2 types.Pokemon) (types.Pokemon, error) {
	if pokemon1.Base.Speed > pokemon2.Base.Speed {
		return pokemon1, nil
	}

	return pokemon2, nil
}

func calculteMultiplier(attacker types.Pokemon, defender types.Pokemon) (float64, error) {
	multiplier := 1.0

	for _, attackerType := range attacker.Type {
		multiplier *= calculateWeaknessesMultiplier(attackerType, defender.Type)
		multiplier *= calculateResistancesMultiplier(attackerType, defender.Type)
		multiplier *= calculateImmunitiesMultiplier(attackerType, defender.Type)
	}

	return multiplier, nil
}

func calculateWeaknessesMultiplier(attackerType types.TypeData, defenderTypes []types.TypeData) float64 {
	multiplier := 1.0

	for _, defenderType := range defenderTypes {
		for _, weakness := range attackerType.Weaknessess {
			if weakness.TypeId == defenderType.Id {
				multiplier *= 2
			}
		}
	}

	return multiplier
}

func calculateResistancesMultiplier(attackerType types.TypeData, defenderTypes []types.TypeData) float64 {
	multiplier := 1.0

	for _, defenderType := range defenderTypes {
		for _, resistance := range attackerType.Resistances {
			if resistance.TypeId == defenderType.Id {
				multiplier *= 0.5
			}
		}
	}

	return multiplier
}

func calculateImmunitiesMultiplier(attackerType types.TypeData, defenderTypes []types.TypeData) float64 {
	multiplier := 1.0

	for _, defenderType := range defenderTypes {
		for _, immunity := range attackerType.Immunities {
			if immunity.TypeId == defenderType.Id {
				multiplier *= 0
			}
		}
	}

	return multiplier
}

func findHighestStatAttacker(attacker types.Pokemon) (highestStatAttacker string) {
	if attacker.Base.Attack > attacker.Base.SpAttack {
		return "Attack"
	}

	return "SpAttack"
}

func findWinner(pokemon1 types.Pokemon, pokemon2 types.Pokemon) (types.Pokemon, error) {
	attacker, err := findHighestSpeed(pokemon1, pokemon2)
	var defender types.Pokemon

	if err != nil {
		return types.Pokemon{}, err
	}

	if attacker.Id == pokemon1.Id {
		defender = pokemon2
	} else {
		defender = pokemon1
	}

	multiplier, err := calculteMultiplier(attacker, defender)

	if err != nil {
		return types.Pokemon{}, err
	}

	highestStatAttacker := findHighestStatAttacker(attacker)

	if highestStatAttacker == "Attack" {
		if float64(attacker.Base.Attack) * multiplier > (float64(defender.Base.HP) + float64(defender.Base.Defense)) {
			return attacker, err
		}
		return defender, err
	} else if highestStatAttacker == "SpAttack" {
		if float64(attacker.Base.SpAttack) * multiplier > (float64(defender.Base.HP) + float64(defender.Base.SpDefense)) {
			return attacker, err
		}

		return defender, err
	}

	return types.Pokemon{}, errors.New("Error finding winner")
}

func CreateBattle(c *gin.Context) {
	param1 := c.Param("pokemon1")
	param2 := c.Param("pokemon2")

	pokemon1, err := findPokemon(param1)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	pokemon2, err := findPokemon(param2)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	winnerFirstTurn, err := findWinner(pokemon1, pokemon2)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"pokemon1": pokemon1,"pokemon2": pokemon2 , "winnerFirstTurn": winnerFirstTurn},)
}
