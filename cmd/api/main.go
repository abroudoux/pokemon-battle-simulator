package main

import (
	controllers "github.com/abroudoux/pokemon-battle-simulator/internal/controllers"
	router "github.com/abroudoux/pokemon-battle-simulator/internal/routes"
)

func main () {
	controllers.InitPokedex()
	router.Router()
}