package auth

import (
	"app/db"
	structs "app/struct"
	"encoding/json"
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

// // GenerateToken creates a random reset token
// func GenerateToken() (string, error) {
// 	bytes := make([]byte, 16)
// 	if _, err := rand.Read(bytes); err != nil {
// 		return "", err
// 	}
// 	return hex.EncodeToString(bytes), nil
// }

// RequestPasswordReset handles the forgot password request
func RequestPasswordReset(w http.ResponseWriter, r *http.Request) {
	var user structs.UserData
	// Decode the JSON request body into the struct
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		StructureResponse("Unable to decode JSON request body", "400", "true", "", w)
		return
	}

	//CHECKING IF THE EMAIL AND PASSWORD FIELDS ARE EMPTY
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
		StructureResponse("An error occured", "400", err.Error(), "", w)
		return
	}

}
