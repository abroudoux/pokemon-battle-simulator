package router

import (
	"net/http"

	"github.com/abroudoux/pokemon-battle-simulator/src/controllers"
	"github.com/gin-gonic/gin"
)

func Router() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello Worouterld",
		})
	})
	router.GET("/pokedex", controllers.GetPokedex)
	router.GET("/pokemon/:id", controllers.GetPokemon)

	router.Run(":8080")
}

