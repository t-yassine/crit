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

	repo := repository.NewAgences()
	service := metier.NewService(repo)
	api.Init(service)

	http.HandleFunc("/healthcheck", api.HandleHealthCheck)
	http.HandleFunc("/agence", api.HandleGetAgence)

	log.Printf("le serveur HTTP démarre sur le port %s\n", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil); err != nil {
		log.Fatalf("le serveur HTTP s'est arrêté inopinément : %v", err)
	}
}
