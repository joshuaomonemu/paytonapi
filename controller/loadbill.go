package controller

import (
	"app/models"
	structs "app/struct"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

func LoadBillers(w http.ResponseWriter, r *http.Request) {
	resp, err := models.LoadBillers()
	if err != nil {
		io.WriteString(w, err.Error())
		return
	}
	//io.WriteString(w, string(resp))
	// var response *structs.ElectricityPaymentOption
	// err = json.Unmarshal(resp, &response)
	// if err != nil {
	// 	io.WriteString(w, err.Error())
	// }

	// jsn, err := json.Marshal(response)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	io.WriteString(w, string(resp))

}

func PayBill(w http.ResponseWriter, r *http.Request) {

	var billreq *structs.UtilBill

	bs, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(406)
	}
	json.Unmarshal(bs, &billreq)

	bs1, _ := json.Marshal(billreq)

	resp, err := models.PayBill(bs1)
	if err != nil {
		io.WriteString(w, err.Error())
		return
	}

	// var response *Structs.ElectricityPaymentOption
	// err = json.Unmarshal(resp, &response)
	// if err != nil {
	// 	io.WriteString(w, err.Error())
	// 	return
	// }

	// jsn, err := json.Marshal(response)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	//fmt.Println(jsn)
	io.WriteString(w, string(resp))
}
