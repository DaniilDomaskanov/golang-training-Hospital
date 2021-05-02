package main

import (
	"log"
	"net"
	"net/http"
	"os"

	"github.com/DaniilDomaskanov/golang-training-Hospital/pkg/api"
	"github.com/DaniilDomaskanov/golang-training-Hospital/pkg/data"
	"github.com/DaniilDomaskanov/golang-training-Hospital/pkg/db"

	"github.com/gorilla/mux"
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
	// 2. create router that allows to set routes
	r := mux.NewRouter()
	// 3. connect to data layer
	doctorData := data.NewDoctor(connection)
	// 4. send data layer to api layer
	api.ServeDoctorResources(r, *doctorData)
	// 5. cors for making requests from any domain
	r.Use(mux.CORSMethodMiddleware(r))

	// 6. start server
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("Server Listen port...")
	}
	if err := http.Serve(listener, r); err != nil {
		log.Fatal("Server has been crashed...")
	}
}
