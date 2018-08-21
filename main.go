package main

import (
	"fmt"
	"github.com/Piszmog/pokeapi-client/net"
	"io/ioutil"
)

const urlBase = "http://pokeapi.co/api/v2/"

func main() {
	httpClient := net.CreateDefaultHttpClient()
	resp, _ := httpClient.Get(urlBase + "pokemon/1/")
	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("%s", string(bodyBytes))
}
