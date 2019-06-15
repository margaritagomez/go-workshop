package main

import (
	"fmt"
)

var pokemons []Pokemon

func init() {
	pokemons = []Pokemon{
		Pokemon{
			PokedexNumber: 1,
			Name:          "Bulbasaur",
			EvolvesTo:     "Ivysaur",
			Types: []Type{
				TypeGrass,
				TypePoison,
			},
		},
	}
}

func main() {
	fmt.Println("Let's get started!")
}
