package main

import (
	"Project/pkg/controllers"
	"Project/pkg/db"
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
	finder:=db.NewAppartmentsStorage()
	router:=mux.NewRouter()
	controller:=controllers.NewControllers(finder)

	router.HandleFunc("/", controller.GetAllHousing).Methods(http.MethodGet)
	router.HandleFunc("/{type:flats|houses}", controller.GetTypeHousing).Methods(http.MethodGet)
	router.HandleFunc("/{id:[0-9]+}", controller.GetOneHousing).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":"+port, router))
}
