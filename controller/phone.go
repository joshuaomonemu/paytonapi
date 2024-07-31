package controller

import (
	"app/helper"
	"app/models"
	"io"
	"net/http"
)

func PhonePay(w http.ResponseWriter, r *http.Request) {

	// Generate the full request ID
	reqID, _ := helper.GenerateRequestID(12)

	provider := r.Header.Get("provider")
	amount := r.Header.Get("amount")
	phone := r.Header.Get("phone")

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

	// err = json.Unmarshal(resp, &response)
	// if err != nil {
	// 	io.WriteString(w, err.Error())
	// 	return
	// }
	// simp, _ := json.Marshal(response)

	io.WriteString(w, string(resp))

}
