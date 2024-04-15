package publicapi

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"

	"orgchart/api/server/orgchartpublic"
	"orgchart/pkg/orgchart/app/model"
	"orgchart/pkg/orgchart/app/service"
)

type PublicAPI orgchartpublic.StrictServerInterface

func NewPublicAPI(
	branchService service.BranchService,
	employeeService service.EmployeeService,
) PublicAPI {
	return &publicAPI{
		branchService:   branchService,
		employeeService: employeeService,
	}
}

type publicAPI struct {
	branchService   service.BranchService
	employeeService service.EmployeeService
}

func (p publicAPI) CreateBranch(_ context.Context, request orgchartpublic.CreateBranchRequestObject) (orgchartpublic.CreateBranchResponseObject, error) {
	branchID, err := p.branchService.CreateBranch(request.Body.City, request.Body.Address)
	if err != nil {
		return nil, err
	}
	return orgchartpublic.CreateBranch200JSONResponse{BranchId: branchID.String()}, nil
}

func (p publicAPI) ListBranches(_ context.Context, _ orgchartpublic.ListBranchesRequestObject) (orgchartpublic.ListBranchesResponseObject, error) {
	branchList, err := p.branchService.GetBranchList()
	if err != nil {
		return nil, err
	}

	branches := make([]orgchartpublic.BranchPreview, 0, len(branchList))
	for _, branch := range branchList {
		branchEmployees, err2 := p.employeeService.GetBranchEmployees(branch.ID())
		if err2 != nil {
			return nil, err2
		}
		branches = append(branches, orgchartpublic.BranchPreview{
			Address:          branch.Address(),
			City:             branch.City(),
			CountOfEmployees: int64(len(branchEmployees)),
		})
	}
	return orgchartpublic.ListBranches200JSONResponse{Branches: branches}, err
}

func (p publicAPI) GetBranchInfo(_ context.Context, request orgchartpublic.GetBranchInfoRequestObject) (orgchartpublic.GetBranchInfoResponseObject, error) {
	branchID, err := uuid.Parse(request.BranchId)
	if err != nil {
		return nil, err
	}
	branch, err := p.branchService.GetBranchInfo(branchID)
	if err != nil {
		return nil, err
	}
	branchEmployees, err := p.employeeService.GetBranchEmployees(branch.ID())
	if err != nil {
		return nil, err
	}

	employees := make([]orgchartpublic.EmployeePreview, 0, len(branchEmployees))
	for _, employee := range branchEmployees {
		employees = append(employees, mapEmployeeModelToAPI(employee))
	}
	return orgchartpublic.GetBranchInfo200JSONResponse{
		City:      branch.City(),
		Address:   branch.Address(),
		Employees: employees,
	}, nil
}

func (p publicAPI) DeleteBranch(_ context.Context, request orgchartpublic.DeleteBranchRequestObject) (orgchartpublic.DeleteBranchResponseObject, error) {
	branchID, err := uuid.Parse(request.BranchId)
	if err != nil {
		return nil, err
	}
	branchEmployees, err := p.employeeService.GetBranchEmployees(branchID)
	if err != nil {
		return nil, err
	}
	if len(branchEmployees) > 0 {
		return nil, errors.WithStack(model.ErrBranchHasEmployees)
	}

	err = p.branchService.DeleteBranch(branchID)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (p publicAPI) UpdateBranchInfo(_ context.Context, request orgchartpublic.UpdateBranchInfoRequestObject) (orgchartpublic.UpdateBranchInfoResponseObject, error) {
	branchID, err := uuid.Parse(request.BranchId)
	if err != nil {
		return nil, err
	}

	branch, err := model.NewBranch(
		branchID,
		request.Body.City,
		request.Body.Address,
	)
	if err != nil {
		return nil, err
	}
	err = p.branchService.UpdateBranch(branch)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (p publicAPI) CreateEmployee(_ context.Context, request orgchartpublic.CreateEmployeeRequestObject) (orgchartpublic.CreateEmployeeResponseObject, error) {
	birthDate, err := time.Parse(time.DateOnly, request.Body.BirthDate)
	if err != nil {
		return nil, err
	}
	hireDate, err := time.Parse(time.DateOnly, request.Body.HireDate)
	if err != nil {
		return nil, err
	}
	branchID, err := uuid.Parse(request.Body.BranchId)
	if err != nil {
		return nil, err
	}
	employeeID, err := p.employeeService.CreateEmployee(
		branchID,
		request.Body.FirstName,
		request.Body.LastName,
		request.Body.MiddleName,
		request.Body.JobTitle,
		request.Body.Phone,
		request.Body.Email,
		model.Gender(request.Body.Gender),
		birthDate,
		hireDate,
		&request.Body.Comment,
		nil,
	)
	return orgchartpublic.CreateEmployee200JSONResponse{EmployeeId: employeeID.String()}, err
}

func (p publicAPI) GetEmployeeInfo(_ context.Context, request orgchartpublic.GetEmployeeInfoRequestObject) (orgchartpublic.GetEmployeeInfoResponseObject, error) {
	employeeID, err := uuid.Parse(request.EmployeeId)
	if err != nil {
		return nil, err
	}
	employee, err := p.employeeService.GetEmployeeInfo(employeeID)
	return orgchartpublic.GetEmployeeInfo200JSONResponse{
		FirstName:  employee.FirstName(),
		LastName:   employee.LastName(),
		MiddleName: employee.MiddleName(),
		JobTitle:   employee.JobTitle(),
		Phone:      employee.Phone(),
		Email:      employee.Email(),
		Gender:     int(employee.Gender()),
		BirthDate:  employee.BirthDate().String(),
		HireDate:   employee.HireDate().String(),
		Comment:    employee.Comment(),
	}, err
}

func (p publicAPI) DeleteEmployee(_ context.Context, request orgchartpublic.DeleteEmployeeRequestObject) (orgchartpublic.DeleteEmployeeResponseObject, error) {
	employeeID, err := uuid.Parse(request.EmployeeId)
	if err != nil {
		return nil, err
	}
	err = p.employeeService.DeleteEmployee(employeeID)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (p publicAPI) UpdateEmployeeInfo(_ context.Context, request orgchartpublic.UpdateEmployeeInfoRequestObject) (orgchartpublic.UpdateEmployeeInfoResponseObject, error) {
	employeeID, err := uuid.Parse(request.EmployeeId)
	if err != nil {
		return nil, err
	}
	branchID, err := uuid.Parse(request.Body.BranchId)
	if err != nil {
		return nil, err
	}
	birthDate, err := time.Parse(time.DateOnly, request.Body.BirthDate)
	if err != nil {
		return nil, err
	}
	hireDate, err := time.Parse(time.DateOnly, request.Body.HireDate)
	if err != nil {
		return nil, err
	}

	employee, err := model.NewEmployee(
		employeeID,
		branchID,
		request.Body.FirstName,
		request.Body.LastName,
		request.Body.MiddleName,
		request.Body.JobTitle,
		request.Body.Phone,
		request.Body.Email,
		model.Gender(request.Body.Gender),
		birthDate,
		hireDate,
		request.Body.Comment,
		request.Body.AvatarPath,
	)
	if err != nil {
		return nil, err
	}
	err = p.employeeService.UpdateEmployee(employee)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func mapEmployeeModelToAPI(employee model.Employee) orgchartpublic.EmployeePreview {
	return orgchartpublic.EmployeePreview{
		FirstName:  employee.FirstName(),
		LastName:   employee.LastName(),
		MiddleName: employee.MiddleName(),
		JobTitle:   employee.JobTitle(),
	}
}
