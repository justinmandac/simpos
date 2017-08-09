package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// JSONResponse generic model for JSON responses.
type JSONResponse struct {
	Err     int         `json:"err"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// Account user account model
type Account struct {
	ID   string
	Name string
}

// AccountMainHandler index.
func AccountMainHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Index handled\n")
	data := JSONResponse{Err: 0, Message: "Success", Data: nil}
	json.NewEncoder(w).Encode(data)
}

// CreateAccountHandler Create Account
func CreateAccountHandler(w http.ResponseWriter, r *http.Request) {

}

// GetAccountHandler Get Account
func GetAccountHandler(w http.ResponseWriter, r *http.Request) {

}

// UpdateAccountHandler Updates account information.
func UpdateAccountHandler(w http.ResponseWriter, r *http.Request) {

}

// DeleteAccountHandler Delete an account/render account inactive.
func DeleteAccountHandler(w http.ResponseWriter, r *http.Request) {

}

func main() {
	log.Printf("Account Service started at %s", Config.SvcPort)
	r := mux.NewRouter()
	r.HandleFunc("/", AccountMainHandler)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(Config.SvcPort, nil))
}
