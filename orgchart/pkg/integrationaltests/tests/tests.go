package tests

import (
	"fmt"
	"orgchart/pkg/orgchart/app/service"
	"reflect"
)

var (
	defaultBranchCity    = "Yoshkar-Ola"
	defaultBranchAddress = "Volkova 108"

	changedBranchCity    = "Kazan"
	changedBranchAddress = "Voznesenskaya 110"
)

var (
	defaultEmployeeFirstName  = "Ivan"
	defaultEmployeeLastName   = "Petrov"
	defaultEmployeeMiddleName = "Sergeevich"
	defaultEmployeeJobTitle   = "Programmer"
	defaultEmployeePhone      = "+79024563278"
	defaultEmployeeEmail      = "ivan.petrov@email.com"
	defaultEmployeeGender     = 0
	defaultEmployeeBirthDate  = "2004-10-28"
	defaultEmployeeHireDate   = "2006-01-02"
	defaultEmployeeComment    = "good programmer"

	changedEmployeeFirstName = "Oleg"
	changedEmployeeLastName  = "Sergeev"
)

func RunTests(api OrgchartPublicAPI,
	branchService service.BranchService,
	employeeService service.EmployeeService,
) {
	//RunFunctionalTests(api)
	RunComponentTests(branchService, employeeService)
}

func RunFunctionalTests(api OrgchartPublicAPI) {
	testOrgchartPublic(api)
}

func RunComponentTests(
	branchService service.BranchService,
	employeeService service.EmployeeService,
) {
	testComponent(branchService, employeeService)
}

func assert(testingParameter string, reference, value interface{}) {
	if !reflect.DeepEqual(reference, value) {
		panic(fmt.Sprintf("Assertion failed on %s: expected: %v, got %v", testingParameter, reference, value))
	}
}

func assertErr(err error) {
	if err != nil {
		panic(err)
	}
}
