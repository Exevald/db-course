package query

import "trees/pkg/tree/domain/model"

type NodeDataQueryService interface {
	GetNode(nodeID int64) (model.NodeData, error)
	GetParentNode(nodeID int64) (model.NodeData, error)
	GetNodePath(nodeID int64) ([]string, error)
}

func NewNodeDataQueryService() NodeDataQueryService {
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
