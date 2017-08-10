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
	config := &mysql.Config{
		User:   Config.DBUser,
		Passwd: Config.DBPassword,
		DBName: Config.DBName,
	}
	// username:password@protocol(address)/dbname?param=value
	db, err := sqlx.Connect("mysql", config.FormatDSN())
	log.Printf("Account Service started at %s", Config.SvcPort)

	err = db.Ping()

	if err != nil {
		log.Fatal(err)
	}

	r := mux.NewRouter()

	r.HandleFunc("/", AccountMainHandler)
	r.HandleFunc("/account", CreateAccountHandler).Methods(http.MethodPost)

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(Config.SvcPort, nil))
}
