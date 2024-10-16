package auth

import (
	"app/db"
	"app/helper"
	"app/mail"
	structs "app/struct"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"strconv"
)

func StructureResponse(message, status, error1 string, data interface{}, writer http.ResponseWriter) {
	data21 := &structs.UserResponse{
		Message:  message,
		Status:   status,
		Error:    error1,
		UserData: data,
	}
	//STORING STRUCT INTO USER BYTE SLICE
	us_bs, _ := json.Marshal(data21)
	//CONVERTING STATUS CODE TO INT
	statu, _ := strconv.Atoi(status)
	//WRITING THE STATUS CODE TO THE HEADER
	writer.WriteHeader(statu)
	//WRITING BACK THE RESPONSE DATA
	io.WriteString(writer, string(us_bs))

}

func generateOTP(email string) (string, error) {
	// Generate a random 6-digit OTP
	otp := fmt.Sprintf("%04d", rand.Intn(10000))
	return otp, nil
}

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	var user structs.UserData
	// Decode the JSON request body into the struct
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		StructureResponse("Unable to decode JSON request body", "400", "true", "", w)
		return
	}

	if user.Fullname == "" {
		StructureResponse("Empty name field", "400", "true", "", w)
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
	if user.Phone == "" {
		StructureResponse("Empty phone number field", "400", "true", "", w)
		return
	}

	user.Fullname = helper.FormatAndEscape(user.Fullname)
	user.Email = helper.FormatAndEscape(user.Email)
	user.Password = helper.FormatAndEscape(user.Password)
	user.Phone = helper.FormatAndEscape(user.Phone)

	exists, err := db.EmailExists(user.Email)
	if err != nil {
		StructureResponse("Error checking user", "400", err.Error(), "", w)
		return
	}
	if exists {
		StructureResponse("Email already exists", "400", "true", "", w)
		return
	}

	err1 := db.SetUser(&user)
	if err1 != nil {
		StructureResponse("error creating user", "400", err1.Error(), "", w)
		return
	}

	//GENERATE OTP CODE
	otp, _ := generateOTP(user.Email)
	//MAIL OTP TO USERS EMAIL
	mail.OtpMail(user.Email, otp)
	//STORE OTP PIN IN DATABASE
	db.StoreOTP(user.Email, otp)

	StructureResponse("User Created successfully", "200", "", user, w)

}

func VerifyOtp(w http.ResponseWriter, r *http.Request) {
	email := r.Header.Get("email")
	otp := r.Header.Get("otp")

	storedOTP, err := db.GetOTP(email)
	if err != nil {
		StructureResponse("Error fetching OTP", "400", "true", "", w)
		return
	}
	if storedOTP != otp {
		StructureResponse("Invalid OTP", "400", "true", "", w)
		return
	}
	StructureResponse("Account Verified Successfully", "200", "false", "", w)
}
