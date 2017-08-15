package main

import (
	"encoding/json"
	"log"
	"net/http"
	"net/http/httputil"
	"simpos/account/models"

	"github.com/go-sql-driver/mysql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

// AccountMainHandler index.
func AccountMainHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Index handled\n")
	data := models.JSONResponse{Err: 0, Message: "Success", Data: nil}
	json.NewEncoder(w).Encode(data)
}

// CreateAccountHandler Create Account
func CreateAccountHandler(w http.ResponseWriter, r *http.Request) {
	requestDump, _ := httputil.DumpRequest(r, true)
	var account models.AccountRequest

	log.Printf("%s\n", string(requestDump))

	if err := json.NewDecoder(r.Body).Decode(&account); err != nil {
		panic(err)
	}

	request := models.GeneratePassword(account)

	w.Header().Set("Content-Type", "application/json")

	_, err := db.Exec(`INSERT INTO account (username, password) VALUES (?, ?);`, request.Username, request.Password)

	if err != nil {
		log.Printf("%s", err)
		json.NewEncoder(w).Encode(models.JSONResponse{Err: 1, Message: "Error", Data: nil})
		return
	}

	json.NewEncoder(w).Encode(models.JSONResponse{Err: 0, Message: "Success", Data: request})
}

// GetAccountHandler Authenticates the account provided in the request.
func GetAccountHandler(w http.ResponseWriter, r *http.Request) {
	requestDump, _ := httputil.DumpRequest(r, true)
	var account models.AccountRequest
	var dbAccount models.Account

	log.Printf("%s\n", string(requestDump))

	if err := json.NewDecoder(r.Body).Decode(&account); err != nil {
		panic(err)
	}

	err := db.QueryRowx(`SELECT * FROM account WHERE username=?`, account.Username).StructScan(&dbAccount)

	w.Header().Set("Content-Type", "application/json")

	if !dbAccount.AuthenticatePassword(account.Password) {
		log.Printf("%s\n", err)
		json.NewEncoder(w).Encode(models.JSONResponse{Err: 1, Message: "Error", Data: nil})
		return
	}
	log.Printf("%s authenticated\n", account.Username)
	json.NewEncoder(w).Encode(models.JSONResponse{Err: 0, Message: "Success", Data: nil})
}

// UpdateAccountHandler Updates account information.
func UpdateAccountHandler(w http.ResponseWriter, r *http.Request) {

}

// DeleteAccountHandler Delete an account/render account inactive.
func DeleteAccountHandler(w http.ResponseWriter, r *http.Request) {

}

func main() {
	config := &mysql.Config{
		User:   Config.DBUser,
		Passwd: Config.DBPassword,
		DBName: Config.DBName,
	}
	// username:password@protocol(address)/dbname?param=value
	db = sqlx.MustConnect("mysql", config.FormatDSN())
	log.Printf("Account Service started at %s", Config.SvcPort)

	r := mux.NewRouter()

	r.HandleFunc("/", AccountMainHandler)
	r.HandleFunc("/account", CreateAccountHandler).Methods(http.MethodPost)
	r.HandleFunc("/account/auth", GetAccountHandler).Methods(http.MethodPost)

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(Config.SvcPort, nil))
}
