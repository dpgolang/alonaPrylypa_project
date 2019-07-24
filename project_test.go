package Project

import (
	"github.com/alonaprylypa/Project/pkg/controllers"
	"github.com/alonaprylypa/Project/pkg/db"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"testing"
)

func checkError(err error, t *testing.T) {
	if err != nil {
		t.Errorf("An error occurred. %v", err)
	}
}

var finder = db.NewAppartmentsStorage()
var router = mux.NewRouter()
var controller = controllers.NewControllers(finder)

func TestGetAllHousing(t *testing.T) {
	req, err := http.NewRequest("GET", "/apartments", nil)
	checkError(err, t)
	rr := httptest.NewRecorder()
	router.HandleFunc("/apartments", controller.GetAllHousing).Methods("GET")
	router.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Status code differs. Expected %d\n Got %d", http.StatusOK, status)
	}
}
func TestGetTypeHousing(t *testing.T) {
	req, err := http.NewRequest("GET", "/apartments/alala", nil)
	checkError(err, t)
	rr := httptest.NewRecorder()
	router.HandleFunc("/apartments/{type:flats|houses}", controller.GetTypeHousing).Methods("GET")
	router.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("Status code differs. Expected %d\n Got %d", http.StatusNotFound, status)
	}
}
func TestGetOneHousing(t *testing.T) {
	req, err := http.NewRequest("GET", "/apartments/alala", nil)
	checkError(err, t)
	rr := httptest.NewRecorder()
	router.HandleFunc("/apartments/{id:[0-9]+}", controller.GetOneHousing).Methods("GET")
	router.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("Status code differs. Expected %d\n Got %d", http.StatusNotFound, status)
	}
}
func TestGetRealtor(t *testing.T) {
	req, err := http.NewRequest("GET", "/apartments/4/realtor", nil)
	checkError(err, t)
	rr := httptest.NewRecorder()
	router.HandleFunc("/apartments/{id:[0-9]+}/realtor", controller.GetRealtor).Methods("GET")
	router.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusNetworkAuthenticationRequired {
		t.Errorf("Status code differs. Expected %d\n Got %d", http.StatusNetworkAuthenticationRequired, status)
	}
}
func TestLoginPageHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/login", nil)
	checkError(err, t)
	rr := httptest.NewRecorder()
	router.HandleFunc("/login", controller.LoginPageHandler).Methods("GET")
	router.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Status code differs. Expected %d\n Got %d", http.StatusOK, status)
	}
}
func TestIndexPageHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/index", nil)
	checkError(err, t)
	rr := httptest.NewRecorder()
	router.HandleFunc("/index", controller.IndexPageHandler).Methods("GET")
	router.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Status code differs. Expected %d\n Got %d", http.StatusOK, status)
	}
}
func TestRegisterPageHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/register", nil)
	checkError(err, t)
	rr := httptest.NewRecorder()
	router.HandleFunc("/register", controller.RegisterPageHandler).Methods("GET")
	router.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Status code differs. Expected %d\n Got %d", http.StatusOK, status)
	}
}
func TestLogOut(t *testing.T) {
	req, err := http.NewRequest("POST", "/index", nil)
	checkError(err, t)
	rr := httptest.NewRecorder()
	router.HandleFunc("/index", controller.LogOut).Methods("POST")
	router.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusFound {
		t.Errorf("Status code differs. Expected %d\n Got %d", http.StatusFound, status)
	}
}
