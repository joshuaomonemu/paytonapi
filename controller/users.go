package controller

import (
	"app/db"
	"encoding/json"
	"io"
	"net/http"
)

func Users(w http.ResponseWriter, r *http.Request) {

	users := db.GetUser()
	// if err != nil {
	// 	io.WriteString(w, string(err.Error()))
	// 	return
	// }
	bs, _ := json.Marshal(users)

	io.WriteString(w, string(bs))

}

// func UpdateWallet(w http.ResponseWriter, r *http.Request) {
// 	email := r.Header.Get("email")
// 	amount := r.Header.Get("amount")
// 	err := db.UpdateBalance(email, amount)
// 	if err != nil {
// 		io.WriteString(w, "error updating wallet balance")
// 	}
// 	w.WriteHeader(202)
// }
