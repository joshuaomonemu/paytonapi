package controller

import (
	"app/db"
	"app/helper"
	"encoding/json"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

func Users(w http.ResponseWriter, r *http.Request) {

	users, err := db.GetUser()
	if err != nil {
		io.WriteString(w, string(err.Error()))
		return
	}
	bs, err := json.Marshal(users)
	if err != nil {
		io.WriteString(w, string(err.Error()))
		return
	}

	io.WriteString(w, string(bs))

}

func UpdateWallet(w http.ResponseWriter, r *http.Request) {

	email := r.Header.Get("email")
	amount := r.Header.Get("amount")

	params := mux.Vars(r)
	stat := params["id"]
	date := helper.GetDate()
	time := helper.GetTime()

	if stat != "000" {
		trans_stat = "Declined"
		trans := &db.Transaction{
			IconUrl: "assets/images/wallet.png",
			Title:   "Wallet TopUp",
			Date:    date,
			Time:    time,
			Amount:  "₦" + amount,
			Status:  trans_stat,
			User:    email,
		}
		err := db.SetWallets(trans)
		if err != nil {
			io.WriteString(w, err.Error())
			w.WriteHeader(400)
			return
		}
	} else {
		err := db.UpdateBalance(email, amount)
		if err != nil {
			io.WriteString(w, "error updating wallet balance")
			w.WriteHeader(202)
		}
		trans_stat = "Approved"
		trans := &db.Transaction{
			IconUrl: "assets/images/wallet.png",
			Title:   "Wallet TopUp",
			Date:    date,
			Time:    time,
			Amount:  "₦" + amount,
			Status:  trans_stat,
			User:    email,
		}
		err1 := db.SetWallets(trans)
		if err1 != nil {
			io.WriteString(w, err1.Error())
			return
		}
	}
}
