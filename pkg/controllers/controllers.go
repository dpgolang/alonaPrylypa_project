package controllers

import (
	"Project/pkg/models"
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Controller struct {
}

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
func (c *Controller) GetAllHousing(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query("SELECT * FROM living_spaces")
		if err != nil {
			logFatal(err)
		}
		defer rows.Close()
		flats := make([]*models.Flat, 0)
		for rows.Next() {
			fl := new(models.Flat)
			err := rows.Scan(&fl.Id, &fl.Type, &fl.Street, &fl.Price, &fl.Square, &fl.Rooms, &fl.Floor)
			if err != nil {
				logFatal(err)
			}
			flats = append(flats, fl)
		}
		if err = rows.Err(); err != nil {
			logFatal(err)
		}
		json.NewEncoder(w).Encode(flats)
	}
}
func (c *Controller) GetTypeHousing(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		que:=""
		if params["type"]=="flats"{
			que="SELECT * FROM living_spaces WHERE type='Flat'"
		}
		if params["type"]=="houses"{
			que="SELECT * FROM living_spaces WHERE type='House'"
		}
		rows, err := db.Query(que)
		if err != nil {
			logFatal(err)
		}
		defer rows.Close()
		flats := make([]*models.Flat, 0)
		for rows.Next() {
			fl := new(models.Flat)
			err := rows.Scan(&fl.Id, &fl.Type, &fl.Street, &fl.Price, &fl.Square, &fl.Rooms, &fl.Floor)
			if err != nil {
				logFatal(err)
			}
			flats = append(flats, fl)
		}
		if err = rows.Err(); err != nil {
			logFatal(err)
		}
		json.NewEncoder(w).Encode(flats)
	}
}
func (c *Controller) GetFlat(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var fl models.Flat
		params := mux.Vars(r)
		row := db.QueryRow("select * from living_spaces where id = $1", params["id"])
		err := row.Scan(&fl.Id, &fl.Type, &fl.Street, &fl.Price, &fl.Square, &fl.Rooms, &fl.Floor)
		if err != nil {
			logFatal(err)
		}
		json.NewEncoder(w).Encode(fl)
	}
}
