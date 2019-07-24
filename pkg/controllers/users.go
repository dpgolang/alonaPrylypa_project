//This file contains handlers for
//working with users
package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/alonaprylypa/Project/pkg/models"
	"github.com/alonaprylypa/Project/pkg/repos"
	"golang.org/x/crypto/bcrypt"
)

//LoginPageHandler loads html page to login users
func (c *Controller) LoginPageHandler(w http.ResponseWriter, r *http.Request) {
	var body, err = repos.LoadFile("ui/login.html")
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		log.Printf("http page is not available: %v", err)
		return
	}
	fmt.Fprintf(w, body)
}

//LoginHandler collect data from forms and check with users who are registered
//and entered into the database
func (c *Controller) LoginHandler(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "cookie-name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
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
		log.Printf("user doesn't exists:%v", err)
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

//IndexPageHandler loads users html page
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

//LogOut cleans the current user session and redirects into the home page
func (c *Controller) LogOut(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "cookie-name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	session.Values["user"] = models.User{}
	session.Options.MaxAge = -1
	err = session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/", http.StatusFound)
}

//RegisterPageHandler loads html page to registered users
func (c *Controller) RegisterPageHandler(w http.ResponseWriter, r *http.Request) {
	var body, err = repos.LoadFile("ui/register.html")
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		log.Printf("http page is not available: %v", err)
		return
	}
	fmt.Fprintf(w, body)
}

//RegisterHandler checks that password and confirm password are matched.
//And if everything is good, it insertsthe user in the database table
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
		http.Redirect(w, r, "/register", http.StatusPermanentRedirect)
		return
	}
	http.Redirect(w, r, "/login", http.StatusFound)
}
