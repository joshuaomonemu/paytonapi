package controller

import (
	"app/models"
	"io"
	"net/http"
)

func GetGiftCards(w http.ResponseWriter, r *http.Request) {

	// bs, err := ioutil.ReadAll(r.Body)
	// if err != nil {
	// 	w.WriteHeader(406)
	// }
	// json.Unmarshal(bs, &billreq)

	//bs1, _ := json.Marshal(billreq)

	resp, err := models.AllCards()
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
