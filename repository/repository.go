package repository

import "database/sql"

type Repo struct {
	db *sql.DB
}

// index repo
type RepoInterface interface {
	EmployeeRepo
	BookRepo
}

// constructor function
func NewRepo(db *sql.DB) *Repo {
	return &Repo{db: db} // handle dependencies
}
