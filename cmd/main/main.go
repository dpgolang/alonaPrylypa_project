package main

import (
	"Project/pkg/controllers"
	"Project/pkg/driver"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"github.com/joho/godotenv"
	"os"
)
func init(){
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}
func main() {
	port:= os.Getenv("SERVICE_PORT")
	if len(port)==0 {
		log.Fatal("Required parameter service port is not set")
	}
	db := driver.ConnectDB()
	defer db.Close()
	router := mux.NewRouter()
	controller := controllers.Controller{}

	router.HandleFunc("/", controller.GetAllHousing(db)).Methods("GET")
	router.HandleFunc("/{type:flats|houses}", controller.GetTypeHousing(db)).Methods("GET")
	router.HandleFunc("/{id:[0-9]+}", controller.GetFlat(db)).Methods("GET")
	//router.HandleFunc("/login", controller.LoginPageHandler(db)).Methods("GET")
	//router.HandleFunc("/login", controller.LoginHandler(db)).Methods("POST")
	//router.HandleFunc("/register", controller.RegisterPageHandler(db)).Methods("GET")
	//router.HandleFunc("/register", controller.RegisterHandler(db)).Methods("POST")
	//router.HandleFunc("/index/{name}",controller.GetUserPage(db)).Methods("GET")

	log.Fatal(http.ListenAndServe(":"+port, router))
}
