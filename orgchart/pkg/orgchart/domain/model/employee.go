package model

import (
	"context"
	stderrors "errors"
	"github.com/pkg/errors"
	"time"
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
	EmployeeID() uint64
	FirstName() string
	LastName() string
	MiddleName() string
	JobTitle() string
	Email() string
	Gender() Gender
	BirthDate() time.Time
	HireDate() time.Time
}

func NewEmployee(
	id uint64,
	firstName,
	lastName,
	middleName string,
	jobTitle,
	email string,
	gender Gender,
	birthDate,
	hireDate time.Time,
	comment,
	avatarPath string,
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
	// TODO: Add check for valid gender, age and hire date

	return &employee{
		employeeID: id,
		firstName:  firstName,
		lastName:   lastName,
		middleName: middleName,
		jobTitle:   jobTitle,
		email:      email,
		gender:     gender,
		birthDate:  birthDate,
		hireDate:   hireDate,
		comment:    comment,
		avatarPath: avatarPath,
	}, nil
}

type employee struct {
	employeeID uint64
	firstName  string
	lastName   string
	middleName string
	jobTitle   string
	email      string
	gender     Gender
	birthDate  time.Time
	hireDate   time.Time
	comment    string
	avatarPath string
}

func (e employee) EmployeeID() uint64 {
	return e.employeeID
}

func (e employee) FirstName() string {
	return e.firstName
}

func (e employee) LastName() string {
	return e.lastName
}

func (e employee) MiddleName() string {
	return e.middleName
}

func (e employee) JobTitle() string {
	return e.jobTitle
}

func (e employee) Email() string {
	return e.email
}

func (e employee) Gender() Gender {
	return e.gender
}

func (e employee) BirthDate() time.Time {
	return e.birthDate
}

func (e employee) HireDate() time.Time {
	return e.hireDate
}

type EmployeeStorage interface {
	Find(ctx context.Context, id uint64) (employee, error)
	Store(ctx context.Context, employee employee) error
	Delete(ctx context.Context, id uint64) error
}
