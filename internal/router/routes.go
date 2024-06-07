package router

import (
	"net/http"

	"github.com/abroudoux/pokemon-battle-simulator/internal/controllers"
	"github.com/gin-gonic/gin"
)

func Router() {
    router := gin.Default()

    router.GET("/", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "message": "Hello World!",
        })
    })

    router.GET("/pokedex", controllers.GetPokedex)
    router.GET("/pokedex/:pokemon", controllers.GetPokemon)

	router.GET("/battle/:pokemon1/:pokemon2", controllers.CreateBattle)

    // router.GET("/test/:s", controllers.Test)

    router.Run(":8080")
}
