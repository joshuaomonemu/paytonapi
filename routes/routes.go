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
	r.HandleFunc("/billers/categories", controller.GetBillersCategories).Methods("GET")
	r.HandleFunc("/billers/categories/{id}", controller.GetBillersCategoryId).Methods("GET")
	r.HandleFunc("/billers/item/{id}", controller.GetBillerItem).Methods("GET")
	r.HandleFunc("/billers/advice", controller.Advice).Methods("POST")
	r.HandleFunc("/billers/bill", controller.Bill).Methods("POST")
	r.HandleFunc("/airtime/load", controller.Airtime).Methods("POST")
	//r.HandleFunc("/billers/validate", controller.CustomerValidation).Methods("POST")

	log.Fatal(http.ListenAndServe(":"+port, r))
}
