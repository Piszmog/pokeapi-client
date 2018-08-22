package net

import (
	"encoding/json"
	"github.com/Piszmog/pokeapi-client/client"
	"github.com/pkg/errors"
	"net"
	"net/http"
	"time"
)

const (
	urlBase        = "http://pokeapi.co/api/v2/"
	pokemonUrlPath = "pokemon/"
)

type ApiClient struct {
	baseUrl    string
	httpClient *http.Client
}

func (apiClient ApiClient) GetPokemon(identifier string) (*client.Pokemon, error) {
	resp, err := apiClient.httpClient.Get(apiClient.baseUrl + pokemonUrlPath + identifier)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to retrieve details on pokemon %d", identifier)
	}
	defer resp.Body.Close()
	pokemon := &client.Pokemon{}
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(pokemon)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to decode response for pokemon %d", identifier)
	}
	return pokemon, nil
}

// CreateDefaultApiClient creates a default http.Client.
//
// Timeout set to 5 seconds, keep alive set to 30 seconds, TLS handshake timeout set to 5 seconds, and idleConnection set to
// 90 seconds.
func CreateDefaultApiClient() *ApiClient {
	return CreateApiClient(5*time.Second, 30*time.Second, 5*time.Second, 90*time.Second)
}

// CreateApiClient creates a http.Client from the specified timeouts and keep alive.
//
// The client also has the maximum number of idle connections set to 100 and number of connections per host as 100.
func CreateApiClient(timeout time.Duration, keepAlive time.Duration, tlsHandshakeTimeout time.Duration, idleConnection time.Duration) *ApiClient {
	transport := &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   timeout,
			KeepAlive: keepAlive,
			DualStack: true,
		}).DialContext,
		TLSHandshakeTimeout: tlsHandshakeTimeout,
		IdleConnTimeout:     idleConnection,
		MaxIdleConns:        100,
		MaxIdleConnsPerHost: 100,
	}
	httpClient := &http.Client{
		Transport: transport,
		Timeout:   timeout,
	}

	return &ApiClient{
		baseUrl:    urlBase,
		httpClient: httpClient,
	}
}
