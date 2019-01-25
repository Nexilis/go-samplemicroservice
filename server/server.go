package server

import (
	"net/http"
	"time"
)

func configureTls() {
	// TODO: https://blog.cloudflare.com/exposing-go-on-the-internet/
}

func New(mux *http.ServeMux, serverAddress string) (srv *http.Server){
	configureTls()
	srv = &http.Server{
		Addr:         serverAddress,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      mux,
	}
	return
}
