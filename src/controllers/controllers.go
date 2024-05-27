package controllers

import (
	"strconv"

	"github.com/abroudoux/pokemon-battle-simulator/src/models"
	"github.com/gin-gonic/gin"
)

var Pokedex []models.Pokemon
var Counter int

func InitPokedex() {
	Counter = 1

	pokemon1 := models.Pokemon{
		Id: 1,
		Name: "Bulbasaur",
		Type: []string{"Grass", "Poison"},
		Base: []int{45, 49, 49, 65, 65, 45},
	}

	Pokedex = append(Pokedex, pokemon1)
}

func GetPokedex(c * gin.Context) {
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

