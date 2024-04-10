package model

import (
	"context"
	stderrors "errors"

	"github.com/pkg/errors"
)

var (
	ErrBranchNotFound = stderrors.New("branch not found")
	ErrInvalidCity    = stderrors.New("invalid city")
	ErrInvalidAddress = stderrors.New("invalid address")
)

type Branch interface {
	ID() uint64
	City() string
	Address() string
	Employees() []employee
}

func NewBranch(
	id uint64,
	city,
	address string,
) (Branch, error) {
	if city == "" {
		return branch{}, errors.WithStack(ErrInvalidCity)
	}
	if address == "" {
		return branch{}, errors.WithStack(ErrInvalidAddress)
	}

	return &branch{
		branchID:  id,
		city:      city,
		address:   address,
		employees: nil,
	}, nil
}

type branch struct {
	branchID  uint64
	city      string
	address   string
	employees []employee
}

func (b branch) ID() uint64 {
	return b.branchID
}

func (b branch) City() string {
	return b.city
}

func (b branch) Address() string {
	return b.address
}

func (b branch) Employees() []employee {
	return b.employees
}

type BranchStorage interface {
	Find(ctx context.Context, id uint64) (branch, error)
	Store(ctx context.Context, branch branch) error
	Delete(ctx context.Context, id uint64) error
}
