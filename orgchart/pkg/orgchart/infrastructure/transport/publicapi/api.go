package publicapi

import (
	"context"

	"orgchart/api/server/orgchartpublic"
)

type PublicAPI orgchartpublic.StrictServerInterface

func NewPublicAPI() PublicAPI {
	return &publicAPi{}
}

type publicAPi struct {
}

func (p publicAPi) CreateBranch(ctx context.Context, request orgchartpublic.CreateBranchRequestObject) (orgchartpublic.CreateBranchResponseObject, error) {
	panic("implement me")
}

func (p publicAPi) ListBranches(ctx context.Context, request orgchartpublic.ListBranchesRequestObject) (orgchartpublic.ListBranchesResponseObject, error) {
	panic("implement me")
}

func (p publicAPi) GetBranchInfo(ctx context.Context, request orgchartpublic.GetBranchInfoRequestObject) (orgchartpublic.GetBranchInfoResponseObject, error) {
	panic("implement me")
}

func (p publicAPi) DeleteBranch(ctx context.Context, request orgchartpublic.DeleteBranchRequestObject) (orgchartpublic.DeleteBranchResponseObject, error) {
	//TODO implement me
	panic("implement me")
}

func (p publicAPi) UpdateBranchInfo(ctx context.Context, request orgchartpublic.UpdateBranchInfoRequestObject) (orgchartpublic.UpdateBranchInfoResponseObject, error) {
	//TODO implement me
	panic("implement me")
}

func (p publicAPi) CreateEmployee(ctx context.Context, request orgchartpublic.CreateEmployeeRequestObject) (orgchartpublic.CreateEmployeeResponseObject, error) {
	//TODO implement me
	panic("implement me")
}

func (p publicAPi) GetEmployeeInfo(ctx context.Context, request orgchartpublic.GetEmployeeInfoRequestObject) (orgchartpublic.GetEmployeeInfoResponseObject, error) {
	//TODO implement me
	panic("implement me")
}

func (p publicAPi) DeleteEmployee(ctx context.Context, request orgchartpublic.DeleteEmployeeRequestObject) (orgchartpublic.DeleteEmployeeResponseObject, error) {
	//TODO implement me
	panic("implement me")
}

func (p publicAPi) UpdateEmployeeInfo(ctx context.Context, request orgchartpublic.UpdateEmployeeInfoRequestObject) (orgchartpublic.UpdateEmployeeInfoResponseObject, error) {
	//TODO implement me
	panic("implement me")
}
