package main

import (
	"fmt"
	"github.com/Piszmog/pokeapi-client/net"
)

func main() {
	client := net.CreateDefaultHttpClient()
	pokemon, err := client.GetPokemonById(1)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", pokemon)
}
