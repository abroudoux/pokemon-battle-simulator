package controllers

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/abroudoux/pokemon-battle-simulator/internal/models"
	"github.com/gin-gonic/gin"
)

const INVALID_ID string = "Invalid ID"
var Pokedex []models.Pokemon

func InitPokedex() {
    file, err := os.Open("internal/data/pokedex.json")

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

func FindPokemonById(id int) models.Pokemon {
	if id > len(Pokedex) || id < 1 {
		fmt.Println("Pokemon not found")
		return models.Pokemon{}
	}

	for _, pokemon := range Pokedex {
		if pokemon.Id == id {
			return pokemon
		}
	}

	return models.Pokemon{}
}

func FindPokemonByName(name string) models.Pokemon {
	for _, pokemon := range Pokedex {
		if pokemon.Name.French == name {
			return pokemon
		}
		
		if pokemon.Name.English == name {
			return pokemon
		}

		if pokemon.Name.Japanese == name {
			return pokemon
		}

		if pokemon.Name.Chinese == name {
			return pokemon
		}
	}

	return models.Pokemon{}
}

func GetPokemon(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))

    if err != nil {
        c.JSON(400, gin.H{"error": INVALID_ID})
        return
    }

    pokemon := FindPokemonById(id)

	if pokemon.Id != 0 {
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

	pokemon1 := FindPokemonById(id1)
	pokemon2 := FindPokemonById(id2)

	winner := FindHighestSpeed(pokemon1, pokemon2)

	c.JSON(200, gin.H{"pokemon battle": []models.Pokemon{pokemon1, pokemon2}, "winner": winner})
}