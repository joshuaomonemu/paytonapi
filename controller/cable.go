package controller

import (
	"app/db"
	"app/helper"
	"app/models"
	structs "app/struct"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

type VariationDetails struct {
	Name            string `json:"name"`
	VariationCode   string `json:"variation_code"`
	VariationAmount string `json:"variation_amount"`
}

type Content struct {
	Transactions Transaction `json:"transactions"`
}

type Content1 struct {
	Transactions Transaction1 `json:"transactions"`
}

type TransactionDate struct {
	Date         string `json:"date"`
	TimezoneType int    `json:"timezone_type"`
	Timezone     string `json:"timezone"`
}

type DstvResponse struct {
	Code                string  `json:"code"`
	Content             Content `json:"content"`
	ResponseDescription string  `json:"response_description"`
	RequestID           string  `json:"requestId"`
	Amount              string  `json:"amount"`
	TransactionDate     string  `json:"transaction_date"`
	PurchasedCode       string  `json:"purchased_code"`
}
type DataResponse struct {
	Code                string   `json:"code"`
	Content             Content1 `json:"content"`
	ResponseDescription string   `json:"response_description"`
	RequestID           string   `json:"requestId"`
	Amount              string   `json:"amount"`
	TransactionDate     string   `json:"transaction_date"`
	PurchasedCode       string   `json:"purchased_code"`
}
type UtilityResponse struct {
	Code                string  `json:"code"`
	Content             Content `json:"content"`
	ResponseDescription string  `json:"response_description"`
	RequestID           string  `json:"requestId"`
	Amount              string  `json:"amount"`
	TransactionDate     string  `json:"transaction_date"`
	PurchasedCode       string  `json:"purchased_code"`
	ExchangeReference   string  `json:"exchangeReference"`
	ArrearsBalance      string  `json:"arrearsBalance"`
	AppliedToArrears    string  `json:"appliedToArrears"`
	Wallet              string  `json:"wallet"`
	Vat                 string  `json:"vat"`
	InvoiceNumber       string  `json:"invoiceNumber"`
	AppliedToWallet     string  `json:"appliedToWallet"`
	Units               string  `json:"units"`
	Token               string  `json:"token"`
}

type Transaction struct {
	Status              string      `json:"status"`
	ProductName         string      `json:"product_name"`
	UniqueElement       string      `json:"unique_element"`
	UnitPrice           string      `json:"unit_price"`
	Quantity            string      `json:"quantity"`
	ServiceVerification interface{} `json:"service_verification"` // Assuming this can be of any type, hence using interface{}
	Channel             string      `json:"channel"`
	Commission          string      `json:"commission"`
	TotalAmount         string      `json:"total_amount"`
	Discount            interface{} `json:"discount"` // Assuming this can be of any type, hence using interface{}
	Type                string      `json:"type"`
	Email               string      `json:"email"`
	Phone               string      `json:"phone"`
	Name                interface{} `json:"name"` // Assuming this can be of any type, hence using interface{}
	ConvenienceFee      string      `json:"convinience_fee"`
	Amount              string      `json:"amount"`
	Platform            string      `json:"platform"`
	Method              string      `json:"method"`
	TransactionID       string      `json:"transactionId"`
}

type Transaction1 struct {
	Status              string      `json:"status"`
	ProductName         string      `json:"product_name"`
	UniqueElement       string      `json:"unique_element"`
	UnitPrice           string      `json:"unit_price"`
	Quantity            int64       `json:"quantity"`
	ServiceVerification interface{} `json:"service_verification"` // Assuming this can be of any type, hence using interface{}
	Channel             string      `json:"channel"`
	Commission          string      `json:"commission"`
	TotalAmount         string      `json:"total_amount"`
	Discount            interface{} `json:"discount"` // Assuming this can be of any type, hence using interface{}
	Type                string      `json:"type"`
	Email               string      `json:"email"`
	Phone               string      `json:"phone"`
	Name                interface{} `json:"name"` // Assuming this can be of any type, hence using interface{}
	ConvenienceFee      string      `json:"convinience_fee"`
	Amount              string      `json:"amount"`
	Platform            string      `json:"platform"`
	Method              string      `json:"method"`
	TransactionID       string      `json:"transactionId"`
}

type AirtimeResponse struct {
	Code                string `json:"code"`
	ResponseDescription string `json:"response_description"`
	RequestID           string `json:"requestId"`
	TransactionID       string `json:"transactionId"`
	Amount              string `json:"amount"`
	TransactionDate     string `json:"transaction_date"`
	PurchasedCode       string `json:"purchased_code"`
}

type AirtimeDate struct {
	Date         string `json:"date"`
	TimezoneType int    `json:"timezone_type"`
	Timezone     string `json:"timezone"`
}

var (
	resp1 []byte
)

func Dstv(w http.ResponseWriter, r *http.Request) {
	resp, err := models.Dstv()
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

func DstvVerify(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	biller := params["id"]
	provider := "dstv"
	resp, err := models.DstvVerify(biller, provider)
	if err != nil {
		io.WriteString(w, err.Error())
		w.WriteHeader(500)
		return
	}

	io.WriteString(w, string(resp))

}

func DstvPay(w http.ResponseWriter, r *http.Request) {

	// Generate the full request ID
	reqID, _ := helper.GenerateRequestID(12)

	params := mux.Vars(r)
	biller := params["id"]
	provider := "dstv"
	amount := r.Header.Get("amount")
	phone := r.Header.Get("phone")
	email := r.Header.Get("email")
	variation_code := r.Header.Get("variation_code")
	quantity := r.Header.Get("quantity")
	subscription_type := r.Header.Get("subscription_type")
	//note := "Cable Subscription"
	date := helper.GetDate()
	time := helper.GetTime()

	_, err := db.CheckBalance(amount, email)
	if err != nil {
		w.WriteHeader(402)
		return
	}

	resp, err := models.DstvPay(biller, provider, amount, phone, subscription_type, variation_code, quantity, reqID)
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
			IconUrl: "assets/images/cable.png",
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
		w.WriteHeader(400)
		return
	} else {
		trans_stat = "Approved"
		db.WalletTrans(amount, email)
		trans := &db.Transaction{
			IconUrl: "assets/images/cable.png",
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
	}

	simp, _ := json.Marshal(response)

	io.WriteString(w, string(simp))

}

func Gotv(w http.ResponseWriter, r *http.Request) {
	resp, err := models.Gotv()
	if err != nil {
		io.WriteString(w, err.Error())
		w.WriteHeader(500)
		return
	}
	fmt.Println(resp)
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

func GotvVerify(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	biller := params["id"]
	provider := "gotv"
	resp, err := models.GotvVerify(biller, provider)
	if err != nil {
		io.WriteString(w, err.Error())
		w.WriteHeader(500)
		return
	}

	io.WriteString(w, string(resp))

}

func GotvPay(w http.ResponseWriter, r *http.Request) {

	// Generate the full request ID
	reqID, _ := helper.GenerateRequestID(12)

	params := mux.Vars(r)
	biller := params["id"]
	provider := "gotv"
	amount := r.Header.Get("amount")
	phone := r.Header.Get("phone")
	email := r.Header.Get("email")
	quantity := r.Header.Get("quantity")
	variation_code := r.Header.Get("variation_code")
	subscription_type := r.Header.Get("subscription_type")
	date := helper.GetDate()
	time := helper.GetTime()

	if subscription_type == "renew" {
		var err error
		resp1, err = models.GotvPay(biller, provider, amount, phone, subscription_type, reqID)
		if err != nil {
			io.WriteString(w, err.Error())
			w.WriteHeader(500)
			return
		}

	} else {
		var err error
		resp1, err = models.GotvPay1(biller, provider, amount, phone, variation_code, subscription_type, quantity, reqID)
		if err != nil {
			io.WriteString(w, err.Error())
			w.WriteHeader(500)
			return
		}
	}

	var response DstvResponse

	err := json.Unmarshal(resp1, &response)
	if err != nil {
		io.WriteString(w, err.Error())
		return
	} else {
		io.WriteString(w, string(resp1))
	}

	if response.Code != "000" {
		trans_stat = "Declined"
		trans := &db.Transaction{
			IconUrl: "assets/images/cable.png",
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
		w.WriteHeader(400)
		return
	} else {
		trans_stat = "Approved"
		db.WalletTrans(amount, email)
		trans := &db.Transaction{
			IconUrl: "assets/images/cable.png",
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
	}

	simp, _ := json.Marshal(response)

	io.WriteString(w, string(simp))

}

func Star(w http.ResponseWriter, r *http.Request) {
	resp, err := models.Star()
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

func StarVerify(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	biller := params["id"]
	provider := "gotv"
	resp, err := models.StarVerify(biller, provider)
	if err != nil {
		io.WriteString(w, err.Error())
		w.WriteHeader(500)
		return
	}

	io.WriteString(w, string(resp))

}

func StarPay(w http.ResponseWriter, r *http.Request) {

	// Generate the full request ID
	reqID, _ := helper.GenerateRequestID(12)

	params := mux.Vars(r)
	biller := params["id"]
	provider := "gotv"
	amount := r.Header.Get("amount")
	phone := r.Header.Get("phone")
	email := r.Header.Get("email")
	subscription_type := r.Header.Get("subscription_type")
	date := helper.GetDate()
	time := helper.GetTime()

	resp, err := models.StarPay(biller, provider, amount, phone, subscription_type, reqID)
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
			IconUrl: "assets/images/cable.png",
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
		w.WriteHeader(400)
		return
	} else {
		trans_stat = "Approved"
		db.WalletTrans(amount, email)
		trans := &db.Transaction{
			IconUrl: "assets/images/cable.png",
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
	}

	simp, _ := json.Marshal(response)

	io.WriteString(w, string(simp))

}
