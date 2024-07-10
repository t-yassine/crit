package api

import (
	"encoding/json"
	"net/http"
)

func HandleHealthCheck(rw http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		rw.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	status := map[string]string{"status": "OK"}
	statusJSON, err := json.Marshal(status)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusOK)
	rw.Write(statusJSON)
}
