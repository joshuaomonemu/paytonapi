package controller

import (
	"app/db"
	"io"
	"net/http"
)

func GetTrans(w http.ResponseWriter, r *http.Request) {
	db.InsertReq()

	io.WriteString(w, "db.GetTrans()")
}
