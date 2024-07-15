package routes

import (
	"app/controller"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

// Starting route points
var r = mux.NewRouter()

func Routes() {
	//Starting Server and running on port 2020
	//port = os.Getenv("PORT") // Get the port from the environment

	// Api for User
	//r.HandleFunc("/controller/user/get{id}", controller.GetUser).Methods("GET")
	port := os.Getenv("PORT")
	if port == "" {
		port = "2020" // Set a default port for development purposes (can be removed for production)
	}

	//Endpoints and Route points for users
	r.HandleFunc("/airtime/load", controller.Airtime).Methods("POST")
	r.HandleFunc("/util/billers", controller.LoadBillers).Methods("POST")
	r.HandleFunc("/util/bill/pay", controller.PayBill).Methods("POST")
	//r.HandleFunc("/billers/validate", controller.CustomerValidation).Methods("POST")

	log.Fatal(http.ListenAndServe(":"+port, r))
}
