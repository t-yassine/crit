package main

import (
	"flag"
	"fmt"
	"github.com/groupecrit/agences/metier"
	"github.com/groupecrit/agences/repository"
	"log"
	"net/http"

	"github.com/groupecrit/agences/api"
)

func main() {
	var port string
	flag.StringVar(&port, "port", "8080", "Le port utilisé par le serveur HTTP")
	flag.Parse()

	// Initialize the repository that provides access to agency data.
	// This repository interacts with the data source to fetch agency details.
	repo := repository.NewAgences()
	// Create a new service instance, passing in the repository.
	// The service contains business logic and uses the repository to get data.
	service := metier.NewService(repo)
	// Initialize the API package with the service.
	// This makes the service available to the API handlers.
	api.Init(service)

	http.HandleFunc("/healthcheck", api.HandleHealthCheck)

	// The /agency endpoint is used to get information about a specific agency.
	// It expects a query parameter 'code' which specifies the agency code.
	http.HandleFunc("/agency", api.HandleGetAgence)

	log.Printf("le serveur HTTP démarre sur le port %s\n", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil); err != nil {
		log.Fatalf("le serveur HTTP s'est arrêté inopinément : %v", err)
	}
}
