package publicapi

import (
	"context"
	"net/http"

	strictnethttp "github.com/oapi-codegen/runtime/strictmiddleware/nethttp"
	"github.com/pkg/errors"

	"orgchart/api/server/orgchartpublic"
	"orgchart/pkg/orgchart/app/model"
)

func NewErrorsMiddleware() orgchartpublic.StrictMiddlewareFunc {
	return func(f strictnethttp.StrictHTTPHandlerFunc, operationID string) strictnethttp.StrictHTTPHandlerFunc {
		return func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (response interface{}, err error) {
			resp, err := f(ctx, w, r, request)
			switch errors.Cause(err) {
			case model.ErrBranchNotFound:
				return orgchartpublic.NotFoundJSONResponse{
					Code:    orgchartpublic.BranchNotFound,
					Message: err.Error(),
				}, nil
			case model.ErrEmployeeNotFound:
				return orgchartpublic.NotFoundJSONResponse{
					Code:    orgchartpublic.EmployeeNotFound,
					Message: err.Error(),
				}, nil
			case model.ErrInvalidCity:
				return orgchartpublic.BadRequestJSONResponse{
					Code:    orgchartpublic.InvalidCity,
					Message: err.Error(),
				}, nil
			case model.ErrInvalidAddress:
				return orgchartpublic.BadRequestJSONResponse{
					Code:    orgchartpublic.InvalidAddress,
					Message: err.Error(),
				}, nil
			case model.ErrInvalidName:
				return orgchartpublic.BadRequestJSONResponse{
					Code:    orgchartpublic.InvalidName,
					Message: err.Error(),
				}, nil
			case model.ErrInvalidJobTitle:
				return orgchartpublic.BadRequestJSONResponse{
					Code:    orgchartpublic.InvalidJobTitle,
					Message: err.Error(),
				}, nil
			case model.ErrInvalidEmail:
				return orgchartpublic.BadRequestJSONResponse{
					Code:    orgchartpublic.InvalidEmail,
					Message: err.Error(),
				}, nil
			case model.ErrInvalidAge, model.ErrInvalidHireDate:
				return orgchartpublic.BadRequestJSONResponse{
					Code:    orgchartpublic.InvalidDate,
					Message: err.Error(),
				}, nil
			case model.ErrInvalidGender:
				return orgchartpublic.BadRequestJSONResponse{
					Code:    orgchartpublic.InvalidGender,
					Message: err.Error(),
				}, nil
			case nil:
				return resp, nil
			}
			return nil, err
		}
	}
}
