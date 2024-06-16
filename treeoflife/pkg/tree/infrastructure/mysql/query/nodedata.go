package query

import (
	"tree/pkg/tree/app/query"
	"tree/pkg/tree/domain/model"
)

func NewNodeDataQueryService() query.NodeDataQueryService {
	return &nodeDataQueryService{}
}

type nodeDataQueryService struct {
}

func (n nodeDataQueryService) GetNode(nodeID int64) (model.NodeData, error) {
	panic("implement me")
}

func (n nodeDataQueryService) GetParentNode(nodeID int64) (model.NodeData, error) {
	panic("implement me")
}

func (n nodeDataQueryService) GetNodePath(nodeID int64) ([]string, error) {
	panic("implement me")
}
