package auth

import (
	"app/db"
	"app/helper"
	structs "app/struct"
	"encoding/json"
	"net/http"
)

func LoginUser(w http.ResponseWriter, r *http.Request) {
	var user structs.UserData
	// Decode the JSON request body into the struct
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		StructureResponse("Unable to decode JSON request body", "400", "true", "", w)
		return
	}

	if user.Email == "" {
		StructureResponse("Empty email field", "400", "true", "", w)
		return
	}
	if user.Password == "" {
		StructureResponse("Empty password field", "400", "true", "", w)
		return
	}

	user.Email = helper.FormatAndEscape(user.Email)
	user.Password = helper.FormatAndEscape(user.Password)

	exists, err := db.EmailExists(user.Email)
	if err != nil {
		StructureResponse("An error occured", "400", err.Error(), "", w)
		return
	}
	if exists {
		loggedIn, err := db.LoginUser(user.Email, user.Password)
		if err != nil {
			StructureResponse("Error Logging in", "400", err.Error(), "", w)
			return
		}

		if loggedIn {
			us, _ := db.GetUserbyEmail(user.Email)

			StructureResponse("User login successful", "200", "", us, w)
			return
		} else {
			StructureResponse("Invalid password", "400", "true", "", w)
			return
		}
	} else {
		StructureResponse("This email does not exist", "400", "true", "", w)
		return
	}

}
