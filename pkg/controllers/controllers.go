package controllers

import (
	"database/sql"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"github.com/alonaprylypa/Project/pkg/db"
	"github.com/alonaprylypa/Project/pkg/models"
	"github.com/alonaprylypa/Project/pkg/repos"
	"github.com/gorilla/mux"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"golang.org/x/crypto/bcrypt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

type Controller struct {
	Finder db.ApartmentsFinder
}
type User struct {
	Username      string
	Authenticated bool
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
	gob.Register(User{})
	tpl = template.Must(template.ParseGlob("ui/*.html"))
}
func NewControllers(finder db.ApartmentsFinder) *Controller {
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
func GetUser(s *sessions.Session) User {
	val := s.Values["user"]
	var user = User{}
	user, ok := val.(User)
	if !ok {
		return User{Authenticated: false}
	}
	return user
}

//func (c *Controller) SendMail(w http.ResponseWriter, r *http.Request){
//	session,err:=store.Get(r,"cookie-name")
//	if err != nil {
//		http.Error(w,err.Error(),http.StatusInternalServerError)
//		return
//	}
//	params:=mux.Vars(r)
//	val, err := strconv.Atoi(params["id"])
//	if err != nil {
//		w.WriteHeader(http.StatusBadRequest)
//		log.Printf("you had a bad request:%v", err)
//		return
//	}
//	msg, err := c.Finder.GetApartmentById(val)
//	if err != nil {
//		w.WriteHeader(http.StatusInternalServerError)
//		log.Printf("while getting apartments got an error:%v", err)
//		return
//	}
//	user:=GetUser(session)
//	email, err := c.Finder.GetEmail(user.Username)
//	if err != nil {
//		w.WriteHeader(http.StatusInternalServerError)
//		log.Printf("while getting email got an error:%v", err)
//		return
//	}
//	err=smtp.SendMail()
//
//}
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
func (c *Controller) LoginPageHandler(w http.ResponseWriter, r *http.Request) {
	var body, err = repos.LoadFile("ui/login.html")
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		log.Printf("http page is not available: %v", err)
		return
	}
	fmt.Fprintf(w, body)
}
func (c *Controller) LoginHandler(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "cookie-name")

	name := r.FormValue("name")
	pass := r.FormValue("password")
	customer, err := c.Finder.ReturnCustomer(name)
	if err != nil {
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusUnauthorized)
			log.Printf("user doesn't exists:%v", err)
			http.Redirect(w, r, "/register", http.StatusMultipleChoices)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("user is not found:%v", err)
		return
	}
	if err = bcrypt.CompareHashAndPassword([]byte(customer.Password), []byte(pass)); err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		log.Printf("password is incorrect:%v", err)
		http.Redirect(w, r, "/login", http.StatusMultipleChoices)
		return
	}
	user := &User{
		Username:      name,
		Authenticated: true,
	}
	session.Values["user"] = user
	err = session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/index", http.StatusFound)

}
func (c *Controller) IndexPageHandler(w http.ResponseWriter, r *http.Request) {
	var body, err = repos.LoadFile("ui/index.html")
	session, err := store.Get(r, "cookie-name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	user := GetUser(session)
	fmt.Fprintf(w, body, user.Username)
}
func (c *Controller) LogOut(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "cookie-name")

	session.Values["user"] = User{}
	session.Options.MaxAge = -1
	err = session.Save(r, w)
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		log.Printf("http page is not available: %v", err)
		return
	}
	http.Redirect(w, r, "/", http.StatusFound)
}
func (c *Controller) RegisterPageHandler(w http.ResponseWriter, r *http.Request) {
	var body, err = repos.LoadFile("ui/register.html")
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		log.Printf("http page is not available: %v", err)
		return
	}
	fmt.Fprintf(w, body)
}
func (c *Controller) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	customer := models.Customer{r.FormValue("username"), r.FormValue("email"), r.FormValue("password")}
	if r.FormValue("confirmPassword") != customer.Password {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Print(w, "repeat the password correctly")
		http.Redirect(w, r, "/register", http.StatusPermanentRedirect)
		return
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(customer.Password), 8)
	if err != nil {
		log.Println("impossible to hash the password")
		return
	}
	err = c.Finder.RegisterCustomer(customer.UserName, customer.Email, string(hashedPassword))
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		http.Redirect(w, r, "/register", http.StatusMultipleChoices)
		return
	}
	http.Redirect(w, r, "/login", http.StatusFound)
}
