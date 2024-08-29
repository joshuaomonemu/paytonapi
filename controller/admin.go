package controller

import (
	"app/models"
	"io"
	"net/http"
)

func Balance(w http.ResponseWriter, r *http.Request) {
	resp, err := models.Balance()
	if err != nil {
		io.WriteString(w, err.Error())
		w.WriteHeader(500)
		return
	}

	io.WriteString(w, string(resp))

}
