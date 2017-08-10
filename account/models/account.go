package models

// Account user account model.
type Account struct {
	ID         string
	Username   string
	StoredHash string
}

// CreateAccountRequest request expected from the client when creating a new account.
type CreateAccountRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// User user model.
type User struct {
	FirstName     string
	LastName      string
	ContactNumber string
	Email         string
}
