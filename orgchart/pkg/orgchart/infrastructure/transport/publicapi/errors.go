package publicapi

import (
	"context"
	"encoding/json"
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
				w.WriteHeader(404)
				err2 := json.NewEncoder(w).Encode(orgchartpublic.NotFoundJSONResponse{
					Code:    orgchartpublic.BranchNotFound,
					Message: err.Error(),
				})
				return nil, err2
			case model.ErrEmployeeNotFound:
				w.WriteHeader(404)
				err2 := json.NewEncoder(w).Encode(orgchartpublic.NotFoundJSONResponse{
					Code:    orgchartpublic.EmployeeNotFound,
					Message: err.Error(),
				})
				return nil, err2
			case model.ErrInvalidCity:
				w.WriteHeader(400)
				err2 := json.NewEncoder(w).Encode(orgchartpublic.BadRequestJSONResponse{
					Code:    orgchartpublic.InvalidCity,
					Message: err.Error(),
				})
				return nil, err2
			case model.ErrInvalidAddress:
				w.WriteHeader(400)
				err2 := json.NewEncoder(w).Encode(orgchartpublic.BadRequestJSONResponse{
					Code:    orgchartpublic.InvalidAddress,
					Message: err.Error(),
				})
				return nil, err2
			case model.ErrInvalidName:
				w.WriteHeader(400)
				err2 := json.NewEncoder(w).Encode(orgchartpublic.BadRequestJSONResponse{
					Code:    orgchartpublic.InvalidName,
					Message: err.Error(),
				})
				return nil, err2
			case model.ErrInvalidJobTitle:
				w.WriteHeader(400)
				err2 := json.NewEncoder(w).Encode(orgchartpublic.BadRequestJSONResponse{
					Code:    orgchartpublic.InvalidJobTitle,
					Message: err.Error(),
				})
				return nil, err2
			case model.ErrInvalidEmail:
				w.WriteHeader(400)
				err2 := json.NewEncoder(w).Encode(orgchartpublic.BadRequestJSONResponse{
					Code:    orgchartpublic.InvalidEmail,
					Message: err.Error(),
				})
				return nil, err2
			case model.ErrInvalidAge, model.ErrInvalidHireDate:
				w.WriteHeader(400)
				err2 := json.NewEncoder(w).Encode(orgchartpublic.BadRequestJSONResponse{
					Code:    orgchartpublic.InvalidDate,
					Message: err.Error(),
				})
				return nil, err2
			case model.ErrInvalidGender:
				w.WriteHeader(400)
				err2 := json.NewEncoder(w).Encode(orgchartpublic.BadRequestJSONResponse{
					Code:    orgchartpublic.InvalidGender,
					Message: err.Error(),
				})
				return nil, err2
			case nil:
				return resp, nil
			}
			return nil, err
		}
	}
}
