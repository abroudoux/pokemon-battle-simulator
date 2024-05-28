package controllers

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/abroudoux/pokemon-battle-simulator/src/models"
	"github.com/gin-gonic/gin"
)

var Pokedex []models.Pokemon

func InitPokedex() {
    file, err := os.Open("data/pokedex.json")
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

func GetPokemon(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))

    if err != nil {
        c.JSON(400, gin.H{"error": "Invalid ID"})
        return
    }

    for _, pokemon := range Pokedex {
        if pokemon.Id == id {
            c.JSON(200, gin.H{"data": pokemon})
            return
        }
    }

    c.JSON(404, gin.H{"error": "Pokemon not found"})
}
