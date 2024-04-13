package service

import (
	"context"
	"orgchart/pkg/orgchart/app/model"
)

type BranchService interface {
	CreateBranch(ctx context.Context, id uint64, city, address string) (model.Branch, error)
	UpdateBranch(ctx context.Context, branch model.Branch) error
	DeleteBranch(ctx context.Context, id uint64) error
}

func NewBranchService(unitOfWorkFactory UnitOfWorkFactory) BranchService {
	return &branchService{
		unitOfWorkService: unitOfWorkService{unitOfWorkFactory: unitOfWorkFactory},
	}
}

type branchService struct {
	unitOfWorkService
}

func (b branchService) CreateBranch(ctx context.Context, id uint64, city, address string) (branch model.Branch, err error) {
	err = b.executeUnitOfWork(ctx, func(provider StorageProvider) error {
		branchStorage := provider.BranchStorage()
		branch, err = model.NewBranch(id, city, address)
		if err != nil {
			return err
		}
		err = branchStorage.Store(ctx, branch)
		if err != nil {
			return err
		}
		return nil
	})
	return branch, err
}

func (b branchService) UpdateBranch(ctx context.Context, branch model.Branch) error {
	err := b.executeUnitOfWork(ctx, func(provider StorageProvider) error {
		branchStorage := provider.BranchStorage()
		return branchStorage.Update(ctx, branch)
	})
	return err
}

func (b branchService) DeleteBranch(ctx context.Context, id uint64) error {
	err := b.executeUnitOfWork(ctx, func(provider StorageProvider) error {
		branchStorage := provider.BranchStorage()
		return branchStorage.Delete(ctx, id)
	})
	return err
}
