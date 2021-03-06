package main

import (
	"encoding/json"
	"github.com/Piszmog/pokeapi-client/cache"
	"github.com/Piszmog/pokeapi-client/client"
	"github.com/Piszmog/pokeapi-client/net"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

const (
	keyContentType       = "Content-Type"
	valueApplicationJson = "application/json"
	defaultCacheTTL      = 3600
)

type errorResponse struct {
	ErrorMessage string `json:"error_message"`
}

var apiClient *net.ApiClient
var pokemonCacheClient *cache.RedisClient

func main() {
	apiClient = net.CreateDefaultApiClient()
	pokemonCacheClient = cache.CreateLocalRedisClient("pokemon")
	defer pokemonCacheClient.Close()
	router := httprouter.New()
	router.GET("/pokemon", GetPokemon)
	log.Fatal(http.ListenAndServe(":8080", router))
}

func GetPokemon(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	writer.Header().Set(keyContentType, valueApplicationJson)
	query := request.URL.Query()
	identifier := query.Get("id")
	if len(identifier) == 0 {
		identifier = query.Get("name")
	}
	var pokemon client.Pokemon
	err := pokemonCacheClient.Get(identifier, &pokemon)
	if err != nil {
		writer.WriteHeader(500)
		errorResponse := errorResponse{
			ErrorMessage: err.Error(),
		}
		bytes, _ := json.Marshal(errorResponse)
		writer.Write(bytes)
		return
	}
	if pokemon.Id == 0 {
		serverPokemon, err := apiClient.GetPokemon(identifier)
		if err != nil {
			writer.WriteHeader(500)
			errorResponse := errorResponse{
				ErrorMessage: err.Error(),
			}
			bytes, _ := json.Marshal(errorResponse)
			writer.Write(bytes)
			return
		} else if serverPokemon.Id == 0 {
			writer.WriteHeader(404)
			errorResponse := errorResponse{
				ErrorMessage: "Failed to find pokemon " + identifier,
			}
			bytes, _ := json.Marshal(errorResponse)
			writer.Write(bytes)
			return
		}
		pokemonCacheClient.Insert(identifier, serverPokemon)
		pokemon = *serverPokemon
	}
	pokemonCacheClient.SetTTL(identifier, defaultCacheTTL)
	writer.WriteHeader(200)
	bytes, _ := json.Marshal(pokemon)
	writer.Write(bytes)
}
