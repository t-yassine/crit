package api

import (
	"github.com/groupecrit/agences/metier"
	"github.com/stretchr/testify/require"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type mockService struct {
	getAgence func(code string) metier.Agence
}

func (m mockService) GetAgence(code string) metier.Agence {
	return m.getAgence(code)
}

func TestHandleGetAgence(t *testing.T) {
	service := mockService{
		getAgence: func(code string) metier.Agence {
			if code == "431" {
				return metier.Agence{Code: "431", Nom: "CRIT SELESTAT"}
			}
			return metier.Agence{}
		},
	}
	Init(service)

	tests := []struct {
		name       string
		code       string
		wantStatus int
		wantBody   string
	}{
		{"Valid Code", "431", http.StatusOK, `{"Code":"431","Nom":"CRIT SELESTAT"}`},
		{"Invalid Code", "999", http.StatusNotFound, "agency not found"},
		{"Missing Code", "", http.StatusBadRequest, "code is required"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/agency?code="+tt.code, nil)
			rec := httptest.NewRecorder()

			HandleGetAgence(rec, req)

			res := rec.Result()
			defer res.Body.Close()

			// Assert the status code
			require.Equal(t, tt.wantStatus, res.StatusCode)

			// Read the body
			body, err := io.ReadAll(res.Body)
			require.NoError(t, err)

			// Assert the response body
			if tt.wantStatus == http.StatusOK {
				require.JSONEq(t, tt.wantBody, string(body))
			} else {
				require.Equal(t, tt.wantBody, strings.TrimSpace(string(body)))
			}
		})
	}
}
