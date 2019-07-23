package repos

import (
	"encoding/json"
	"github.com/alonaprylypa/Project/pkg/models"
	"github.com/gorilla/sessions"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

func LoadFile(fileName string) (string, error) {
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
func GetUser(s *sessions.Session) models.User {
	val := s.Values["user"]
	var user = models.User{}
	user, ok := val.(models.User)
	if !ok {
		return models.User{Authenticated: false}
	}
	return user
}
func CurrensyExchange(currency string, fl *models.Flat){
	current := time.Now()
	var data models.Privat
	url:="https://api.privatbank.ua/p24api/exchange_rates?json&date="+current.Format("02.01.2006")
	r, err := http.Get(url)
	if err != nil {
		log.Printf("unable to get data from privat:%v",err)
		return
	}
	defer r.Body.Close()
	dec := json.NewDecoder(r.Body)
	dec.Decode(&data)
	for _,val:=range data.ExchangeRate{
		if val.Currency==strings.ToUpper(currency){
			fl.Price=fl.Price*val.SaleRate
		}
	}
}
