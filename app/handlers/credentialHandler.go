package handlers

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"leave-app/entity"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (ts test) CredentialRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Post("/Register", ts.RegisterHandler)
	r.Post("/Login", ts.LoginHandler)

	return r
}

func (ts test) LoginHandler(w http.ResponseWriter, r *http.Request) {
	type Registered struct {
		Registered bool   `json:"registered"`
		Userid     string `json:"userid, omitempty"`
		LoggedIn   bool   `json:"loggedIn"`
	}
	var reg Registered
	pgxpool := ts.dataB
	var reqBody entity.Credentials
	json.NewDecoder(r.Body).Decode(&reqBody)
	userid := ts.GetUserID(reqBody.Email)
	if userid.Valid {
		reg.Registered = true

		password := pgxpool.QueryRow(context.Background(), `select password from dev.usercreds where password = $1`, reqBody.Password)
		var pwd sql.NullString
		error := password.Scan(&pwd)
		if error != nil {
			fmt.Println("An error has occurred two: ", error)
		}
		if pwd.Valid {
			reg.Userid = userid.String
			reg.LoggedIn = true
			response, error := json.Marshal(reg)
			// error := json.NewEncoder(w).Encode(reg)
			fmt.Println(reg)
			if error != nil {
				fmt.Println("An error has occurred in pwd response: ", error)
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(200)
			w.Write(response)
		} else {
			response, error := json.Marshal(reg)
			fmt.Println(reg, response)
			if error != nil {
				fmt.Println("An error has occurred in pwd response: ", error)
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(500)
			w.Write(response)

		}

	} else {

		response, error := json.Marshal(reg)
		if error != nil {
			fmt.Println("An error has occurred in response: ", error)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		w.Write(response)

	}
}

func (ts test) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	pgxpool := ts.dataB
	var reqBody entity.Register
	json.NewDecoder(r.Body).Decode(&reqBody)
	_, error := pgxpool.Exec(context.Background(), `insert into dev.userinfo (name, mobile, email, role) values ($1, $2, $3, $4)`, reqBody.Name, reqBody.Mobile, reqBody.Email, "1")
	if error != nil {
		fmt.Println("An error has occurred: ", error)
	}
	uid := ts.GetUserID(reqBody.Email)
	if uid.Valid {
		_, error = pgxpool.Exec(context.Background(), `insert into dev.usercreds (username, password, userid) values ($1, $2, $3)`, reqBody.Email, reqBody.Password, uid.String)
		if error != nil {
			fmt.Println("An error has occurred: ", error)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
	} else {
		w.WriteHeader(500)
	}
}

func (ts test) GetUserID(email string) sql.NullString {
	pgxpool := ts.dataB
	userid := pgxpool.QueryRow(context.Background(), `select userid from dev.userinfo where email = $1`, email)
	var uid sql.NullString
	error := userid.Scan(&uid)
	if error != nil {
		fmt.Println("An error has occurred two: ", error)
	}
	return uid

}
