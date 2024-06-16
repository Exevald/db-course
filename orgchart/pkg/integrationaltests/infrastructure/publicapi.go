package infrastructure

import (
	"context"
	stderrors "errors"

	"github.com/pkg/errors"

	"orgchart/api/server/orgchartpublic"
	"orgchart/pkg/integrationaltests/tests"
)

func NewOrgchartPublicAPI(client orgchartpublic.ClientWithResponsesInterface) tests.OrgchartPublicAPI {
	return &orgchartPublicAPI{client: client}
}

type orgchartPublicAPI struct {
	client orgchartpublic.ClientWithResponsesInterface
}

func (api *orgchartPublicAPI) CreateBranch(request tests.CreateBranchRequestData) (tests.CreateBranchResponseData, error) {
	response, err := api.client.CreateBranchWithResponse(
		context.Background(),
		orgchartpublic.CreateBranchJSONRequestBody{
			Address: request.Address,
			City:    request.City,
		},
	)
	if err != nil {
		return tests.CreateBranchResponseData{}, err
	}
	if response.HTTPResponse.StatusCode != 200 {
		return tests.CreateBranchResponseData{}, stderrors.New(response.HTTPResponse.Status)
	}
	return tests.CreateBranchResponseData{BranchID: response.JSON200.BranchId}, err
}

func (api *orgchartPublicAPI) ListBranches() (tests.ListBranchesResponseData, error) {
	response, err := api.client.ListBranchesWithResponse(context.Background())
	if err != nil {
		return tests.ListBranchesResponseData{}, err
	}
	if response.HTTPResponse.StatusCode != 200 {
		return tests.ListBranchesResponseData{}, stderrors.New(response.HTTPResponse.Status)
	}
	var branches []tests.BranchPreview
	for _, responseBranch := range response.JSON200.Branches {
		branches = append(branches, tests.BranchPreview{
			Address:          responseBranch.Address,
			City:             responseBranch.City,
			CountOfEmployees: int(responseBranch.CountOfEmployees),
		})
	}
	return tests.ListBranchesResponseData{
		Branches: branches,
	}, nil
}

func (api *orgchartPublicAPI) GetBranchInfo(request tests.GetBranchInfoRequestData) (tests.GetBranchInfoResponseData, error) {
	response, err := api.client.GetBranchInfoWithResponse(
		context.Background(),
		request.BranchID,
	)
	if err != nil {
		return tests.GetBranchInfoResponseData{}, err
	}
	if response.HTTPResponse.StatusCode != 200 {
		return tests.GetBranchInfoResponseData{}, stderrors.New(response.HTTPResponse.Status)
	}
	var employees []tests.EmployeePreview
	for _, responseEmployee := range response.JSON200.Employees {
		employees = append(employees, tests.EmployeePreview{
			FirstName:  responseEmployee.FirstName,
			JobTitle:   responseEmployee.JobTitle,
			LastName:   responseEmployee.LastName,
			MiddleName: responseEmployee.MiddleName,
		})
	}
	return tests.GetBranchInfoResponseData{
		Address:   response.JSON200.Address,
		BranchId:  response.JSON200.BranchId,
		City:      response.JSON200.City,
		Employees: employees,
	}, nil
}

func (api *orgchartPublicAPI) DeleteBranch(request tests.DeleteBranchRequestData) error {
	response, err := api.client.DeleteBranchWithResponse(
		context.Background(),
		request.BranchID,
	)
	if err != nil {
		return err
	}
	if response.HTTPResponse.StatusCode != 200 {
		return stderrors.New(response.HTTPResponse.Status)
	}
	return nil
}

func (api *orgchartPublicAPI) UpdateBranchInfo(request tests.UpdateBranchRequestData) error {
	response, err := api.client.UpdateBranchInfoWithResponse(
		context.Background(),
		request.BranchID,
		orgchartpublic.UpdateBranchInfoJSONRequestBody{
			Address: request.Branch.Address,
			City:    request.Branch.City,
		},
	)
	if err != nil {
		return err
	}
	if response.HTTPResponse.StatusCode != 200 {
		return stderrors.New(response.HTTPResponse.Status)
	}
	return nil
}

