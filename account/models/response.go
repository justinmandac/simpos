package models

// JSONResponse generic model for JSON responses.
type JSONResponse struct {
	Err     int         `json:"err"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
