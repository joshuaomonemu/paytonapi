package controller

import (
	"app/models"
	Structs "app/struct"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func GetBillersCategories(w http.ResponseWriter, r *http.Request) {
	resp, err := models.GetBillersCategories()
	if err != nil {
		io.WriteString(w, fmt.Sprintln(err))
	}

	var response *Structs.BillerCategoriesResponse
	err = json.Unmarshal(resp, &response)
	if err != nil {
		io.WriteString(w, err.Error())
		return
	}

	jsn, err := json.Marshal(response)
	if err != nil {
		log.Fatal(err)
	}
	io.WriteString(w, string(jsn))

}

func GetBillersCategoryId(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	resp, err := models.GetBillersCategoryId(id)
	if err != nil {
		io.WriteString(w, fmt.Sprintln(err))
	}

	var response *Structs.BillersIdResponse
	err = json.Unmarshal(resp, &response)
	if err != nil {
		io.WriteString(w, err.Error())
		return
	}

	jsn, err := json.Marshal(response)
	if err != nil {
		log.Fatal(err)
	}
	io.WriteString(w, string(jsn))

}

func GetBillerItem(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	resp, err := models.GetBillerItem(id)
	if err != nil {
		io.WriteString(w, fmt.Sprintln(err))
	}

	// var response BillersIdResponse
	// err = json.Unmarshal(resp, &response)
	// if err != nil {
	// 	io.WriteString(w, err.Error())
	// 	return
	// }

	// jsn, err := json.Marshal(response)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	io.WriteString(w, string(resp))

}

func CustomerValidation(w http.ResponseWriter, r *http.Request) {
	// params := mux.Vars(r)
	// id := params["id"]

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		io.WriteString(w, "Request format error")
	}

	// var item *Structs.PaymentTransaction

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		log.Fatal("Error marshalling data to JSON:", err)
		return
	}

	resp, err := models.Validate(jsonData)
	if err != nil {
		io.WriteString(w, fmt.Sprintln(err))
	}
	// jsn, err := json.Marshal(response)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	io.WriteString(w, string(resp))

}

func Advice(w http.ResponseWriter, r *http.Request) {
	// params := mux.Vars(r)
	// id := params["id"]

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		io.WriteString(w, "Request format error")
	}
	fmt.Println(string(reqBody))

	// var item *Structs.PaymentTransaction

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		log.Fatal("Error marshalling data to JSON:", err)
		return
	}

	resp, err := models.Advice(jsonData)
	if err != nil {
		io.WriteString(w, fmt.Sprintln(err))
	}
	// jsn, err := json.Marshal(response)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	io.WriteString(w, string(resp))

}

func Bill(w http.ResponseWriter, r *http.Request) {
	models.Bill()
}
