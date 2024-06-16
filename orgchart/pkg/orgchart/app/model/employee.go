package model

import (
	stderrors "errors"
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type Gender uint8

const (
	Male   Gender = iota
	Female Gender = iota
)

var (
	ErrEmployeeNotFound = stderrors.New("employee not found")
	ErrInvalidName      = stderrors.New("invalid first name")
	ErrInvalidJobTitle  = stderrors.New("invalid job title")
	ErrInvalidEmail     = stderrors.New("invalid email")
	ErrInvalidGender    = stderrors.New("invalid gender")
	ErrInvalidAge       = stderrors.New("invalid age")
	ErrInvalidHireDate  = stderrors.New("invalid hire date")
)

type Employee interface {
	EmployeeID() uuid.UUID
	BranchID() uuid.UUID
	FirstName() string
	LastName() string
	MiddleName() string
	JobTitle() string
	Phone() string
	Email() string
	Gender() Gender
	BirthDate() time.Time
	HireDate() time.Time
	Comment() *string
	AvatarPath() *string
}

func NewEmployee(
	employeeID,
	branchID uuid.UUID,
	firstName,
	lastName,
	middleName,
	jobTitle,
	phone,
	email string,
	gender Gender,
	birthDate,
	hireDate time.Time,
	comment,
	avatarPath *string,
) (Employee, error) {
	if firstName == "" || lastName == "" || middleName == "" {
		return nil, errors.WithStack(ErrInvalidName)
	}
	if jobTitle == "" {
		return nil, errors.WithStack(ErrInvalidJobTitle)
	}
	if email == "" {
		return nil, errors.WithStack(ErrInvalidEmail)
	}
	if !isGenderValid(gender) {
		return nil, errors.WithStack(ErrInvalidGender)
	}
	if !isDateValid(birthDate) {
		return nil, errors.WithStack(ErrInvalidAge)
	}
	if !isDateValid(hireDate) || hireDate.Before(birthDate) {
		return nil, errors.WithStack(ErrInvalidHireDate)
	}
	return &employee{
		employeeID: employeeID,
		branchID:   branchID,
		firstName:  firstName,
		lastName:   lastName,
		middleName: middleName,
		jobTitle:   jobTitle,
		phone:      phone,
		email:      email,
		gender:     gender,
		birthDate:  birthDate,
		hireDate:   hireDate,
		comment:    comment,
		avatarPath: avatarPath,
	}, nil
}

func LoadEmployee(
	employeeID,
	branchID uuid.UUID,
	firstName,
	lastName,
	middleName,
	jobTitle,
	phone,
	email string,
	gender Gender,
	birthDate,
	hireDate time.Time,
	comment,
	avatarPath *string,
) Employee {
	return &employee{
		employeeID: employeeID,
		branchID:   branchID,
		firstName:  firstName,
		lastName:   lastName,
		middleName: middleName,
		jobTitle:   jobTitle,
		phone:      phone,
		email:      email,
		gender:     gender,
		birthDate:  birthDate,
		hireDate:   hireDate,
		comment:    comment,
		avatarPath: avatarPath,
	}
}

type employee struct {
	employeeID uuid.UUID
	branchID   uuid.UUID
	firstName  string
	lastName   string
	middleName string
	jobTitle   string
	phone      string
	email      string
	gender     Gender
	birthDate  time.Time
	hireDate   time.Time
	comment    *string
	avatarPath *string
}

func (e *employee) EmployeeID() uuid.UUID {
	return e.employeeID
}

func (e *employee) BranchID() uuid.UUID {
	return e.branchID
}

func (e *employee) FirstName() string {
	return e.firstName
}

func (e *employee) LastName() string {
	return e.lastName
}

func (e *employee) MiddleName() string {
	return e.middleName
}

func (e *employee) JobTitle() string {
	return e.jobTitle
}

func (e *employee) Email() string {
	return e.email
}

func (e *employee) Gender() Gender {
	return e.gender
}

func (e *employee) BirthDate() time.Time {
	return e.birthDate
}

func (e *employee) HireDate() time.Time {
	return e.hireDate
}

func (e *employee) Phone() string {
	return e.phone
}

func (e *employee) Comment() *string {
	return e.comment
}

func (e *employee) AvatarPath() *string {
	return e.avatarPath
}

type EmployeeRepository interface {
	Find(id uuid.UUID) (Employee, error)
	Store(employee Employee) error
	Delete(id uuid.UUID) error
	ListBranchEmployees(branchID uuid.UUID) ([]Employee, error)
	GetCountOfBranchEmployees(branchID uuid.UUID) (int, error)
}

func isGenderValid(gender Gender) bool {
	if gender == Male || gender == Female {
		return true
	}
	return false
}

func isDateValid(date time.Time) bool {
	now := time.Now()
	if date.Before(now) || date.IsZero() {
		return false
	}
	return true
}
