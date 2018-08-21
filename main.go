package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	client := &http.Client{}
	resp, _ := client.Get("http://pokeapi.co/api/v2/pokemon/1/")
	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("%s", string(bodyBytes))
}
