package controller

import (
	"app/db"
	"app/helper"
	"app/mail"
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
	note := "Airtime Purchase"
	date := helper.GetDate()
	time := helper.GetTime()

	_, err := db.CheckBalance(amount, email)
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

	var response AirtimeResponse

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
			Amount:  "₦" + amount,
			Status:  trans_stat,
			User:    email,
		}
		err := db.SetTransaction(trans)
		if err != nil {
			io.WriteString(w, err.Error())
			w.WriteHeader(400)
			return
		}
	} else {
		db.WalletTrans(amount, email)
		trans_stat = "Approved"
		trans := &db.Transaction{
			IconUrl: "assets/images/airtime.png",
			Title:   provider,
			Date:    date,
			Time:    time,
			Amount:  "₦" + amount,
			Status:  trans_stat,
			User:    email,
		}
		err := db.SetTransaction(trans)
		if err != nil {
			io.WriteString(w, err.Error())
			return
		}

		mail.AirtimeMail(email, note, phone, amount)
	}

	simp, _ := json.Marshal(response)

	io.WriteString(w, string(simp))

}
