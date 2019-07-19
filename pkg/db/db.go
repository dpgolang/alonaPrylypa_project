package db

import (
	"Project/pkg/models"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"os"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "190117"
	dbname   = "flats"
)
type ApartmentsFinder interface{
	GetAllApartments()([]models.Flat,error)
	GetApartmentById(id int)(models.Flat,error)
	GetFlats()([]models.Flat,error)
	GetHouses()([]models.Flat,error)
}
type ApartmentsStorage struct {
	db *sql.DB
}
func NewAppartmentsStorage()ApartmentsFinder{
	db,err:=ConnectDB()
	if err!=nil{
		log.Printf("can't connect to db:%v",err)
		os.Exit(1)
	}
	return &ApartmentsStorage{db}
}
func (s ApartmentsStorage) GetAllApartments()([]models.Flat,error){
	rows, err := s.db.Query("SELECT * FROM living_spaces")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	flats := make([]models.Flat, 0)
	for rows.Next() {
		fl := new(models.Flat)
		err := rows.Scan(&fl.Id,&fl.Type,&fl.Street,&fl.Price,&fl.Square,&fl.Rooms,&fl.Floor)
		if err != nil {
			return nil, err
		}
		flats = append(flats, *fl)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return flats,nil
}
func (s ApartmentsStorage) GetFlats()([]models.Flat,error){
	rows, err := s.db.Query("SELECT * FROM living_spaces where type='Flat'")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	flats := make([]models.Flat, 0)
	for rows.Next() {
		fl := new(models.Flat)
		err := rows.Scan(&fl.Id,&fl.Type,&fl.Street,&fl.Price,&fl.Square,&fl.Rooms,&fl.Floor)
		if err != nil {
			return nil, err
		}
		flats = append(flats, *fl)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return flats,nil
}
func (s ApartmentsStorage) GetHouses()([]models.Flat,error){
	rows, err := s.db.Query("SELECT * FROM living_spaces where type='House'")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	flats := make([]models.Flat, 0)
	for rows.Next() {
		fl := new(models.Flat)
		err := rows.Scan(&fl.Id,&fl.Type,&fl.Street,&fl.Price,&fl.Square,&fl.Rooms,&fl.Floor)
		if err != nil {
			return nil, err
		}
		flats = append(flats, *fl)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return flats,nil
}
func (s ApartmentsStorage) GetApartmentById(id int)(models.Flat,error){
	var fl models.Flat
	row := s.db.QueryRow("select * from living_spaces where id = $1", id)
	err := row.Scan(&fl.Id,&fl.Type,&fl.Street,&fl.Price,&fl.Square,&fl.Rooms,&fl.Floor)
	if err != nil {
		return models.Flat{}, err
	}
	return fl,nil
}
func ConnectDB() (*sql.DB,error) {
	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err!=nil{
		return nil, err
	}
	err = db.Ping()
	if err!=nil{
		return nil, err
	}
	return db, nil
}
