package repository

import (
	"github.com/groupecrit/agences/metier"
)

var agences = map[string]string{
	"431": "CRIT SELESTAT",
	"433": "CRIT COLMAR Généraliste",
	"434": "CRIT COLMAR Tertiaire -Industrie - TL",
	"437": "CRIT DIDENHEIM Automobile - Industrie",
	"438": "CRIT DIDENHEIM Commerce Agro Evene",
	"440": "CRIT ALTKIRCH",
	"442": "CRIT MULHOUSE Industrie",
	"443": "CRIT MULHOUSE BTP",
	"444": "CRIT SAINT LOUIS",
	"445": "CRIT GUEBWILLER Generaliste",
}

type Agences struct{}

func NewAgences() Agences {
	return Agences{}
}

func (a Agences) Get(code string) metier.Agence {
	nom := agences[code]
	return metier.Agence{Code: code, Nom: nom}
}
