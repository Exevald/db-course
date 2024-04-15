package service

import (
	"github.com/google/uuid"

	"orgchart/pkg/orgchart/app/model"
)

type BranchService interface {
	CreateBranch(city, address string) (uuid.UUID, error)
	UpdateBranch(branch model.Branch) error
	DeleteBranch(branchID uuid.UUID) error
	GetBranchInfo(branchID uuid.UUID) (model.Branch, error)
	GetBranchList() ([]model.Branch, error)
}

func NewBranchService(repository model.BranchRepository) BranchService {
	return &branchService{
		branchRepository: repository,
	}
}

type branchService struct {
	branchRepository model.BranchRepository
}

func (b *branchService) CreateBranch(city, address string) (uuid.UUID, error) {
	branchID := uuid.New()
	branch, err := model.NewBranch(branchID, city, address)
	if err != nil {
		return uuid.UUID{}, err
	}
	err = b.branchRepository.Store(branch)
	if err != nil {
		return uuid.UUID{}, err
	}
	return branchID, err
}

func (b *branchService) UpdateBranch(branch model.Branch) error {
	return b.branchRepository.Store(branch)
}

func (b *branchService) DeleteBranch(id uuid.UUID) error {
	return b.branchRepository.Delete(id)
}

func (b *branchService) GetBranchInfo(branchID uuid.UUID) (model.Branch, error) {
	return b.branchRepository.Find(branchID)
}

func (b *branchService) GetBranchList() ([]model.Branch, error) {
	branchList, err := b.branchRepository.ListBranches()
	if err != nil {
		return nil, err
	}
	return branchList, nil
}
