package main

import (
	"fmt"
	"github.com/Piszmog/pokeapi-client/net"
	"io/ioutil"
)

func main() {
	httpClient := net.CreateDefaultHttpClient()
	resp, _ := httpClient.Get("http://pokeapi.co/api/v2/pokemon/1/")
	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("%s", string(bodyBytes))
}
