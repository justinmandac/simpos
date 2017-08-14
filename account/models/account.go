package models

import "golang.org/x/crypto/bcrypt"

// Account user account model.
type Account struct {
	ID         string
	Username   string
	StoredHash string
}

// AuthenticatePassword checks if the provided password matches
func (acct *Account) AuthenticatePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(acct.StoredHash), []byte(password))
	// nil means passwords match
	return err == nil
}

// CreateAccountRequest request expected from the client when creating a new account.
type CreateAccountRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// GeneratePassword creates a password. acct is the request object containing the plaintext password
func GeneratePassword(acct CreateAccountRequest) *CreateAccountRequest {
	var req *CreateAccountRequest
	password := []byte(acct.Password)
	passwordHash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	req = &CreateAccountRequest{Username: acct.Username, Password: string(passwordHash)}
	return req
}

// User user model.
type User struct {
	FirstName     string
	LastName      string
	ContactNumber string
	Email         string
}
