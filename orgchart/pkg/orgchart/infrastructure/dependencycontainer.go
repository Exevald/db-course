package infrastructure

type DependencyContainer struct{}

func NewDependencyContainer() (*DependencyContainer, error) {
	return &DependencyContainer{}, nil
}
