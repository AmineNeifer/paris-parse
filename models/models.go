package models

import "github.com/gocql/gocql"

type Response struct {
	Records []struct {
		Tournage Fields `json:"fields"`
	}
}

type Fields struct {
	ID             gocql.UUID `json:"id" cql:"id"`
	NomTournage    string     `json:"nom_tournage" cql:"name"`
	NomProducteur  string     `json:"nom_producteur" cql:"producer"`
	NomRealisateur string     `json:"nom_realisateur" cql:"director"`
	CoordX         float64    `json:"coord_x" cql:"coord_x"`
	CoordY         float64    `json:"coord_y" cql:"coord_y"`
	TypeTournage   string     `json:"type_tournage" cql:"type"`
	ArdtLieu       int        `json:"ardt_lieu,string" cql:"place_ardt"`
	IDLieu         string     `json:"id_lieu" cql:"place_id"`
	AdresseLieu    string     `json:"adresse_lieu" cql:"address"`
	AnneeTournage  int        `json:"annee_tournage,string" cql:"year"`
}