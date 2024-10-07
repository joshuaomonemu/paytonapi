package auth

import (
	"app/helper"
	structs "app/struct"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	var user structs.UserData

	// Decode the JSON request body into the struct
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Unable to decode JSON request body", http.StatusBadRequest)
		return
	}

	if user.Fname == "" {
		io.WriteString(w, "No Firstname")
		w.WriteHeader(202)
		return
	}

	if user.Lname == "" {
		io.WriteString(w, "No Lastname")
		w.WriteHeader(202)
		return
	}
	if user.Email == "" {
		io.WriteString(w, "No Email")
		w.WriteHeader(202)
		return
	}
	if user.Password == "" {
		io.WriteString(w, "No Password")
		w.WriteHeader(202)
		return
	}
	if user.Phone == "" {
		io.WriteString(w, "No Phone Number")
		w.WriteHeader(202)
		return
	}

	firstname := helper.FormatAndEscape(user.Fname)
	lastname := helper.FormatAndEscape(user.Lname)
	email := helper.FormatAndEscape(user.Email)
	password := helper.FormatAndEscape(user.Password)
	phone := helper.FormatAndEscape(user.Phone)
	fmt.Println(firstname, lastname, email, password, phone)

	// Access the fields
	//fmt.Printf("First Name: %s, Last Name: %s, Email: %s\n", user.Fname, user.Lname, user.Email)
}
