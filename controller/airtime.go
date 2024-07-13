package controller

import (
	"app/models"
	"io"
	"net/http"
)

func Airtime(w http.ResponseWriter, r *http.Request) {

	number := r.Header.Get("Number")
	amount := r.Header.Get("Amount")
	resp, err := models.Airtime(number, amount)
	if err != nil {
		io.WriteString(w, err.Error())
	}

	io.WriteString(w, resp)

}
