package auth

import (
	"app/db"
	"app/mail"
	structs "app/struct"
	"encoding/hex"
	"encoding/json"
	"io"
	"math/rand"
	"net/http"
)

// User model
// type User struct {
// 	ID          uint
// 	Email       string
// 	Password    string
// 	Token       string
// 	TokenExpiry time.Time
// }

type PageData struct {
	Token string
	Error string
}

func generateResetToken() (string, error) {
	// Generate a random 32-byte token
	token := make([]byte, 32)
	if _, err := rand.Read(token); err != nil {
		return "", err
	}
	return hex.EncodeToString(token), nil
}

// RequestPasswordReset handles the forgot password request
func RequestPasswordReset(w http.ResponseWriter, r *http.Request) {
	var user structs.UserData
	// Decode the JSON request body into the struct
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		StructureResponse("Unable to decode JSON request body", "400", "true", "", w)
		return
	}

	//CHECKING IF THE EMAIL FIELDS ARE EMPTY
	if user.Email == "" {
		StructureResponse("Empty email field", "400", "true", "", w)
		return
	}

	// Check if user exists
	exists, err := db.EmailExists(user.Email)
	if err != nil {
		StructureResponse("An error occured", "400", err.Error(), "", w)
		return
	}
	if !exists {
		StructureResponse("Email does not exist", "400", "Email does not exist", "", w)
		return
	}

	//Generate request token for password reset
	token, err := generateResetToken()
	if err != nil {
		StructureResponse("An error occured", "400", err.Error(), "", w)
		return
	}

	// Store the token in the database
	err = db.StoreResetToken(user.Email, token)
	if err != nil {
		StructureResponse("An error occured", "400", err.Error(), "", w)
		return
	}

	linke := "https://payton.jitssolutions.com/reset.html?token=" + token

	mail.ResetMail(user.Email, linke)

	StructureResponse("Email Sent", "200", "false", "", w)

}

func SetPassword(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	token := r.FormValue("token")
	password := r.FormValue("password")
	confirmPassword := r.FormValue("confirm_password")

	if password != confirmPassword {
		io.WriteString(w, "Passwords are not thesame")
		return
	}

	db.ResetPassword(token, password)

}
