package repository

import (
	"sesi_8/model"
	"sesi_8/repository/query"
)

// interface employee
type EmployeeRepo interface {
	CreateEmployee(in model.Employee) (res model.Employee, err error)
}

func (r Repo) CreateEmployee(in model.Employee) (res model.Employee, err error) {
	err = r.db.QueryRow(
		query.AddEmployee,
		in.Fullname,
		in.Email,
		in.Age,
		in.Division,
	).Scan(
		&res.ID,
		&res.Fullname,
		&res.Email,
		&res.Age,
		&res.Division,
	)
	if err != nil {
		return res, err
	}
	return res, err
}
