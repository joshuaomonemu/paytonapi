package controller

import (
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
	provider := r.Header.Get("provider")
	amount := r.Header.Get("amount")
	phone := r.Header.Get("phone")
	variation_code := r.Header.Get("variation_code")

	resp, err := models.ElectPay1(biller, provider, amount, phone, variation_code, reqID)
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
	simp, _ := json.Marshal(response)

	io.WriteString(w, string(simp))

}
