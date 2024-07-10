package api

import (
	"encoding/json"
	"github.com/groupecrit/agences/metier"
	"net/http"
)

// Service defines the methods that any service providing
// agency information must implement.
type Service interface {
	// GetAgence returns the details of an agency identified by the given code.
	GetAgence(code string) metier.Agence
}

// service is a package-level variable that holds the implementation
// of the Service interface. This allows the handlers to access the
// service without having to pass it explicitly.
var service Service

// Init initializes the package-level service variable with the given
// implementation of the Service interface. This function must be called
// before any handlers that depend on the service are used.
func Init(s Service) {
	service = s
}

// HandleGetAgence handles HTTP GET requests for retrieving agency information.
// It expects a query parameter "code" which specifies the agency code.
// If the code is missing or the method is not GET, it responds with an
// appropriate HTTP status code. On success, it returns the agency details
// as a JSON object.
func HandleGetAgence(rw http.ResponseWriter, r *http.Request) {
	// Only allow GET requests
	if r.Method != http.MethodGet {
		rw.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// Get the agency code from the query parameters
	code := r.URL.Query().Get("code")
	if code == "" {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte("code is required"))
		return
	}

	// Retrieve the agency details using the service
	agence := service.GetAgence(code)
	if agence.Nom == "" {
		rw.WriteHeader(http.StatusNotFound)
		rw.Write([]byte("agency not found"))
		return
	}

	// Marshal the agency details into JSON
	agenceJSON, err := json.Marshal(agence)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Set the content type to JSON and write the response
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	rw.Write(agenceJSON)
}
