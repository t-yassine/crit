package metier

import (
	"testing"
)

func TestService_GetAgence(t *testing.T) {
	t.Parallel()

	s := NewService(mockRepository{
		get: func(code string) Agence {
			return Agence{
				Code: code,
				Nom:  "nomAgence",
			}
		},
	})

	agence := s.GetAgence("431")

	if agence.Nom != "nomAgence" {
		t.Errorf("le nom de l'agence attendu est %q, obtenu %q", "nomAgence", agence.Nom)
	}
}

type mockRepository struct {
	get func(code string) Agence
}

func (r mockRepository) Get(code string) Agence {
	return r.get(code)
}
