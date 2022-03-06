package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/AmineNeifer/tournage-paris/models"
	"github.com/gocql/gocql"
)

func getHTTPResponse(url string) models.Response {
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

func activateModel() {

	url := "https://opendata.paris.fr/api/records/1.0/search/?dataset=lieux-de-tournage-a-paris&q=&rows=10000"
	host := "127.0.0.1"

	// get data
	resp := getHTTPResponse(url)

	records := resp.Records

	cluster := gocql.NewCluster(host)
	cluster.Consistency = gocql.Quorum
	cluster.Timeout = 3 * time.Second // 3000ms

	session, err := cluster.CreateSession()
	errorHandle(err)
	defer session.Close()

	createKeyspace("paris_open_data", session)

	createTable("filming", "paris_open_data", session)

	storeRecords("paris_open_data", "filming", session, records)

}

func createKeyspace(keyspace string, session *gocql.Session) {
	var query string = fmt.Sprintf("CREATE KEYSPACE IF NOT EXISTS %s"+
	" WITH REPLICATION = {'class' : 'SimpleStrategy', 'replication_factor' : 2}", keyspace)
	err := session.Query(query).Exec()
	errorHandle(err)
}

func createTable(table string, keyspace string, session *gocql.Session) {
	var query string = fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s.%s"+ 
	" (id timeuuid PRIMARY KEY, name text, producer text, director text,"+
	" coord_x double, coord_y double, type text, place_ardt int, place_id text,"+
	" address text, year int)", keyspace, table)
	// if you encounter an error here, try encreasing write_request_timeout_in_ms in /etc/cassandra/cassandra.yaml
	// or checkout https://stackoverflow.com/questions/42922757/getting-timeout-error-with-gocql
	err := session.Query(query).Exec()
	errorHandle(err)
}

func storeRecords(keyspace string, table string, session *gocql.Session, records []struct {
	Tournage models.Fields "json:\"fields\""
}) {
	for _, record := range records {
		r := record.Tournage
		query := fmt.Sprintf("INSERT INTO %s.%s"+
			" (id,name,producer,director,coord_x,coord_y,type,place_ardt,place_id,address,year)"+
			" VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", keyspace, table)
		err := session.Query(query, gocql.TimeUUID(), r.NomTournage, r.NomProducteur, r.NomRealisateur, r.CoordX, r.CoordY, r.TypeTournage, r.ArdtLieu, r.IDLieu, badSpaceHandle(r.AdresseLieu), r.AnneeTournage).Exec()
		errorHandle(err)
	}
}
