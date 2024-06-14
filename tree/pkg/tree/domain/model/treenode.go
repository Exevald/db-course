package model

import (
	stderrors "errors"

	"github.com/pkg/errors"
)

var (
	ErrTreeNodeNotFound        = stderrors.New("tree node not found")
	ErrInvalidNodeId           = stderrors.New("invalid node id")
	ErrInvalidTreeNode         = stderrors.New("invalid tree node")
	ErrInvalidEntityName       = stderrors.New("invalid entity node name")
	ErrInvalidEntityConfidence = stderrors.New("invalid entity confidence")
	ErrInvalidParentId         = stderrors.New("invalid parent id")
)

type TreeNode interface {
	ID() int64
	Name() string
	IsExtinct() bool
	Confidence() int
	ParentNodeID() int64
	Children() []*TreeNode
}

func NewTreeNode(
	id int64,
	name string,
	isExtinct bool,
	confidence int,
	parentNodeID int64,
) (TreeNode, error) {
	if name == "" {
		return nil, errors.WithStack(ErrInvalidEntityName)
	}
	if confidence < 0 {
		return nil, errors.WithStack(ErrInvalidEntityConfidence)
	}
	if parentNodeID < 0 || parentNodeID == id {
		return nil, errors.WithStack(ErrInvalidParentId)
	}

	return &treeNode{
		id:           id,
		name:         name,
		isExtinct:    isExtinct,
		confidence:   confidence,
		parentNodeID: parentNodeID,
		children:     nil,
	}, nil
}

func LoadTreeNode(
	id int64,
	name string,
	isExtinct bool,
	confidence int,
	parentNodeID int64,
	children []*TreeNode,
) TreeNode {
	return &treeNode{
		id:           id,
		name:         name,
		isExtinct:    isExtinct,
		confidence:   confidence,
		parentNodeID: parentNodeID,
		children:     children,
	}
}

type treeNode struct {
	id           int64
	name         string
	isExtinct    bool
	confidence   int
	parentNodeID int64
	children     []*TreeNode
}

func (t *treeNode) ID() int64 {
	return t.id
}

func (t *treeNode) Name() string {
	return t.name
}

func (t *treeNode) IsExtinct() bool {
	return t.isExtinct
}

func (t *treeNode) Confidence() int {
	return t.confidence
}

func (t *treeNode) ParentNodeID() int64 {
	return t.parentNodeID
}

func (t *treeNode) Children() []*TreeNode {
	return t.children
}
