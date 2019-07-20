package repos

import (
	"github.com/alonaprylypa/Project/pkg/controllers"
	"github.com/gorilla/sessions"
	"io/ioutil"
)

func LoadFile(fileName string) (string, error) {
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
func GetUser(s *sessions.Session) controllers.User{
	val:=s.Values["user"]
	var user = controllers.User{}
	user,ok:=val.(controllers.User)
	if !ok{
		return controllers.User{Authenticated:false}
	}
	return user
}