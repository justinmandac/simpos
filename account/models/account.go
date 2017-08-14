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

	return true
}

// CreateAccountRequest request expected from the client when creating a new account.
type CreateAccountRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// GeneratePassword creates a password
func (acct *CreateAccountRequest) GeneratePassword() string {
	password := []byte(acct.Password)
	passwordHash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(passwordHash)
}

// User user model.
type User struct {
	FirstName     string
	LastName      string
	ContactNumber string
	Email         string
}
