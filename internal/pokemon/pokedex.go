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
