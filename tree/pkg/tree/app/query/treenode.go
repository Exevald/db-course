package query

import "trees/pkg/tree/domain/model"

type TreeNodeQueryService interface {
	GetTree() (model.TreeNode, error)
	GetSubTree(nodeID int64) (model.TreeNode, error)
	GetChildren(nodeID int64) ([]model.TreeNode, error)
}

func NewTreeNodeQueryService() TreeNodeQueryService {
	return &treeNodeQueryService{}
}

type treeNodeQueryService struct {
}

func (t treeNodeQueryService) GetTree() (model.TreeNode, error) {
	panic("implement me")
}

func (t treeNodeQueryService) GetSubTree(nodeID int64) (model.TreeNode, error) {
	panic("implement me")
}

func (t treeNodeQueryService) GetChildren(nodeID int64) ([]model.TreeNode, error) {
	panic("implement me")
}
