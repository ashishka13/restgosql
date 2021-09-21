package models

// Customer ....
type Customer struct {
	ID    int
	FName string `json:"fname"`
	LName string `json:"lname"`
}
