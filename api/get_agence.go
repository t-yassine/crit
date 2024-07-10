package api

import (
	"encoding/json"
	"github.com/groupecrit/agences/metier"
	"net/http"
)

type Service interface {
	GetAgence(code string) metier.Agence
}

var service Service

func Init(s Service) {
	service = s
}

func HandleGetAgence(rw http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		rw.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	code := r.URL.Query().Get("code")
	if code == "" {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("code is required"))
		return
	}

	agence := service.GetAgence(code)
	if agence.Nom == "" {
		rw.WriteHeader(http.StatusNotFound)
		rw.Write([]byte("agency not found"))
		return
	}

	agenceJSON, err := json.Marshal(agence)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	rw.Write(agenceJSON)
}
