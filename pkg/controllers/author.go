package controllers

//import (
//	"Project/pkg/models"
//	"Project/pkg/repos"
//	"database/sql"
//	"golang.org/x/crypto/bcrypt"
//
//	"fmt"
//	"net/http"
//)
//
//func (c *Controller) LoginPageHandler(db *sql.DB) http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		var body, err = repos.LoadFile("ui/login.html")
//		if err != nil {
//			logFatal(err)
//		}
//		fmt.Fprintf(w, body)
//	}
//}
//func (c *Controller) LoginHandler(db *sql.DB) http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		name := r.FormValue("name")
//		pass := r.FormValue("password")
//		result := db.QueryRow("select password from users where username=$1", name)
//		storedCreds := &models.Customer{}
//		err := result.Scan(&storedCreds.Password)
//		if err != nil {
//			if err == sql.ErrNoRows {
//				logFatal(err)
//			}
//			logFatal(err)
//		}
//		if err = bcrypt.CompareHashAndPassword([]byte(storedCreds.Password), []byte(pass)); err != nil {
//			logFatal(err)
//		}
//	}
//}
//func (c *Controller) RegisterPageHandler(db *sql.DB) http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		var body, err = repos.LoadFile("ui/register.html")
//		if err != nil {
//			w.WriteHeader(http.StatusBadRequest)
//			return
//		}
//		fmt.Fprintf(w, body)
//	}
//}
//func (c *Controller) RegisterHandler(db *sql.DB) http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		r.ParseForm()
//		customer := models.Customer{r.FormValue("username"), r.FormValue("email"), r.FormValue("password")}
//		if r.FormValue("confirmPassword") != customer.Password {
//			fmt.Fprintln(w, "Enter a correct password!")
//			return
//		}
//		http.Redirect(w,r,"/",301)
//		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(customer.Password), 8)
//		if _, err = db.Query("insert into users values ($1,$2,$3)", customer.UserName, string(hashedPassword), customer.Email); err != nil {
//			logFatal(err)
//		}
//	}
//
//}
//func (c *Controller) GetUserPage(db *sql.DB) http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		r.ParseForm()
//		customer := models.Customer{r.FormValue("username"), r.FormValue("email"), r.FormValue("password")}
//		if r.FormValue("confirmPassword") != customer.Password {
//			fmt.Fprintln(w, "Enter a correct password!")
//		}
//		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(customer.Password), 8)
//		if _, err = db.Query("insert into users values ($1,$2,$3)", customer.UserName, string(hashedPassword), customer.Email); err != nil {
//			logFatal(err)
//		}
//	}
//}
