package controller

import (
	"app/auth"
	"io"
	"net/http"
)

func Auth(w http.ResponseWriter, r *http.Request) {
	token, err := auth.Auth2()
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		io.WriteString(w, "Token error")
	}

	io.WriteString(w, token)

}
