package controller

import (
	"app/db"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

func GetTrans(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	user := params["id"]
	resp := db.GetTransactions(user)

	io.WriteString(w, string(resp))
}

func UpdateBalance(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	user := params["id"]
	amount := r.Header.Get("amount")
	err := db.UpdateBalance(user, amount)
	if err != nil {
		w.WriteHeader(402)
		io.WriteString(w, "error updating balance")
	}
}
