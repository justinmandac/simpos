package main

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httputil"
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
	requestDump, _ := httputil.DumpRequest(r, true)
	var account models.CreateAccountRequest

	log.Printf("%s\n", string(requestDump))

	if err := json.NewDecoder(r.Body).Decode(&account); err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.JSONResponse{Err: 0, Message: "Success", Data: nil})
}

// GetAccountHandler Authenticates the account provided in the request.
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
	r.HandleFunc("/account", CreateAccountHandler).Methods(http.MethodPost)

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(Config.SvcPort, nil))
}
