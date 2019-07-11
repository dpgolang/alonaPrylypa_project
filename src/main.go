package main

import (
	"./github.com/gorilla/mux"
	_ "./github.com/lib/pq"
	"database/sql"
	"fmt"
	"log"
	"net/http"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "190117"
	dbname   = "postgres"
)

type Flat struct {
	Id     int
	Type   string
	Street string
	Price  int
	Square float64
	Rooms  int
	Floor  int
}

var db *sql.DB

func init() {
	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/flats", getFlats).Methods("GET")
	router.HandleFunc("/flat/{id:[0-9]+}", getFlat).Methods("GET")
	router.HandleFunc("/flats/add", addFlat).Methods("POST")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func getFlats(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT * FROM living_spaces")
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	defer rows.Close()
	flats := make([]*Flat, 0)
	for rows.Next() {
		fl := new(Flat)
		err := rows.Scan(&fl.Id, &fl.Type, &fl.Street, &fl.Price, &fl.Square, &fl.Rooms, &fl.Floor)
		if err != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}
		flats = append(flats, fl)
	}
	if err = rows.Err(); err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	for _, fl := range flats {
		fmt.Fprintf(w, "%s,%s,%s,%s,%s,%s,%s", fl.Id, fl.Type, fl.Street, fl.Price, fl.Square, fl.Rooms, fl.Floor)
	}

}

func getFlat(w http.ResponseWriter, r *http.Request) {

}

func addFlat(w http.ResponseWriter, r *http.Request) {

}
