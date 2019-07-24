//Package connects to the database and works with it
package db

import (
	"database/sql"
	"fmt"

	"github.com/alonaprylypa/Project/pkg/models"

	"log"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	"os"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Printf("No .env file found:%v", err)
	}
}

type DateFinder interface {
	GetAllApartments() ([]models.Flat, error)
	GetApartmentById(id int) (models.Flat, error)
	GetFlats() ([]models.Flat, error)
	GetHouses() ([]models.Flat, error)
	RegisterCustomer(name string, email string, pass string) (err error)
	ReturnCustomer(name string) (models.Customer, error)
	GetRealtorDate(id int) (models.Realtor, error)
}
type Storage struct {
	db *sql.DB
}

func NewAppartmentsStorage() DateFinder {
	db, err := ConnectDB()
	if err != nil {
		log.Printf("can't connect to db:%v", err)
		os.Exit(1)
	}
	return &Storage{db}
}
func (s Storage) GetAllApartments() ([]models.Flat, error) {
	rows, err := s.db.Query("SELECT * FROM living_spaces")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	flats := make([]models.Flat, 0)
	for rows.Next() {
		fl := new(models.Flat)
		err := rows.Scan(&fl.Id, &fl.Type, &fl.Street, &fl.Price, &fl.Square, &fl.Rooms, &fl.Floor, &fl.Realtor)
		if err != nil {
			return nil, err
		}
		flats = append(flats, *fl)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return flats, nil
}
func (s Storage) GetFlats() ([]models.Flat, error) {
	rows, err := s.db.Query("SELECT * FROM living_spaces where type='Flat'")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	flats := make([]models.Flat, 0)
	for rows.Next() {
		fl := new(models.Flat)
		err := rows.Scan(&fl.Id, &fl.Type, &fl.Street, &fl.Price, &fl.Square, &fl.Rooms, &fl.Floor, &fl.Realtor)
		if err != nil {
			return nil, err
		}
		flats = append(flats, *fl)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return flats, nil
}
func (s Storage) GetHouses() ([]models.Flat, error) {
	rows, err := s.db.Query("SELECT * FROM living_spaces where type='House'")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	flats := make([]models.Flat, 0)
	for rows.Next() {
		fl := new(models.Flat)
		err := rows.Scan(&fl.Id, &fl.Type, &fl.Street, &fl.Price, &fl.Square, &fl.Rooms, &fl.Floor, &fl.Realtor)
		if err != nil {
			return nil, err
		}
		flats = append(flats, *fl)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return flats, nil
}
func (s Storage) GetApartmentById(id int) (models.Flat, error) {
	var fl models.Flat
	row := s.db.QueryRow("select * from living_spaces where id = $1", id)
	err := row.Scan(&fl.Id, &fl.Type, &fl.Street, &fl.Price, &fl.Square, &fl.Rooms, &fl.Floor, &fl.Realtor)
	if err != nil {
		return models.Flat{}, err
	}
	return fl, nil
}
func (s Storage) RegisterCustomer(name string, email string, pass string) (err error) {
	_, err = s.db.Exec("insert into users values ($1, $2, $3)", name, email, pass)
	return
}
func (s Storage) ReturnCustomer(name string) (models.Customer, error) {
	result := s.db.QueryRow("select password from users where username=$1", name)
	storedCreds := models.Customer{}
	err := result.Scan(&storedCreds.Password)
	if err != nil {
		return models.Customer{}, err
	}
	return storedCreds, nil
}
func (s Storage) GetRealtorDate(id int) (models.Realtor, error) {
	result := s.db.QueryRow("SELECT r.name, r.number, r.email FROM realtor r left JOIN living_spaces liv ON liv.realtor = r.name WHERE liv.id = $1", id)
	storedCreds := models.Realtor{}
	err := result.Scan(&storedCreds.Name, &storedCreds.Phone, &storedCreds.Email)
	if err != nil {
		return models.Realtor{}, err
	}
	return storedCreds, nil
}
func ConnectDB() (*sql.DB, error) {
	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		os.Getenv("SERVICE_HOST"), os.Getenv("SERVICE_PORT_BD"),
		os.Getenv("SERVICE_USER"), os.Getenv("SERVICE_PASSWORD"), os.Getenv("SERVICE_DBNAME"))

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
