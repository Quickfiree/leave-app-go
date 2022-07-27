package server

import (
	"leave-app/dbConnection"
	"leave-app/handlers"
	"net/http"

	"github.com/go-chi/chi/v5"
)

// Initialise connection to db and router

func Router() {
	r := chi.NewRouter()
	conn := dbConnection.DBConn()
	testing := handlers.NewTest(conn)
	r.Mount("/employee", testing.EmployeeRouter())
	r.Mount("/cred", testing.CredentialRouter())
	// r.Mount("/manager", testing.ManagerRouter())
	http.ListenAndServe(":8080", r)
}
