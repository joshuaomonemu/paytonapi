package controller

import (
	"app/helper"
	"app/models"
	structs "app/struct"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

func Smile(w http.ResponseWriter, r *http.Request) {
	resp, err := models.Smile()
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

func SmileVerify(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	biller := params["id"]
	provider := "smile-direct"
	resp, err := models.SmileVerify(biller, provider)
	if err != nil {
		io.WriteString(w, err.Error())
		w.WriteHeader(500)
		return
	}

	io.WriteString(w, string(resp))

}

func SmilePay(w http.ResponseWriter, r *http.Request) {

	// Generate the full request ID
	reqID, _ := helper.GenerateRequestID(12)

	params := mux.Vars(r)
	biller := params["id"]
	provider := "smile-direct"
	amount := r.Header.Get("amount")
	phone := r.Header.Get("phone")
	variation_code := r.Header.Get("variation_code")

	resp, err := models.DstvPay(biller, provider, amount, phone, variation_code, reqID)
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
