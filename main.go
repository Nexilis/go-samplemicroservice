package main

import (
	"log"
	"net/http"
	"os"
	"samplemicroservice/home"
	"samplemicroservice/server"
)

// based on https://youtu.be/bM6N-vgPlyQ
func main () {
	logger := log.New(os.Stdout, "my-microservice ", log.LstdFlags | log.Lshortfile)

	logger.Println("starting service...")

	mux := http.NewServeMux()
	configureRouting(mux, logger)

	srv := server.New(mux, ":8088")
	err := srv.ListenAndServe()
	if err != nil {
		logger.Fatalf("server failed to start: %v", err)
	}
}

func configureRouting(mux *http.ServeMux, logger *log.Logger) {
	h := home.NewHandlers(logger)

	mux.HandleFunc("/", h.Home)
}

