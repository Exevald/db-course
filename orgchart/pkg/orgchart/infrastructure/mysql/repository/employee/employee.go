package employee

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"

	"orgchart/pkg/orgchart/app/model"
	"orgchart/pkg/orgchart/common/mysql"
)

func NewEmployeeRepository(client mysql.Client) model.EmployeeRepository {
	return &employeeRepository{client: client}
}

type employeeRepository struct {
	client mysql.Client
}

func (e employeeRepository) Find(id uuid.UUID) (model.Employee, error) {
	const sqlQuery = `SELECT
    					employee_id,
    					branch_id,
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
					  WHERE employee_id = ?
					 `
	binaryID, err := id.MarshalBinary()
	if err != nil {
		return nil, err
	}

	row := e.client.QueryRow(sqlQuery, binaryID)
	var employee sqlEmployee
	if err := row.Scan(
		&employee.employeeID,
		&employee.branchID,
		&employee.firstName,
		&employee.lastName,
		&employee.middleName,
		&employee.jobTitle,
		&employee.phone,
		&employee.email,
		&employee.gender,
		&employee.birthDate,
		&employee.hireDate,
		&employee.comment,
		&employee.avatarPath,
	); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.WithStack(model.ErrEmployeeNotFound)
		}
		return nil, errors.WithStack(err)
	}
	return sqlEmployeeToEmployee(employee)
}

func (e employeeRepository) Store(employee model.Employee) error {
	const sqlQuery = `INSERT INTO employee(
                     	employee_id,
                     	branch_id,
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
					 ) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
					 ON DUPLICATE KEY UPDATE
						employee_id = VALUES(employee_id),
						branch_id = VALUES(branch_id),
						first_name = VALUES(first_name),
                     	last_name = VALUES(last_name),
                     	middle_name = VALUES(middle_name),
                     	job_title = VALUES(job_title),
                     	phone = VALUES(phone),
                     	email = VALUES(email),
                     	gender = VALUES(gender),
                     	birth_date = VALUES(birth_date),
                     	hire_date = VALUES(hire_date),
                     	comment = VALUES(comment),
                     	avatar_path = VALUES(avatar_path)
					  `
	binaryEmployeeID, err := employee.EmployeeID().MarshalBinary()
	if err != nil {
		return errors.WithStack(err)
	}
	binaryBranchID, err := employee.BranchID().MarshalBinary()
	if err != nil {
		return errors.WithStack(err)
	}
	_, err = e.client.Exec(
		sqlQuery,
		binaryEmployeeID,
		binaryBranchID,
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
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (e employeeRepository) Delete(id uuid.UUID) error {
	const sqlQuery = `DELETE FROM employee WHERE employee_id = ?`
	binaryID, err := id.MarshalBinary()
	if err != nil {
		return errors.WithStack(err)
	}
	_, err = e.client.Exec(sqlQuery, binaryID)
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (e employeeRepository) ListBranchEmployees(branchID uuid.UUID) ([]model.Employee, error) {
	const sqlQuery = `SELECT
    					e.employee_id,
    					e.branch_id,
    					e.first_name,
    					e.last_name,
    					e.middle_name,
    					e.job_title,
    					e.phone,
    					e.email,
    					e.gender,
    					e.birth_date,
    					e.hire_date,
    					e.comment,
    					e.avatar_path
					  FROM employee e
					  INNER JOIN branch b ON b.branch_id = e.branch_id
					  WHERE b.branch_id = ?
					 `
	binaryID, err := branchID.MarshalBinary()
	if err != nil {
		return nil, err
	}
	var employeeList []model.Employee
	rows, err := e.client.Query(sqlQuery, binaryID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.WithStack(model.ErrBranchNotFound)
		}
		return nil, errors.WithStack(err)
	}
	defer rows.Close()
	for rows.Next() {
		var employee sqlEmployee
		if err = rows.Scan(
			&employee.employeeID,
			&employee.branchID,
			&employee.firstName,
			&employee.lastName,
			&employee.middleName,
			&employee.jobTitle,
			&employee.phone,
			&employee.email,
			&employee.gender,
			&employee.birthDate,
			&employee.hireDate,
			&employee.comment,
			&employee.avatarPath,
		); err != nil {
			return nil, errors.WithStack(err)
		}
		modelEmployee, err2 := sqlEmployeeToEmployee(employee)
		if err2 != nil {
			return nil, err2
		}
		employeeList = append(employeeList, modelEmployee)
	}
	if err = rows.Err(); err != nil {
		return nil, errors.WithStack(err)
	}
	return employeeList, nil
}

type sqlEmployee struct {
	employeeID []byte    `db:"employee_id"`
	branchID   []byte    `db:"branch_id"`
	firstName  string    `db:"first_name"`
	lastName   string    `db:"last_name"`
	middleName string    `db:"middle_name"`
	jobTitle   string    `db:"job_title"`
	phone      string    `db:"phone"`
	email      string    `db:"email"`
	gender     uint8     `db:"gender"`
	birthDate  time.Time `db:"birth_date"`
	hireDate   time.Time `db:"hire_date"`
	comment    *string   `db:"comment"`
	avatarPath *string   `db:"avatar_path"`
}

func sqlEmployeeToEmployee(s sqlEmployee) (model.Employee, error) {
	employeeID, err := uuid.FromBytes(s.employeeID)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	branchID, err := uuid.FromBytes(s.branchID)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return model.LoadEmployee(
		employeeID,
		branchID,
		s.firstName,
		s.lastName,
		s.middleName,
		s.jobTitle,
		s.phone,
		s.email,
		model.Gender(s.gender),
		s.birthDate,
		s.hireDate,
		s.comment,
		s.avatarPath,
	), nil
}
