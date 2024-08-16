package controller

import (
	"app/db"
	"encoding/json"
	"io"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {

	io.WriteString(w, "<h1>WELCOME TO PAYTON</h1>")

}

func Users(w http.ResponseWriter, r *http.Request) {

	users := db.GetUser()
	bs, _ := json.Marshal(users)

	io.WriteString(w, string(bs))

}

func UpdateWallet(w http.ResponseWriter, r *http.Request) {
	email := r.Header.Get("email")
	amount := r.Header.Get("amount")
	err := db.UpdateBalance(email, amount)
	if err != nil {
		io.WriteString(w, "error updating wallet balance")
	}
	w.WriteHeader(202)
}
