package controllers

import (
	"Project/pkg/db"
	"Project/pkg/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type Controller struct {
	Finder db.ApartmentsFinder
}
func NewControllers(finder db.ApartmentsFinder)*Controller{
	return &Controller{Finder:finder}
}
func (c *Controller) GetAllHousing(w http.ResponseWriter, r *http.Request){
	apartments,err:=c.Finder.GetAllApartments()
	if err!=nil{
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("while getting apartments got an error: %v",err)
		return
	}
	json.NewEncoder(w).Encode(apartments)
}
func (c *Controller) GetTypeHousing(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		var(
			flats []models.Flat
			err error
		)

		if params["type"]=="flats"{
			flats,err=c.Finder.GetFlats()
		}else if params["type"]=="houses"{
			flats,err=c.Finder.GetHouses()
		}else{
			w.WriteHeader(http.StatusNotFound)
			log.Println("Unexpected type")
			return
		}
		if err!=nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Printf("while getting apartments got an error:%v",err)
			return
		}

		json.NewEncoder(w).Encode(flats)
	}

func (c *Controller) GetOneHousing(w http.ResponseWriter, r *http.Request){
		params := mux.Vars(r)
		val,err:=strconv.Atoi(params["id"])
		if err!=nil{
			w.WriteHeader(http.StatusBadRequest)
			log.Printf("you had a bad request:%v",err)
			return
		}
		fl,err:=c.Finder.GetApartmentById(val)
		if err!=nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Printf("while getting apartments got an error:%v",err)
			return
		}
		json.NewEncoder(w).Encode(fl)
	}

