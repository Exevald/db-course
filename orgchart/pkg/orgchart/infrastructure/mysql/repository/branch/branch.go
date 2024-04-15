package branch

import (
	"database/sql"

	"github.com/google/uuid"
	"github.com/pkg/errors"

	"orgchart/pkg/orgchart/app/model"
	"orgchart/pkg/orgchart/common/mysql"
)

func NewBranchRepository(client mysql.Client) model.BranchRepository {
	return &branchRepository{client: client}
}

type branchRepository struct {
	client mysql.Client
}

func (b *branchRepository) Find(id uuid.UUID) (model.Branch, error) {
	const sqlQuery = `SELECT branch_id, city, address FROM branch WHERE branch_id=?`
	binaryID, err := id.MarshalBinary()
	if err != nil {
		return nil, err
	}

	row := b.client.QueryRow(sqlQuery, binaryID)
	var branch sqlBranch
	if err := row.Scan(&branch.branchID, &branch.city, &branch.address); err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.WithStack(model.ErrBranchNotFound)
		}
		return nil, errors.WithStack(err)
	}
	return sqlBranchToBranch(branch)
}

func (b *branchRepository) Store(branch model.Branch) error {
	const sqlQuery = `INSERT INTO branch 
    					(branch_id, city, address) 
					  VALUES (?, ?, ?)
					  ON DUPLICATE KEY UPDATE
					   	branch_id = VALUES(branch_id),
					   	city = VALUES(city),
						address = VALUES(address)
					 `
	binaryID, err := branch.ID().MarshalBinary()
	if err != nil {
		return errors.WithStack(err)
	}
	_, err = b.client.Exec(sqlQuery, binaryID, branch.City(), branch.Address())
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (b *branchRepository) Delete(id uuid.UUID) error {
	const sqlQuery = `DELETE FROM branch WHERE branch_id=?`
	binaryID, err := id.MarshalBinary()
	if err != nil {
		return err
	}

	_, err = b.client.Exec(sqlQuery, binaryID)
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (b *branchRepository) ListBranches() ([]model.Branch, error) {
	const sqlQuery = `SELECT branch_id, city, address FROM branch`
	var branchList []model.Branch

	rows, err := b.client.Query(sqlQuery)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	defer rows.Close()
	for rows.Next() {
		var branch sqlBranch
		if err = rows.Scan(&branch.branchID, &branch.city, &branch.address); err != nil {
			return nil, errors.WithStack(err)
		}
		modelBranch, err2 := sqlBranchToBranch(branch)
		if err2 != nil {
			return nil, err2
		}
		branchList = append(branchList, modelBranch)
	}
	if err = rows.Err(); err != nil {
		return nil, errors.WithStack(err)
	}
	return branchList, nil
}

type sqlBranch struct {
	branchID []byte `db:"branch_id"`
	city     string `db:"city"`
	address  string `db:"address"`
}

func sqlBranchToBranch(s sqlBranch) (model.Branch, error) {
	branchID, err := uuid.FromBytes(s.branchID)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return model.LoadBranch(
		branchID,
		s.city,
		s.address,
	), nil
}
