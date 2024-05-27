package main

import (
	controllers "github.com/abroudoux/pokemon-battle-simulator/src/controllers"
	router "github.com/abroudoux/pokemon-battle-simulator/src/routes"
)

func main () {
	controllers.InitPokedex()
	router.Router()
}