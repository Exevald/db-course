package service

import "trees/pkg/tree/domain/model"

type TreeService interface {
	SaveTree(root model.TreeNode) error
	AddNode(node model.NodeData, parentID int64) error
	MoveSubTree(treeNodeID int64, newParentID int64) error
	DeleteSubTree(treeNodeID int64) error
}

func NewTreeService(
	treeNodeRepository model.TreeNodeRepository,
	nodeDataRepository model.NodeDataRepository,
) TreeService {
	return &treeService{
		treeNodeRepository: treeNodeRepository,
		nodeDataRepository: nodeDataRepository,
	}
}

type treeService struct {
	treeNodeRepository model.TreeNodeRepository
	nodeDataRepository model.NodeDataRepository
}

func (t treeService) SaveTree(root model.TreeNode) error {
	panic("implement me")
}

func (t treeService) AddNode(node model.NodeData, parentID int64) error {
	panic("implement me")
}

func (t treeService) MoveSubTree(treeNodeID int64, newParentID int64) error {
	panic("implement me")
}

func (t treeService) DeleteSubTree(treeNodeID int64) error {
	panic("implement me")
}
