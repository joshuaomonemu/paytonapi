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
	r.HandleFunc("/auth/token", controller.Auth).Methods("POST")

	r.HandleFunc("/airtime/load", controller.Airtime).Methods("POST")

	r.HandleFunc("/util/billers", controller.LoadBillers).Methods("GET")
	r.HandleFunc("/util/bill/pay", controller.PayBill).Methods("POST")

	r.HandleFunc("/giftcard/all", controller.GetGiftCards).Methods("GET")

	r.HandleFunc("/cable/dstv", controller.Dstv).Methods("GET")
	r.HandleFunc("/cable/dstv/verify/{id}", controller.DstvVerify).Methods("POST")
	r.HandleFunc("/cable/dstv/pay/{id}", controller.DstvPay).Methods("POST")

	r.HandleFunc("/cable/gotv", controller.Gotv).Methods("GET")
	r.HandleFunc("/cable/gotv/verify/{id}", controller.GotvVerify).Methods("POST")
	r.HandleFunc("/cable/gotv/pay/{id}", controller.GotvPay).Methods("POST")

	r.HandleFunc("/internet/smile", controller.Smile).Methods("GET")
	r.HandleFunc("/internet/smile/verify/{id}", controller.SmileVerify).Methods("POST")
	r.HandleFunc("/internet/smile/pay/{id}", controller.SmilePay).Methods("POST")

	r.HandleFunc("/data/all/{id}", controller.Data).Methods("GET")
	r.HandleFunc("/data/all/pay/{id}", controller.DataPay).Methods("POST")

	r.HandleFunc("/util/elect/verify/{id}", controller.ElectVerify).Methods("POST")
	r.HandleFunc("/util/elect/pay/{id}", controller.ElectPay1).Methods("POST")

	r.HandleFunc("/db/trans", controller.GetTrans).Methods("GET")

	r.HandleFunc("/user/transactions", controller.Transactions).Methods("GET")

	log.Fatal(http.ListenAndServe(":"+port, r))
}
