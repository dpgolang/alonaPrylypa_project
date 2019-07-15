package models

type Flat struct {
	Id     int
	Type   string
	Street string
	Price  int
	Square float64
	Rooms  int
	Floor  int
}

type Customer struct{
	UserName string
	Email string
	Password string
}