package model

import (
	stderrors "errors"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

var (
	ErrBranchNotFound     = stderrors.New("branch not found")
	ErrInvalidCity        = stderrors.New("invalid city")
	ErrInvalidAddress     = stderrors.New("invalid address")
	ErrBranchHasEmployees = stderrors.New("branch has employees")
)

type Branch interface {
	ID() uuid.UUID
	City() string
	Address() string
}

func NewBranch(
	id uuid.UUID,
	city,
	address string,
) (Branch, error) {
	if city == "" {
		return &branch{}, errors.WithStack(ErrInvalidCity)
	}
	if address == "" {
		return &branch{}, errors.WithStack(ErrInvalidAddress)
	}

	return &branch{
		branchID: id,
		city:     city,
		address:  address,
	}, nil
}

func LoadBranch(
	id uuid.UUID,
	city,
	address string,
) Branch {
	return &branch{
		branchID: id,
		city:     city,
		address:  address,
	}
}

type branch struct {
	branchID uuid.UUID
	city     string
	address  string
}

func (b *branch) ID() uuid.UUID {
	return b.branchID
}

func (b *branch) City() string {
	return b.city
}

func (b *branch) Address() string {
	return b.address
}

type BranchRepository interface {
	Find(id uuid.UUID) (Branch, error)
	Store(branch Branch) error
	Delete(id uuid.UUID) error
	ListBranches() ([]Branch, error)
}
