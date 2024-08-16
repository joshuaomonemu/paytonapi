package controller

import (
	"app/db"
	"app/helper"
	"app/models"
	"encoding/json"
	"io"
	"net/http"
)

var (
	trans_stat = ""
)

func PhonePay(w http.ResponseWriter, r *http.Request) {

	// Generate the full request ID
	reqID, _ := helper.GenerateRequestID(12)

	provider := r.Header.Get("provider")
	amount := r.Header.Get("amount")
	phone := r.Header.Get("phone")
	email := r.Header.Get("email")
	date := helper.GetDate()
	time := helper.GetTime()

	err, _ := db.CheckBalance(amount, email)
	if err != nil {
		w.WriteHeader(402)
		return
	}

	if provider == "0" {
		provider = "mtn"
	} else if provider == "1" {
		provider = "airtel"
	} else if provider == "2" {
		provider = "etisalat"
	} else if provider == "3" {
		provider = "glo"
	}

	resp, err := models.PhonePay(provider, amount, phone, reqID)
	if err != nil {
		io.WriteString(w, err.Error())
		w.WriteHeader(500)
		return
	}

	var response DstvResponse

	err = json.Unmarshal(resp, &response)
	if err != nil {
		io.WriteString(w, err.Error())
		return
	}

	if response.Code != "000" {
		trans_stat = "Declined"
		trans := &db.Transaction{
			IconUrl: "assets/images/airtime.png",
			Title:   provider,
			Date:    date,
			Time:    time,
			Amount:  amount,
			Status:  trans_stat,
			User:    email,
		}
		db.SetTransaction(trans)
		w.WriteHeader(400)
		return
	} else {
		trans_stat = "Approved"
		db.WalletTrans(amount, email)
		trans_stat = "Declined"
		trans := &db.Transaction{
			IconUrl: "assets/images/airtime.png",
			Title:   provider,
			Date:    date,
			Time:    time,
			Amount:  amount,
			Status:  trans_stat,
			User:    email,
		}
		db.SetTransaction(trans)
	}

	// err = json.Unmarshal(resp, &response)
	// if err != nil {
	// 	io.WriteString(w, err.Error())
	// 	return
	// }
	// simp, _ := json.Marshal(response)

	io.WriteString(w, string(resp))

}