func (api *orgchartPublicAPI) CreateEmployee(request tests.CreateEmployeeRequestData) (tests.CreateEmployeeResponseData, error) {
	response, err := api.client.CreateEmployeeWithResponse(
		context.Background(),
		orgchartpublic.CreateEmployeeJSONRequestBody{
			BirthDate:  request.BirthDate,
			BranchId:   request.BranchId,
			Comment:    request.Comment,
			Email:      request.Email,
			FirstName:  request.FirstName,
			Gender:     request.Gender,
			HireDate:   request.HireDate,
			JobTitle:   request.JobTitle,
			LastName:   request.LastName,
			MiddleName: request.MiddleName,
			Phone:      request.Phone,
			PhotoPath:  request.PhotoPath,
		},
	)
	if err != nil {
		return tests.CreateEmployeeResponseData{}, errors.WithStack(err)
	}
	if response.HTTPResponse.StatusCode != 200 {
		return tests.CreateEmployeeResponseData{}, stderrors.New(response.HTTPResponse.Status)
	}
	return tests.CreateEmployeeResponseData{
		EmployeeId: response.JSON200.EmployeeId,
	}, nil
}

func (api *orgchartPublicAPI) GetEmployeeInfo(request tests.GetEmployeeInfoRequestData) (tests.GetEmployeeInfoResponseData, error) {
	response, err := api.client.GetEmployeeInfoWithResponse(
		context.Background(),
		request.EmployeeId,
	)
	if err != nil {
		return tests.GetEmployeeInfoResponseData{}, err
	}
	if response.HTTPResponse.StatusCode != 200 {
		return tests.GetEmployeeInfoResponseData{}, stderrors.New(response.HTTPResponse.Status)
	}
	return tests.GetEmployeeInfoResponseData{
		BirthDate:  response.JSON200.BirthDate,
		BranchID:   response.JSON200.BranchId,
		Comment:    *response.JSON200.Comment,
		Email:      response.JSON200.Email,
		FirstName:  response.JSON200.FirstName,
		Gender:     response.JSON200.Gender,
		HireDate:   response.JSON200.HireDate,
		JobTitle:   response.JSON200.JobTitle,
		LastName:   response.JSON200.LastName,
		MiddleName: response.JSON200.MiddleName,
		Phone:      response.JSON200.Phone,
		PhotoPath:  response.JSON200.AvatarPath,
	}, nil
}

func (api *orgchartPublicAPI) DeleteEmployee(request tests.DeleteEmployeeRequestData) error {
	response, err := api.client.DeleteEmployeeWithResponse(
		context.Background(),
		request.EmployeeId,
	)
	if err != nil {
		return err
	}
	if response.HTTPResponse.StatusCode != 200 {
		return stderrors.New(response.HTTPResponse.Status)
	}
	return nil
}

func (api *orgchartPublicAPI) UpdateEmployeeInfo(request tests.UpdateEmployeeInfoRequestData) error {
	response, err := api.client.UpdateEmployeeInfoWithResponse(
		context.Background(),
		request.EmployeeId,
		orgchartpublic.UpdateEmployeeInfoJSONRequestBody{
			AvatarPath: request.Employee.PhotoPath,
			BirthDate:  request.Employee.BirthDate,
			BranchId:   request.Employee.BranchId,
			Comment:    &request.Employee.Comment,
			Email:      request.Employee.Email,
			EmployeeId: request.EmployeeId,
			FirstName:  request.Employee.FirstName,
			Gender:     request.Employee.Gender,
			HireDate:   request.Employee.HireDate,
			JobTitle:   request.Employee.JobTitle,
			LastName:   request.Employee.LastName,
			MiddleName: request.Employee.MiddleName,
			Phone:      request.Employee.Phone,
		},
	)
	if err != nil {
		return err
	}
	if response.HTTPResponse.StatusCode != 200 {
		return stderrors.New(response.HTTPResponse.Status)
	}
	return nil
}
