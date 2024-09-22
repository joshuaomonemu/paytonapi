package controller

import (
	"app/db"
	"encoding/json"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

type Transaction12 struct {
	IconUrl string `json:"icon_url"`
	Title   string `json:"title"`
	Date    string `json:"date"`
	Time    string `json:"time"`
	Amount  string `json:"amount"`
	Status  string `json:"status"`
}

func Transactions(w http.ResponseWriter, r *http.Request) {

	transactions := []Transaction12{
		{
			IconUrl: "assets/images/data.png",
			Title:   "Buy Data",
			Date:    "19 September, 2023",
			Time:    "5:40 PM",
			Amount:  "₦3,609.00",
			Status:  "completed",
		},
		{
			IconUrl: "assets/images/cable.png",
			Title:   "Buy DsTv",
			Date:    "19 September, 2023",
			Time:    "5:31 PM",
			Amount:  "₦22,060.00",
			Status:  "completed",
		},
		{
			IconUrl: "assets/images/electricity.png",
			Title:   "Recharge Meter",
			Date:    "17 September, 2023",
			Time:    "1:11 PM",
			Amount:  "₦11,030.00",
			Status:  "completed",
		},
		{
			IconUrl: "assets/images/airtime.png",
			Title:   "Buy Airtime",
			Date:    "17 September, 2023",
			Time:    "3:30 PM",
			Amount:  "₦5,000.00",
			Status:  "completed",
		},
	}
	bs, _ := json.Marshal(transactions)

	io.WriteString(w, string(bs))

}

func GetTrans(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	user := params["id"]
	resp, err := db.GetTransactions(user)
	if err != nil {
		io.WriteString(w, err.Error())
		return
	}

	io.WriteString(w, string(resp))
}

func GetWalletTrans(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	user := params["id"]
	resp, err := db.GetWalletTransactions(user)
	if err != nil {
		w.WriteHeader(204)
		io.WriteString(w, err.Error())
		return
	}

	io.WriteString(w, string(resp))
}
