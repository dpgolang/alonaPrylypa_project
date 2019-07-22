package controllers

import (
	//"database/sql"
	"encoding/gob"
	"encoding/json"
	"github.com/alonaprylypa/Project/pkg/db"
	"github.com/alonaprylypa/Project/pkg/models"
	"github.com/alonaprylypa/Project/pkg/repos"
	"github.com/gorilla/mux"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"html/template"
	"log"
	"net/http"
	//"net/smtp"
	"strconv"
)

type Controller struct {
	Finder db.DateFinder
}
var store *sessions.CookieStore
var tpl *template.Template

func init() {
	authKeyOne := securecookie.GenerateRandomKey(64)
	encryptionKeyOne := securecookie.GenerateRandomKey(32)
	store = sessions.NewCookieStore(authKeyOne, encryptionKeyOne)
	store.Options = &sessions.Options{
		MaxAge:   60 * 15,
		HttpOnly: true,
	}
	gob.Register(models.User{})
	tpl = template.Must(template.ParseGlob("ui/*.html"))
}
func NewControllers(finder db.DateFinder) *Controller {
	return &Controller{Finder: finder}
}
func (c *Controller) GetAllHousing(w http.ResponseWriter, r *http.Request) {
	apartments, err := c.Finder.GetAllApartments()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("while getting apartments got an error: %v", err)
		return
	}
	json.NewEncoder(w).Encode(apartments)
}
//func (c *Controller) SendMail(w http.ResponseWriter, r *http.Request) {
//	session, err := store.Get(r, "cookie-name")
//	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth || err != nil {
//		if err != nil {
//			http.Error(w, err.Error(), http.StatusInternalServerError)
//			return
//		}
//		http.Error(w, "You should sign in to check this page", http.StatusForbidden)
//		return
//	}


func (c *Controller) GetTypeHousing(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var (
		flats []models.Flat
		err   error
	)

	if params["type"] == "flats" {
		flats, err = c.Finder.GetFlats()
	} else if params["type"] == "houses" {
		flats, err = c.Finder.GetHouses()
	} else {
		w.WriteHeader(http.StatusNotFound)
		log.Println("Unexpected type")
		return
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("while getting apartments got an error:%v", err)
		return
	}

	json.NewEncoder(w).Encode(flats)
}

func (c *Controller) GetOneHousing(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	val, err := strconv.Atoi(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("you had a bad request:%v", err)
		return
	}
	fl, err := c.Finder.GetApartmentById(val)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("while getting apartments got an error:%v", err)
		return
	}
	json.NewEncoder(w).Encode(fl)
}
func (c *Controller) GetRealtor(w http.ResponseWriter, r * http.Request){
	session, err := store.Get(r, "cookie-name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	user := repos.GetUser(session)
	if !user.Authenticated{
		log.Printf("user should sign in to check this page:%v", err)
		http.Redirect(w, r, "/login", http.StatusMultipleChoices)
		return
	}
	params:=mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Printf("you had a bad request:%v", err)
		return
	}
	realtor,err:=c.Finder.GetRealtorDate(id)
}

