package models

type Response struct {
	Records []struct {
		Tournage struct {
			NomTournage    string  `json:"nom_tournage"`
			NomProducteur  string  `json:"nom_producteur"`
			NomRealisateur string  `json:"nom_realisateur"`
			CoordX         float64 `json:"coord_x"`
			CoordY         float64 `json:"coord_y"`
			TypeTournage   string  `json:"type_tournage"`
			ArdtLieu       string  `json:"ardt_lieu"`
			IDLieu         string  `json:"id_lieu"`
			AdresseLieu    string  `json:"adresse_lieu"`
			AnneeTournage  string  `json:"annee_tournage"`
		} `json:"fields"`
	}
}
