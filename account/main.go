package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type jsonResponse struct {
	Err     int         `json:"err"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// AccountMainHandler index.
func AccountMainHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Index handled\n")

	json.NewEncoder(w).Encode(jsonResponse{Err: 0, Message: "Success", Data: nil})
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
