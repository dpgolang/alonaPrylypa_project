package main

import (
	"github.com/alonaprylypa/Project/pkg/controllers"
	"github.com/alonaprylypa/Project/pkg/db"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Printf("No .env file found:%v", err)
	}
}
func main() {
	port := os.Getenv("SERVICE_PORT")
	if len(port) == 0 {
		port = "8000"
	}
	finder := db.NewAppartmentsStorage()
	router := mux.NewRouter()
	controller := controllers.NewControllers(finder)

	router.HandleFunc("/housing", controller.GetAllHousing).Methods(http.MethodGet)
	router.HandleFunc("/housing/{type:flats|houses}", controller.GetTypeHousing).Methods(http.MethodGet)
	router.HandleFunc("/housing/{id:[0-9]+}", controller.GetOneHousing).Methods(http.MethodGet)
	//router.HandleFunc("/{id:[0-9]+}", controller.SendMail).Methods(http.MethodPost)
	router.HandleFunc("/login", controller.LoginPageHandler).Methods(http.MethodGet)
	router.HandleFunc("/login", controller.LoginHandler).Methods(http.MethodPost)
	router.HandleFunc("/index", controller.IndexPageHandler).Methods(http.MethodGet)
	router.HandleFunc("/index", controller.LogOut).Methods(http.MethodPost)
	router.HandleFunc("/register", controller.RegisterPageHandler).Methods(http.MethodGet)
	router.HandleFunc("/register", controller.RegisterHandler).Methods(http.MethodPost)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
