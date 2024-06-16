package query

import "tree/pkg/tree/domain/model"

type TreeNodeQueryService interface {
	GetTree() (model.TreeNode, error)
	GetSubTree(nodeID int64) (model.TreeNode, error)
	GetChildren(nodeID int64) ([]model.TreeNode, error)
}
