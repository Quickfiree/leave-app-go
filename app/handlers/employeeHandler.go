package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"leave-app/entity"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v4/pgxpool"
)

type test struct {
	dataB *pgxpool.Pool
}

func NewTest(db *pgxpool.Pool) test {
	return test{db}

}

func (ts test) EmployeeRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Get("/GetLeaveCount", ts.GetLeaveCount)
	r.Post("/ApplyForLeave", ts.ApplyForLeave)

	return r

}

func (ts test) GetHandler(w http.ResponseWriter, r *http.Request) {
	pgxpool := ts.dataB
	data, error := pgxpool.Query(context.Background(), `select userid, name, mobile, email, role from dev.userinfo`)
	if error != nil {
		fmt.Println("An error has occurred: ", error)
	}
	var rowArray []entity.Rows

	for data.Next() {
		var row entity.Rows
		error = data.Scan(&row.Userid, &row.Name, &row.Mobile, &row.Email, &row.Role)
		if error != nil {
			fmt.Println("An error has occured")
		}
		rowArray = append(rowArray, row)
	}
	response, error := json.Marshal(rowArray)
	if error != nil {
		fmt.Println("An error has occured")
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func (ts test) GetLeaveCount(w http.ResponseWriter, r *http.Request) {
	pgxpool := ts.dataB
	var reqBody entity.Rows
	json.NewDecoder(r.Body).Decode(&reqBody)
	userid := reqBody.Userid
	data := pgxpool.QueryRow(context.Background(), `select casualleave, sickleave, earnedleave from dev.leavecount where userid = $1`, userid)
	var writeData [3]int
	error := data.Scan(&writeData[0], &writeData[1], &writeData[2])
	if error != nil {
		fmt.Println("An error has occurred: ", error)
	}
	w.Header().Set("Content-Type", "application/json")
	req, error := json.Marshal(writeData)
	if error != nil {
		fmt.Println("An error has occured")
	}
	w.Write(req)
}

func (ts test) ApplyForLeave(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Triggered")
	pgxpool := ts.dataB
	var reqBody entity.ApplyForLeave
	userid := reqBody.Userid
	_, error := pgxpool.Exec(context.Background(), `insert into dev.leaverequests (leavetype, noofdays, startdate, enddate, userid, status, reason) values ($1, $2, $3, $4, $5, $6, $7)`, reqBody.LeaveType, reqBody.NoOfDays, reqBody.StartDate, reqBody.EndDate, userid, reqBody.Status, reqBody.Reason)
	if error != nil {
		fmt.Println("An error has occurred two: ", error)
	}
	req, error := json.Marshal("Applied for leave.")
	if error != nil {
		fmt.Println("An error has occurred three: ", error)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(req)
}
