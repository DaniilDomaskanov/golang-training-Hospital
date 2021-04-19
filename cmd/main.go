package main

import (
	"fmt"
	"golang-training-Hospital/pkg/data"
	"golang-training-Hospital/pkg/db"
	"log"
	"os"
)

var (
	host     = os.Getenv("DB_USERS_HOST")
	port     = os.Getenv("DB_USERS_PORT")
	user     = os.Getenv("DB_USERS_USER")
	dbname   = os.Getenv("DB_USERS_DBNAME")
	password = os.Getenv("DB_USERS_PASSWORD")
	sslmode  = os.Getenv("DB_USERS_SSL")
)

func init() {
	if host == "" {
		host = "localhost"
	}
	if port == "" {
		port = "5432"
	}
	if user == "" {
		user = "postgres"
	}
	if dbname == "" {
		dbname = "Hospital"
	}
	if password == "" {
		password = "root"
	}
	if sslmode == "" {
		sslmode = "disable"
	}
}

func main() {
	connection, err := db.GetConnection(host, port, user, dbname, password, sslmode)
	if err != nil {
		log.Fatalf("can't connect to database, err2: %v", err)
	}
	doctor := data.NewDoctor(connection)
	doc := data.Doctors{
		FirstName:        "Valery",
		LastName:         "Aleksandrov",
		DateOfBirth:      "2017-01-22",
		Salary:           "12131.121",
		CurrentBusyState: false,
		GenderId:         2,
		SpecialityId:     1,
	}
	id, errCreat := doctor.CreateDoctor(doc)
	if errCreat != nil {
		log.Println(errCreat)
	}
	fmt.Println("Inserted doctor id is:", id)
	errDel := doctor.DeleteDoctor(5)
	if errDel != nil {
		log.Println(errDel)
	}
	valuesData := data.Doctors{
		FirstName: "Dmitry",
		LastName:  "Putkov",
	}
	idSearch, err2 := doctor.UpdateDoctor(2, valuesData)
	if err2 != nil {
		log.Println(err2)
	}
	fmt.Println("Updated doctor id is:", idSearch)
	doctors, err := doctor.ReadAll()
	if err != nil {
		log.Println(err)
	}
	log.Println(doctors)
	doctor.ExecInnerJoin()
}
