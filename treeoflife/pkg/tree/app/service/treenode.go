package service

import (
	"tree/pkg/tree/domain/model"
)

type TreeNodeService interface {
	SaveTree(root model.TreeNode) error
	AddNode(node model.NodeData, parentID int64) error
	MoveSubTree(treeNodeID int64, newParentID int64) error
	DeleteSubTree(treeNodeID int64) error
}

type treeNodeService struct {
}
