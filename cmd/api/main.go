package main

import (
	pokemon "github.com/abroudoux/pokemon-battle-simulator/internal/pokemon"
	router "github.com/abroudoux/pokemon-battle-simulator/internal/router"
)

func main () {
	pokemon.InitPokedex()
	router.Router()
}