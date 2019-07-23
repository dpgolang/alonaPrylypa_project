package controllers

import (
	"fmt"
	"github.com/alonaprylypa/Project/pkg/models"
	"github.com/alonaprylypa/Project/pkg/repos"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

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
	if err != nil {
		log.Printf("user doesn't login:%v", err)
		http.Redirect(w, r, r.Header.Get("Referer"), 302)
		return
	}
	name := r.FormValue("name")
	pass := r.FormValue("password")
	customer, err := c.Finder.ReturnCustomer(name)
	if err != nil {

		log.Printf("user doesn't exists:%v", err)
		http.Redirect(w, r, r.Header.Get("Referer"), 302)
		return
	}
	if err = bcrypt.CompareHashAndPassword([]byte(customer.Password), []byte(pass)); err != nil {
		log.Printf("password is incorrect:%v", err)
		http.Redirect(w, r, r.Header.Get("Referer"), 302)
		return
	}
	user := &models.User{
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
	user := repos.GetUser(session)
	fmt.Fprintf(w, body, user.Username)
}
func (c *Controller) LogOut(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "cookie-name")

	session.Values["user"] = models.User{}
	session.Options.MaxAge = -1
	err = session.Save(r, w)
	if err != nil {
		w.WriteHeader(http.StatusPermanentRedirect)
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
