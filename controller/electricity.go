package controller

import (
	"app/db"
	"app/helper"
	"app/models"
	"encoding/json"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

func ElectVerify(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	biller := params["id"]
	provider := r.Header.Get("provider")
	typer := r.Header.Get("type")

	switch provider {
	case "1":
		provider = "ikeja-electric"
	case "2":
		provider = "eko-electric"
	case "3":
		provider = "kano-electric"
	case "4":
		provider = "portharcourt-electric"
	case "5":
		provider = "jos-electric"
	case "6":
		provider = "ibadan-electric"
	case "7":
		provider = "kaduna-electric"
	case "8":
		provider = "abuja-electric"
	case "9":
		provider = "enugu-electric"
	case "10":
		provider = "benin-electric"
	case "11":
		provider = "aba-electric"
	case "12":
		provider = "yola-electric"

	}
	resp, err := models.ElectVerify(biller, provider, typer)
	if err != nil {
		io.WriteString(w, err.Error())
		w.WriteHeader(500)
		return
	}

	io.WriteString(w, string(resp))

}

func ElectPay1(w http.ResponseWriter, r *http.Request) {

	// Generate the full request ID
	reqID, _ := helper.GenerateRequestID(12)

	params := mux.Vars(r)
	biller := params["id"]
	//
	provider := r.Header.Get("provider")
	amount := r.Header.Get("amount")
	phone := r.Header.Get("phone")
	variation_code := r.Header.Get("variation_code")
	email := r.Header.Get("email")
	//note := "Utility Bill"
	date := helper.GetDate()
	time := helper.GetTime()
	//
	switch provider {
	case "1":
		provider = "ikeja-electric"
	case "2":
		provider = "eko-electric"
	case "3":
		provider = "kano-electric"
	case "4":
		provider = "portharcourt-electric"
	case "5":
		provider = "jos-electric"
	case "6":
		provider = "ibadan-electric"
	case "7":
		provider = "kaduna-electric"
	case "8":
		provider = "abuja-electric"
	case "9":
		provider = "enugu-electric"
	case "10":
		provider = "benin-electric"
	case "11":
		provider = "aba-electric"
	case "12":
		provider = "yola-electric"

	}

	_, err := db.CheckBalance(amount, email)
	if err != nil {
		w.WriteHeader(402)
		return
	}

	resp, err := models.ElectPay1(biller, provider, amount, phone, variation_code, reqID)
	if err != nil {
		io.WriteString(w, err.Error())
		w.WriteHeader(500)
		return
	}

	var response UtilityResponse

	err = json.Unmarshal(resp, &response)
	if err != nil {
		io.WriteString(w, err.Error())
		return
	}

	if response.Code != "000" {
		trans_stat = "Declined"
		trans := &db.Transaction{
			IconUrl: "assets/images/electricity.png",
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
		trans_stat = "Approved"
		db.WalletTrans(amount, email)
		trans := &db.Transaction{
			IconUrl: "assets/images/electricity.png",
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

		//mail.AirtimeMail(email, note, phone, amount)
	}
	simp, _ := json.Marshal(response)

	io.WriteString(w, string(simp))

}
