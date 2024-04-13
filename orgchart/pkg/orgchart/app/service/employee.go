package service

import (
	"context"
	"time"

	"orgchart/pkg/orgchart/app/model"
)

type EmployeeService interface {
	CreateEmployee(
		ctx context.Context,
		id uint64,
		firstName,
		lastName,
		middleName,
		jobTitle,
		email string,
		gender model.Gender,
		birthDate time.Time,
		hireDate time.Time,
		comment,
		avatarPath string,
	) (model.Employee, error)
	UpdateEmployee(ctx context.Context, employee model.Employee) error
	DeleteEmployee(ctx context.Context, id uint64) error
}

func NewEmployeeService(unitOfWorkFactory UnitOfWorkFactory) EmployeeService {
	return &employeeService{
		unitOfWorkService: unitOfWorkService{unitOfWorkFactory: unitOfWorkFactory},
	}
}

type employeeService struct {
	unitOfWorkService
}

func (e employeeService) CreateEmployee(
	ctx context.Context,
	id uint64,
	firstName,
	lastName,
	middleName,
	jobTitle,
	email string,
	gender model.Gender,
	birthDate,
	hireDate time.Time,
	comment,
	avatarPath string,
) (employee model.Employee, err error) {
	err = e.executeUnitOfWork(ctx, func(provider StorageProvider) error {
		employeeStorage := provider.EmployeeStorage()
		employee, err = model.NewEmployee(
			id,
			firstName,
			lastName,
			middleName,
			jobTitle,
			email,
			gender,
			birthDate,
			hireDate,
			comment,
			avatarPath,
		)
		if err != nil {
			return err
		}
		return employeeStorage.Store(ctx, employee)
	})
	return employee, nil
}

func (e employeeService) UpdateEmployee(ctx context.Context, employee model.Employee) error {
	err := e.executeUnitOfWork(ctx, func(provider StorageProvider) error {
		employeeStorage := provider.EmployeeStorage()
		return employeeStorage.Update(ctx, employee)
	})
	return err
}

func (e employeeService) DeleteEmployee(ctx context.Context, id uint64) error {
	err := e.executeUnitOfWork(ctx, func(provider StorageProvider) error {
		employeeStorage := provider.EmployeeStorage()
		return employeeStorage.Delete(ctx, id)
	})
	return err
}
