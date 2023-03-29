package model

type Employee struct {
	ID       int    `json:"id" db:"id"`
	Fullname string `json:"full_name" db:"full_name"`
	Email    string `json:"email" db:"email"`
	Age      int    `json:"age" db:"age"`
	Division string `json:"division" db:"division"`
}
