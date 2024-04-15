package service

import (
	"time"

	"github.com/google/uuid"

	"orgchart/pkg/orgchart/app/model"
)

type EmployeeService interface {
	CreateEmployee(
		branchID uuid.UUID,
		firstName,
		lastName,
		middleName,
		jobTitle,
		phone,
		email string,
		gender model.Gender,
		birthDate time.Time,
		hireDate time.Time,
		comment,
		avatarPath *string,
	) (uuid.UUID, error)
	UpdateEmployee(employee model.Employee) error
	DeleteEmployee(id uuid.UUID) error
	GetEmployeeInfo(employeeID uuid.UUID) (model.Employee, error)
	GetBranchEmployees(branchID uuid.UUID) ([]model.Employee, error)
}

func NewEmployeeService(repository model.EmployeeRepository) EmployeeService {
	return &employeeService{employeeRepository: repository}
}

type employeeService struct {
	employeeRepository model.EmployeeRepository
}

func (e *employeeService) CreateEmployee(
	branchID uuid.UUID,
	firstName,
	lastName,
	middleName,
	jobTitle,
	phone,
	email string,
	gender model.Gender,
	birthDate,
	hireDate time.Time,
	comment,
	avatarPath *string,
) (uuid.UUID, error) {
	employeeID := uuid.New()
	employee, err := model.NewEmployee(
		employeeID,
		branchID,
		firstName,
		lastName,
		middleName,
		jobTitle,
		phone,
		email,
		gender,
		birthDate,
		hireDate,
		comment,
		avatarPath,
	)
	if err != nil {
		return uuid.UUID{}, err
	}
	err = e.employeeRepository.Store(employee)
	return employeeID, err
}

func (e *employeeService) UpdateEmployee(employee model.Employee) error {
	return e.employeeRepository.Store(employee)
}

func (e *employeeService) DeleteEmployee(employeeID uuid.UUID) error {
	return e.employeeRepository.Delete(employeeID)
}

func (e *employeeService) GetEmployeeInfo(employeeID uuid.UUID) (model.Employee, error) {
	return e.employeeRepository.Find(employeeID)
}

func (e *employeeService) GetBranchEmployees(branchID uuid.UUID) ([]model.Employee, error) {
	return e.employeeRepository.ListBranchEmployees(branchID)
}
