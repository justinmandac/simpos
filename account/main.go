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

func writeJSONResponse(w http.ResponseWriter, errcode int, message string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.JSONResponse{Err: errcode, Message: message, Data: data})
}

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

	_, err := db.Exec(`INSERT INTO account (username, password) VALUES (?, ?);`, request.Username, request.Password)

	if err != nil {
		log.Printf("%s", err)
		writeJSONResponse(w, 1, "Error", nil)
		return
	}
	writeJSONResponse(w, 0, "Success", nil)
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

	if !dbAccount.AuthenticatePassword(account.Password) {
		log.Printf("%s\n", err)
		writeJSONResponse(w, 1, "Error", nil)
		return
	}
	log.Printf("%s authenticated\n", account.Username)
	writeJSONResponse(w, 1, "Error", nil)
}

// UpdateAccountHandler Updates the account password.
func UpdateAccountHandler(w http.ResponseWriter, r *http.Request) {
	var dbAccount models.Account
	var reqAccount models.AccountRequest
	vars := mux.Vars(r)

	if err := json.NewDecoder(r.Body).Decode(&reqAccount); err != nil {
		log.Printf("%s\n", err)
		writeJSONResponse(w, 2, "Invalid", err)
		return
	}

	reqAccount.Username = vars["username"]

	err := db.QueryRowx(`SELECT * FROM account WHERE username=?;`, reqAccount.Username).StructScan(&dbAccount)

	// Username not found
	if err != nil {
		writeJSONResponse(w, 1, "Error", err)
		return
	}

	targetAccount := models.GeneratePassword(reqAccount)
	_, err = db.Exec(`UPDATE account SET password = ? WHERE username = ?`, targetAccount.Password, targetAccount.Username)

	// Username not found
	if err != nil {
		writeJSONResponse(w, 1, "Error", err)
		return
	}
	writeJSONResponse(w, 0, "Success", nil)
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
	r.HandleFunc("/account/{username}", UpdateAccountHandler).Methods(http.MethodPut)
	r.HandleFunc("/account/{username}", DeleteAccountHandler).Methods(http.MethodDelete)

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(Config.SvcPort, nil))
}
