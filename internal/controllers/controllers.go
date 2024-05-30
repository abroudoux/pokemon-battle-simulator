package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/abroudoux/pokemon-battle-simulator/internal/models"
	"github.com/abroudoux/pokemon-battle-simulator/internal/utils"
	"github.com/gin-gonic/gin"
)

var Pokedex []models.Pokemon
const INVALID_ID string = "Invalid ID"
const POKEMON_NOT_FOUND string = "Pokemon not found"

func InitPokedex() {
    file, err := os.Open("../../internal/data/pokedex.json")

    if err != nil {
        fmt.Println(err, "Error opening file")
        return
    }

    defer file.Close()

    err = json.NewDecoder(file).Decode(&Pokedex)

    if err != nil {
        fmt.Println(err, "Error decoding JSON")
        return
    }

    fmt.Println("Pokedex initialized with", len(Pokedex), "pokemons")
}

func GetPokedex(c *gin.Context) {
    if len(Pokedex) == 0 {
        c.JSON(404, gin.H{"error": "No Pokemon in Pokedex"})
        return
    }

    c.JSON(200, gin.H{"data": Pokedex})
}

func FindPokemonById(id int) (models.Pokemon, error) {
	if id > len(Pokedex) || id < 1 {
		return models.Pokemon{}, errors.New(POKEMON_NOT_FOUND)
	}

	for _, pokemon := range Pokedex {
		if pokemon.Id == id {
			return pokemon, nil
		}
	}

	return models.Pokemon{}, errors.New(POKEMON_NOT_FOUND)
}

func FindPokemonByName(name string) (models.Pokemon, error) {
	for _, pokemon := range Pokedex {
		if pokemon.Name.French == name ||
			pokemon.Name.English == name ||
			pokemon.Name.Japanese == name ||
			pokemon.Name.Chinese == name {
			return pokemon, nil
		}
	}
	return models.Pokemon{}, errors.New(POKEMON_NOT_FOUND)
}

func GetPokemon(c *gin.Context) {
	param := c.Param("pokemon")

	paramType := utils.CheckString(param)

	if paramType == "mixed" {
		c.JSON(400, gin.H{"error": "Invalid parameter: mixed characters"})
		return
	}

	if paramType == "digit" {
		paramInt, err := strconv.Atoi(param)
		if err != nil {
			c.JSON(400, gin.H{"error": INVALID_ID})
			return
		}

		pokemon, err := FindPokemonById(paramInt)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{"pokemon": pokemon})
		return
	}

	if paramType == "letter" {
		pokemon, err := FindPokemonByName(param)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, gin.H{"pokemon": pokemon})
		return
	}
}

func FindHighestSpeed(pokemon1 models.Pokemon, pokemon2 models.Pokemon) models.Pokemon {
	if pokemon1.Base.Speed > pokemon2.Base.Speed {
		return pokemon1
	}

	return pokemon2
}

func CreateBattle(c *gin.Context) {
	id1, err := strconv.Atoi(c.Param("pokemon1"))

	if err != nil {
		c.JSON(400, gin.H{"error": INVALID_ID})
		return
	}

	id2, err := strconv.Atoi(c.Param("pokemon2"))

	if err != nil {
		c.JSON(400, gin.H{"error": INVALID_ID})
		return
	}

	pokemon1, err := FindPokemonById(id1)

	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}

	pokemon2, err := FindPokemonById(id2)

	if err != nil {
		c.JSON(404, gin.H{"error": err.Error()})
		return
	}

	winner := FindHighestSpeed(pokemon1, pokemon2)

	c.JSON(200, gin.H{"pokemon battle": []models.Pokemon{pokemon1, pokemon2}, "winner": winner})
}

// func Test(c *gin.Context) {
// 	s := c.Param("s")

// 	sType := utils.CheckString(s)

// 	if sType == "mixed" {
// 		c.JSON(400, gin.H{"test": "mixed"})
// 		return
// 	}

// 	if sType == "digit" {
// 		sInt, err := strconv.Atoi(s)

// 		if err != nil {
// 			c.JSON(400, gin.H{"error": INVALID_ID})
// 			return
// 		}

// 		pokemon := FindPokemonById(sInt)

// 		if pokemon.Id != 0 {
// 			c.JSON(200, gin.H{"pokemon": pokemon})
// 			return
// 		}

// 		c.JSON(200, gin.H{"pokemon": pokemon})
// 		return
// 	}

// 	if sType == "letter" {
// 		pokemon := FindPokemonByName(s)
// 		c.JSON(200, gin.H{"pokemon": pokemon})
// 		return 
// 	}
// }