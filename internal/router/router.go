package router

import (
	pokemon "github.com/abroudoux/pokemon-battle-simulator/internal/pokemon"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
    router := gin.Default()

    router.GET("/pokedex", pokemon.GetPokedex)
    router.GET("/pokedex/:pokemon", pokemon.GetPokemon)

	router.GET("/battle/:pokemon1/:pokemon2", pokemon.CreateBattle)

    router.Run(":151")
}
