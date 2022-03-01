package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/AmineNeifer/tournage-paris/models"
)

func GetHTTPResponse(url string) models.Response {
	get, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer get.Body.Close()
	body, err := io.ReadAll(get.Body)
	if err != nil {
		log.Fatal(err)
	}
	resp := models.Response{}
	if err := json.Unmarshal(body, &resp); err != nil {
		log.Fatal(err)
	}
	return resp
}
