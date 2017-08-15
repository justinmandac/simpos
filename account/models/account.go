package models

import "golang.org/x/crypto/bcrypt"

// Account user account model.
type Account struct {
	ID         string `db:"Id"`
	Username   string `db:"username"`
	StoredHash string `db:"password"`
}

// AuthenticatePassword checks if the provided password matches
func (acct *Account) AuthenticatePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(acct.StoredHash), []byte(password))
	// nil means passwords match
	return err == nil
}

// AccountRequest request expected from the client when creating a new account.
type AccountRequest struct {
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
}

// GeneratePassword creates a password. acct is the request object containing the plaintext password
func GeneratePassword(acct AccountRequest) AccountRequest {
	var req AccountRequest
	password := []byte(acct.Password)
	passwordHash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	req = AccountRequest{Username: acct.Username, Password: string(passwordHash)}
	return req
}

// User user model.
type User struct {
	FirstName     string
	LastName      string
	ContactNumber string
	Email         string
}
