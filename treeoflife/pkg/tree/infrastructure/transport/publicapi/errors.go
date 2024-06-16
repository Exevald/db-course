package publicapi

import (
	"context"
	"net/http"

	strictnethttp "github.com/oapi-codegen/runtime/strictmiddleware/nethttp"
	"github.com/pkg/errors"

	"tree/api/server/treepublic"
)

func NewErrorsMiddleware() treepublic.StrictMiddlewareFunc {
	return func(f strictnethttp.StrictHTTPHandlerFunc, operationID string) strictnethttp.StrictHTTPHandlerFunc {
		return func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (response interface{}, err error) {
			resp, err := f(ctx, w, r, request)
			switch errors.Cause(err) {
			case nil:
				return resp, nil
			}
			return nil, err
		}
	}
}
