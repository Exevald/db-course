package query

import "tree/pkg/tree/domain/model"

type NodeDataQueryService interface {
	GetNode(nodeID int64) (model.NodeData, error)
	GetParentNode(nodeID int64) (model.NodeData, error)
	GetNodePath(nodeID int64) ([]string, error)
}
