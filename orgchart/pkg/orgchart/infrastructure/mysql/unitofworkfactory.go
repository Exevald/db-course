package mysql

import (
	"context"

	"orgchart/pkg/orgchart/app/model"
	"orgchart/pkg/orgchart/app/service"
	"orgchart/pkg/orgchart/infrastructure/mysql/storage/branch"
	"orgchart/pkg/orgchart/infrastructure/mysql/storage/employee"
	"orgchart/pkg/orgchart/infrastructure/mysql/utils"
)

func NewUnitOfWorkFactory(factory utils.UnitOfWorkFactory) service.UnitOfWorkFactory {
	return &unitOfWorkFactory{factory: factory}
}

type unitOfWorkFactory struct {
	factory utils.UnitOfWorkFactory
}

func (u unitOfWorkFactory) NewUnitOfWork(ctx context.Context, lockName string) (service.UnitOfWork, error) {
	return &unitOfWork{}, nil
}

type unitOfWork struct {
	uow utils.UnitOfWork
}

func (u unitOfWork) BranchStorage() model.BranchStorage {
	return branch.NewBranchStorage(u.uow.ClientContext())
}

func (u unitOfWork) EmployeeStorage() model.EmployeeStorage {
	return employee.NewEmployeeStorage(u.uow.ClientContext())
}

func (u unitOfWork) Complete(err error) error {
	return u.uow.Complete(err)
}
