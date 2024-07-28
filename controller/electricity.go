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
