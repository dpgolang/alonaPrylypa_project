package models

type Flat struct {
	Id     int     `json:"id"`
	Type   string  `json:"type"`
	Street string  `json:"street"`
	Price  int     `json:"price"`
	Square float64 `json:"square"`
	Rooms  int     `json:"rooms"`
	Floor  int     `json:"floor"`
	Realtor string `json:"realtor"`
}
type Customer struct {
	UserName string
	Email    string
	Password string
}
type User struct {
	Username      string
	Authenticated bool
}
type Realtor struct{
	Name string
	Phone int
	Email string
}