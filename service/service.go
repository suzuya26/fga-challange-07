package service

import "sesi_8/repository"

type Service struct {
	repo repository.RepoInterface
}

// index service
type ServiceInterface interface {
	EmployeeService
	BookService
}

func NewService(repo repository.RepoInterface) ServiceInterface {
	return &Service{repo: repo}
}
