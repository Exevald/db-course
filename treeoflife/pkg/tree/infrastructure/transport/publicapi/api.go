package publicapi

import (
	"context"

	"tree/api/server/treepublic"
)

type PublicAPI treepublic.StrictServerInterface

func NewPublicAPI() PublicAPI {
	return &publicAPI{}
}

type publicAPI struct{}

func (p publicAPI) GetTree(ctx context.Context, request treepublic.GetTreeRequestObject) (treepublic.GetTreeResponseObject, error) {
	panic("implement me")
}

func (p publicAPI) GetChildren(ctx context.Context, request treepublic.GetChildrenRequestObject) (treepublic.GetChildrenResponseObject, error) {
	panic("implement me")
}

func (p publicAPI) AddNode(ctx context.Context, request treepublic.AddNodeRequestObject) (treepublic.AddNodeResponseObject, error) {
	panic("implement me")
}

func (p publicAPI) GetNodePath(ctx context.Context, request treepublic.GetNodePathRequestObject) (treepublic.GetNodePathResponseObject, error) {
	panic("implement me")
}

func (p publicAPI) GetNode(ctx context.Context, request treepublic.GetNodeRequestObject) (treepublic.GetNodeResponseObject, error) {
	panic("implement me")
}

func (p publicAPI) GetParentNode(ctx context.Context, request treepublic.GetParentNodeRequestObject) (treepublic.GetParentNodeResponseObject, error) {
	panic("implement me")
}

func (p publicAPI) SaveTree(ctx context.Context, request treepublic.SaveTreeRequestObject) (treepublic.SaveTreeResponseObject, error) {
	panic("implement me")
}

func (p publicAPI) DeleteSubTree(ctx context.Context, request treepublic.DeleteSubTreeRequestObject) (treepublic.DeleteSubTreeResponseObject, error) {
	panic("implement me")
}

func (p publicAPI) MoveSubTree(ctx context.Context, request treepublic.MoveSubTreeRequestObject) (treepublic.MoveSubTreeResponseObject, error) {
	panic("implement me")
}

func (p publicAPI) GetSubTree(ctx context.Context, request treepublic.GetSubTreeRequestObject) (treepublic.GetSubTreeResponseObject, error) {
	panic("implement me")
}
