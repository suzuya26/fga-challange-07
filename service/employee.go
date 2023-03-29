package service

import "sesi_8/model"

type EmployeeService interface {
	CreateEmployee(in model.Employee) (res model.Employee, err error)
}

func (s *Service) CreateEmployee(in model.Employee) (res model.Employee, err error) {
	//call repo
	res, err = s.repo.CreateEmployee(in)
	if err != nil {
		return res, err
	}
	return res, nil
}
