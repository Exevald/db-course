package model

import "github.com/pkg/errors"

type NodeData interface {
	ID() int64
	Name() string
	IsExtinct() bool
	Confidence() int
}

func NewNodeData(id int64, name string, isExtinct bool, confidence int) (NodeData, error) {
	if name == "" {
		return nil, errors.WithStack(ErrInvalidEntityName)
	}
	if confidence < 0 {
		return nil, errors.WithStack(ErrInvalidEntityConfidence)
	}

	return &nodeData{
		id:         id,
		name:       name,
		isExtinct:  isExtinct,
		confidence: confidence,
	}, nil
}

func LoadNodeData(id int64, name string, isExtinct bool, confidence int) NodeData {
	return &nodeData{
		id:         id,
		name:       name,
		isExtinct:  isExtinct,
		confidence: confidence,
	}
}

type nodeData struct {
	id         int64
	name       string
	isExtinct  bool
	confidence int
}

func (n *nodeData) ID() int64 {
	return n.id
}

func (n *nodeData) Name() string {
	return n.name
}

func (n *nodeData) IsExtinct() bool {
	return n.isExtinct
}

func (n *nodeData) Confidence() int {
	return n.confidence
}

type NodeDataRepository interface {
	Get(id int64) (NodeData, error)
	Store(node NodeData) error
	Delete(id int64) error
}
