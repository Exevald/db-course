package branch

import (
	"context"

	"orgchart/pkg/orgchart/app/model"
	"orgchart/pkg/orgchart/infrastructure/mysql/utils"
)

func NewBranchStorage(client utils.ClientContext) model.BranchStorage {
	return &branchStorage{client: client}
}

type branchStorage struct {
	client utils.ClientContext
}

func (b *branchStorage) Find(ctx context.Context, id uint64) (model.Branch, error) {
	const sqlQuery = `SELECT 
    				    city,
    				    address			    
					  FROM branch 
					  WHERE id=?
					  `

}

func (b *branchStorage) Store(ctx context.Context, branch model.Branch) error {
	panic("implement me")
}

func (b *branchStorage) Update(ctx context.Context, branch model.Branch) error {
	panic("implement me")
}

func (b *branchStorage) Delete(ctx context.Context, id uint64) error {
	panic("implement me")
}

var sqlxBranch struct {
	city      string `db:"city"`
	address   string `db:"address"`
	employees []model.Employee
}
