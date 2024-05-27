package controllers

import (
	"strconv"

	"github.com/abroudoux/pokemon-battle-simulator/src/models"
	"github.com/gin-gonic/gin"
)

var Pokedex []models.Pokemon

func InitPokedex() {
	pokemon1 := models.Pokemon{
		Id: 1,
		Name: {
			English: "Bulbasaur",
			Japanese: "フシギダネ",
			Chinese: "妙蛙种子",
			French: "Bulbizarre",
		},
		Type: {"Grass", "Poison"},
		Base: {
			HP: 45,
			Attack: 49,
			Defense: 49,
			SpAttack: 65,
			SpDefense: 65,
			Speed: 45,
		},
	}

	Pokedex = append(Pokedex, pokemon1)
}

// func InitPokedex() {
// 	file, err := os.Open("data/pokedex.json")

// 	if err != nil {
// 		fmt.Println(err, "Error opening file")
// 		return
// 	}

// 	defer file.Close()

// 	var pokedex []models.Pokemon
// 	err = json.NewDecoder(file).Decode(&pokedex)

// 	if err != nil {	
// 		fmt.Println(err)
// 		return
// 	}

// 	for _, pokemon := range pokedex {
// 		fmt.Print("Pokemon: ", pokemon)
// 	}

// 	Pokedex = pokedex
// }

func GetPokedex(c *gin.Context) {
c.JSON(200, gin.H{"pokedex": Pokedex})
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

