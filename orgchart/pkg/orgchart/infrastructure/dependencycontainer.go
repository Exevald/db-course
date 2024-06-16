package infrastructure

import (
	"orgchart/pkg/common/mysql"
	"orgchart/pkg/orgchart/app/service"
	"orgchart/pkg/orgchart/infrastructure/mysql/repository/branch"
	"orgchart/pkg/orgchart/infrastructure/mysql/repository/employee"
)

func NewDependencyContainer(connector mysql.Connector) (*DependencyContainer, error) {
	branchRepository := branch.NewBranchRepository(connector.Client())
	employeeRepository := employee.NewEmployeeRepository(connector.Client())
	branchService := service.NewBranchService(branchRepository)
	employeeService := service.NewEmployeeService(employeeRepository)

	return &DependencyContainer{
		branchService:   branchService,
		employeeService: employeeService,
	}, nil
}

type DependencyContainer struct {
	branchService   service.BranchService
	employeeService service.EmployeeService
}

func (container *DependencyContainer) BranchService() service.BranchService {
	return container.branchService
}

func (container *DependencyContainer) EmployeeService() service.EmployeeService {
	return container.employeeService
}
