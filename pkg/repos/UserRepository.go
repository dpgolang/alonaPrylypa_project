package repos

import "io/ioutil"

func UserIsValid(uName, pwd string) bool {
	_uName, _pwd, isValid := "alyonka", "190117", false
	if uName == _uName && pwd == _pwd {
		isValid = true
	}
	return isValid
}
func IsEmpty(data string) bool {
	if len(data) <= 0 {
		return true
	} else {
		return false
	}
}
func LoadFile(fileName string) (string, error) {
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
