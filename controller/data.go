package controller

import (
	"app/db"
	"app/helper"
	"app/mail"
	"app/models"
	structs "app/struct"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

func Data(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	network := params["id"]
	if network == "0" {
		network = "mtn-data"
	} else if network == "1" {
		network = "airtel-data"
	} else if network == "2" {
		network = "etisalat-data"
	} else if network == "3" {
		network = "glo-data"
	}
	fmt.Println(network)
	resp, err := models.Data(network)
	if err != nil {
		io.WriteString(w, err.Error())
		w.WriteHeader(500)
		return
	}
	var response structs.Response
	err = json.Unmarshal(resp, &response)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}

	var variationDetails []VariationDetails

	// Loop through variations and populate the new struct
	for _, variation := range response.Content.Variations {
		variationDetails = append(variationDetails, VariationDetails{
			Name:            variation.Name,
			VariationCode:   variation.VariationCode,
			VariationAmount: variation.VariationAmount,
		})
	}

	// Marshal the VariationDetails slice back to JSON
	detailsJSON, err := json.Marshal(variationDetails)
	if err != nil {
		panic(err)
	}
	io.WriteString(w, string(detailsJSON))

}

func DataPay(w http.ResponseWriter, r *http.Request) {

	// Generate the full request ID
	reqID, _ := helper.GenerateRequestID(12)

	params := mux.Vars(r)
	biller := params["id"]

	provider := r.Header.Get("provider")
	amount := r.Header.Get("amount")
	phone := r.Header.Get("phone")
	variation_code := r.Header.Get("variation_code")
	email := r.Header.Get("email")
	note := "Data Purchase"
	date := helper.GetDate()
	time := helper.GetTime()

	if provider == "0" {
		provider = "mtn-data"
	} else if provider == "1" {
		provider = "airtel-data"
	} else if provider == "2" {
		provider = "etisalat-data"
	} else if provider == "3" {
		provider = "glo-data"
	}
	io.WriteString(w, biller+provider+amount+phone+variation_code+reqID)

	resp, err := models.DataPay(biller, provider, amount, phone, variation_code, reqID)
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
			IconUrl: "assets/images/data.png",
			Title:   provider,
			Date:    date,
			Time:    time,
			Amount:  amount + "₦",
			Status:  trans_stat,
			User:    email,
		}
		err := db.SetTransaction(trans)
		if err != nil {
			io.WriteString(w, err.Error())
			w.WriteHeader(400)
			return
		}
	} else if response.Code == "011" {
		trans_stat = "Approved"
		db.WalletTrans(amount, email)
		trans := &db.Transaction{
			IconUrl: "assets/images/data.png",
			Title:   provider,
			Date:    date,
			Time:    time,
			Amount:  amount + "₦",
			Status:  trans_stat,
			User:    email,
		}
		err := db.SetTransaction(trans)
		if err != nil {
			io.WriteString(w, err.Error())
			return
		}

		mail.AirtimeMail(email, note, phone, amount)
	} else {
		trans_stat = "Approved"
		db.WalletTrans(amount, email)
		trans := &db.Transaction{
			IconUrl: "assets/images/data.png",
			Title:   provider,
			Date:    date,
			Time:    time,
			Amount:  amount + "₦",
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
