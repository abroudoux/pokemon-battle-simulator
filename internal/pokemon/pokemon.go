package pokemon

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"

	types "github.com/abroudoux/pokemon-battle-simulator/internal/types"
	utils "github.com/abroudoux/pokemon-battle-simulator/internal/utils"
	"github.com/gin-gonic/gin"
)

var Pokedex []types.Pokemon
const INVALID_ID string = "Invalid ID"
const POKEMON_NOT_FOUND string = "Pokemon not found"
const FILE_PATH string = "../../internal/data/pokedex.json"

func InitPokedex() {
    file, err := os.Open(FILE_PATH)

    if err != nil {
        fmt.Println(err, "Error opening file")
        return
    }

    defer file.Close()
    err = json.NewDecoder(file).Decode(&Pokedex)

    if err != nil {
        fmt.Println(err, "Error decoding file")
        return
    }

    fmt.Println("Pokedex initialized with", len(Pokedex), "pokemons")
}

func findHighestSpeed(pokemon1 types.Pokemon, pokemon2 types.Pokemon) types.Pokemon {
	if pokemon1.Base.Speed > pokemon2.Base.Speed {
		return pokemon1
	}

	return pokemon2
}

func findWinner(pokemon1 types.Pokemon, pokemon2 types.Pokemon) (float64) {
	attacker := findHighestSpeed(pokemon1, pokemon2)

	var defender types.Pokemon

	if attacker.Id == pokemon1.Id {
		defender = pokemon2
	} else {
		defender = pokemon1
	}

	multiplier := calculteMultiplier(attacker, defender)

	return multiplier
}

func calculteMultiplier(attacker types.Pokemon, defender types.Pokemon) float64 {
	multiplier := 1.0

	for _, attackerType := range attacker.Type {
		multiplier *= calculateWeaknessesMultiplier(attackerType, defender.Type)
		multiplier *= calculateResistancesMultiplier(attackerType, defender.Type)
		multiplier *= calculateImmunitiesMultiplier(attackerType, defender.Type)
	}

	return multiplier
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


func findPokemonById(id int) (types.Pokemon, error) {
	if id > len(Pokedex) || id < 1 {
		return types.Pokemon{}, errors.New(INVALID_ID)
	}

	for _, pokemon := range Pokedex {
		if pokemon.Id == id {
			return pokemon, nil
		}
	}

	return types.Pokemon{}, errors.New(POKEMON_NOT_FOUND)
}

func findPokemonByName(name string) (types.Pokemon, error) {
	for _, pokemon := range Pokedex {
		if pokemon.Name.French == name ||
			pokemon.Name.English == name ||
			pokemon.Name.Japanese == name ||
			pokemon.Name.Chinese == name {

			return pokemon, nil
		}
	}
	return types.Pokemon{}, errors.New(POKEMON_NOT_FOUND)
}

func findPokemon(param string) (types.Pokemon, error) {
	paramType := utils.CheckParamType(param)

	if paramType == "mixed" {
		return types.Pokemon{}, errors.New("Invalid parameter: mixed characters")
	}

	if paramType == "digit" {
		id, err := strconv.Atoi(param)

		if err != nil {
			return types.Pokemon{}, errors.New("Invalid ID")
		}

		return findPokemonById(id)
	}

	if paramType == "letter" {
		param = utils.CapitalizeFirstLetter(param)

		return findPokemonByName(param)
	}

	return types.Pokemon{}, errors.New("Invalid parameter type")
}

func GetPokedex(c *gin.Context) {
    if len(Pokedex) == 0 {
        c.JSON(404, gin.H{"error": "No Pokemon in Pokedex"})
        return
    }

	if Pokedex == nil {
		c.JSON(404, gin.H{"error": "Pokedex not initialized"})
		return
	}

    c.JSON(200, gin.H{"Pokedex": Pokedex})
}

func GetPokemon(c *gin.Context) {
	param := c.Param("pokemon")
	pokemon, err := findPokemon(param)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"pokemon": pokemon})
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

	fastest := findHighestSpeed(pokemon1, pokemon2)
	mutiplier := findWinner(pokemon1, pokemon2)

	c.JSON(200, gin.H{"fastest": fastest, "pokemon1": pokemon1, "pokemon2": pokemon2, "multiplier": mutiplier},)
}
