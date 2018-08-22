package main

import (
	"encoding/json"
	"github.com/Piszmog/pokeapi-client/net"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

const (
	keyContentType       = "Content-Type"
	valueApplicationJson = "application/json"
)

var client *net.ApiClient

type errorResponse struct {
	Path         string `json:"path"`
	ErrorMessage string `json:"error_message"`
}

func main() {
	client = net.CreateDefaultApiClient()
	router := httprouter.New()
	router.GET("/pokemon", GetPokemon)
	log.Fatal(http.ListenAndServe(":8080", router))
}

func GetPokemon(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	query := request.URL.Query()
	identifier := query.Get("id")
	if len(identifier) == 0 {
		identifier = query.Get("name")
	}
	pokemon, err := client.GetPokemon(identifier)
	writer.Header().Set(keyContentType, valueApplicationJson)
	if err != nil {
		writer.WriteHeader(500)
		errorResponse := errorResponse{
			Path:         "/pokemon/" + identifier,
			ErrorMessage: err.Error(),
		}
		bytes, _ := json.Marshal(errorResponse)
		writer.Write(bytes)
	} else if pokemon.Id == 0 {
		writer.WriteHeader(404)
		errorResponse := errorResponse{
			Path:         "/pokemon/" + identifier,
			ErrorMessage: "Failed to find pokemon " + identifier,
		}
		bytes, _ := json.Marshal(errorResponse)
		writer.Write(bytes)
	} else {
		writer.WriteHeader(200)
		bytes, _ := json.Marshal(pokemon)
		writer.Write(bytes)
	}
}
