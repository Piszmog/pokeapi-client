package net

import (
	"net"
	"net/http"
	"time"
)

// CreateDefaultHttpClient creates a default http.Client.
//
// Timeout set to 5 seconds, keep alive set to 30 seconds, TLS handshake timeout set to 5 seconds, and idleConnection set to
// 90 seconds.
func CreateDefaultHttpClient() *http.Client {
	return CreateHttpClient(5*time.Second, 30*time.Second, 5*time.Second, 90*time.Second)
}

// CreateHttpClient creates a http.Client from the specified timeouts and keep alive.
//
// The client also has the maximum number of idle connections set to 100 and number of connections per host as 100.
func CreateHttpClient(timeout time.Duration, keepAlive time.Duration, tlsHandshakeTimeout time.Duration, idleConnection time.Duration) *http.Client {
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
	return &http.Client{
		Transport: transport,
		Timeout:   timeout,
	}
}
