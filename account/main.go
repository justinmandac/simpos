package main

import (
	"encoding/json"
	"log"
	"net/http"
	"simpos/account/models"

	"github.com/gorilla/mux"
)

// AccountMainHandler index.
func AccountMainHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Index handled\n")
	data := models.JSONResponse{Err: 0, Message: "Success", Data: nil}
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
