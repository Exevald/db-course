package infrastructure

import "tree/pkg/tree/common/mysql"

func NewDependencyContainer(connector mysql.Connector) (*DependencyContainer, error) {
	return &DependencyContainer{}, nil
}

type DependencyContainer struct {
}
