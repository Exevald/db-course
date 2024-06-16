package query

import (
	"tree/pkg/tree/app/query"
	"tree/pkg/tree/domain/model"
)

func NewTreeNodeQueryService() query.TreeNodeQueryService {
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
