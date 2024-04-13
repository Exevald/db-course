package employee

import (
	"context"
	"database/sql"
	"time"

	"github.com/pkg/errors"

	"orgchart/pkg/orgchart/app/model"
	"orgchart/pkg/orgchart/infrastructure/mysql/utils"
)

func NewEmployeeStorage(client utils.ClientContext) model.EmployeeStorage {
	return &employeeStorage{client: client}
}

type employeeStorage struct {
	client utils.ClientContext
}

func (s *employeeStorage) Find(ctx context.Context, id uint64) (model.Employee, error) {
	const sqlQuery = `SELECT
    					first_name,
    					last_name,
    					middle_name,
    					job_title,
    					phone,
    					email,
    					gender,
    					birth_date,
    					hire_date,
    					comment,
    					avatar_path
					  FROM employee
					  WHERE id = ?
					 `
	var employee sqlxEmployee
	err := s.client.GetContext(ctx, &employee, sqlQuery, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.WithStack(model.ErrEmployeeNotFound)
		}
		return nil, errors.WithStack(err)
	}
}

func (s *employeeStorage) Store(ctx context.Context, employee model.Employee) error {
	const sqlQuery = `INSERT INTO employee
    				  	  (
    				  	   first_name,
    				  	   last_name,
    				  	   middle_name,
    					   job_title,
    					   phone,
    					   email,
    					   gender,
    					   birth_date,
    					   hire_date,
    					   comment,
    					   avatar_path
    				  	  )
					  VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	_, err := s.client.ExecContext(ctx, sqlQuery,
		employee.FirstName(),
		employee.LastName(),
		employee.MiddleName(),
		employee.JobTitle(),
		employee.Phone(),
		employee.Email(),
		employee.Gender(),
		employee.BirthDate(),
		employee.HireDate(),
		employee.Comment(),
		employee.AvatarPath(),
	)
	return errors.WithStack(err)
}

func (s *employeeStorage) Update(ctx context.Context, employee model.Employee) error {
	const sqlQuery = `UPDATE employee
    				  SET
    				  	   first_name = ?,
    				  	   last_name = ?,
    				  	   middle_name = ?,
    					   job_title = ?,
    					   phone = ?,
    					   email = ?,
    					   gender = ?,
    					   birth_date = ?,
    					   hire_date = ?,
    					   comment = ?,
    					   avatar_path = ?
					  WHERE id = ?
					 `
	_, err := s.client.ExecContext(ctx, sqlQuery,
		employee.FirstName(),
		employee.LastName(),
		employee.MiddleName(),
		employee.JobTitle(),
		employee.Phone(),
		employee.Email(),
		employee.Gender(),
		employee.BirthDate(),
		employee.HireDate(),
		employee.Comment(),
		employee.AvatarPath(),
		employee.EmployeeID(),
	)
	return errors.WithStack(err)
}

func (s *employeeStorage) Delete(ctx context.Context, id uint64) error {
	const sqlQuery = `DELETE FROM employee WHERE id = ?`
	_, err := s.client.ExecContext(ctx, sqlQuery, id)
	return errors.WithStack(err)
}

//func (s *employeeStorage) FindBranchEmployees(ctx context.Context, branchId uint64) ([]model.Employee, error) {
//	const sqlQuery = `SELECT
//    					first_name,
//    					last_name,
//    					middle_name,
//    					job_title,
//    					phone,
//    					email,
//    					gender,
//    					birth_date,
//    					hire_date,
//    					comment,
//    					avatar_path
//					  FROM employee
//					  INNER JOIN branch b ON b.id = ?
//					 `
//	var employees []sqlxEmployee
//	employee := Em
//}

type sqlxEmployee struct {
	EmployeeID uint64       `db:"id"`
	FirstName  string       `db:"first_name"`
	LastName   string       `db:"last_name"`
	MiddleName string       `db:"middle_name"`
	JobTitle   string       `db:"job_title"`
	Phone      string       `db:"phone"`
	Email      string       `db:"email"`
	Gender     model.Gender `db:"gender"`
	BirthDate  time.Time    `db:"birth_date"`
	HireDate   time.Time    `db:"hire_date"`
	Comment    string       `db:"comment"`
	AvatarPath string       `db:"avatar_path"`
}
