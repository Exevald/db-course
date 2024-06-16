package tests

import (
	"orgchart/pkg/orgchart/app/model"
	"orgchart/pkg/orgchart/app/service"
	"time"
)

func testComponent(branchService service.BranchService, employeeService service.EmployeeService) {
	testComponentCreateAndUpdateBranchScenario(branchService)
	testComponentCreateAndUpdateEmployeeScenario(branchService, employeeService)
	testComponentChangeBranchAndEmployeeScenario(branchService, employeeService)
}

func testComponentCreateAndUpdateBranchScenario(branchService service.BranchService) {
	branchId, err := branchService.CreateBranch("Yoshkar-Ola", "Volkova 108")
	assertErr(err)
	branchInfo, err := branchService.GetBranchInfo(branchId)
	assertErr(err)
	assertBranchFields(branchInfo, "Yoshkar-Ola", "Volkova 108")

	updatedBranchInfo, err := model.NewBranch(
		branchId,
		"Kazan",
		"Voznesenskaya 110",
	)
	assertErr(err)
	err = branchService.UpdateBranch(updatedBranchInfo)
	assertErr(err)
	branchInfo, err = branchService.GetBranchInfo(branchId)
	assertErr(err)
	assertBranchFields(branchInfo, "Kazan", "Voznesenskaya 110")
}

func testComponentCreateAndUpdateEmployeeScenario(branchService service.BranchService, employeeService service.EmployeeService) {
	branchId, err := branchService.CreateBranch("Yoshkar-Ola", "Voznesenskaya 110")
	assertErr(err)
	employeeId, err := employeeService.CreateEmployee(
		branchId,
		"Ivan",
		"Petrov",
		"Sergeevich",
		"Programmer",
		"+79023524374",
		"ivan.petrov@email.com",
		model.Male,
		time.Now().Add(time.Hour*10),
		time.Now().Add(time.Hour*346),
		nil,
		nil,
	)
	assertErr(err)
	employeeInfo, err := employeeService.GetEmployeeInfo(employeeId)
	assertErr(err)
	assert("employee first name", "Ivan", employeeInfo.FirstName())
	assert("employee last name", "Petrov", employeeInfo.LastName())
	assert("employee middle name", "Sergeevich", employeeInfo.MiddleName())
	assert("employee job title", "Programmer", employeeInfo.JobTitle())
	assert("employee phone number", "+79023524374", employeeInfo.Phone())
	assert("employee email", "ivan.petrov@email.com", employeeInfo.Email())
	assert("employee gender", model.Male, employeeInfo.Gender())

	updatedEmployeeInfo, err := model.NewEmployee(
		employeeId,
		branchId,
		"Oleg",
		"Sergeev",
		"Sergeevich",
		"Programmer",
		"+79023524374",
		"ivan.petrov@email.com",
		model.Male,
		time.Now().Add(time.Hour*10),
		time.Now().Add(time.Hour*346),
		nil,
		nil,
	)
	assertErr(err)
	err = employeeService.UpdateEmployee(updatedEmployeeInfo)
	assertErr(err)
	employeeInfo, err = employeeService.GetEmployeeInfo(employeeId)
	assertErr(err)
	assert("employee first name", "Oleg", employeeInfo.FirstName())
	assert("employee last name", "Sergeev", employeeInfo.LastName())

	err = employeeService.DeleteEmployee(employeeId)
	assertErr(err)
	branchEmployees, err := employeeService.GetBranchEmployees(branchId)
	assertErr(err)
	assert("branch employees", 0, len(branchEmployees))
}

func testComponentChangeBranchAndEmployeeScenario(branchService service.BranchService, employeeService service.EmployeeService) {
	branchId1, err := branchService.CreateBranch("Yoshkar-Ola", "Voznesenskaya 110")
	assertErr(err)
	branchId2, err := branchService.CreateBranch("Kazan", "Rubinova 44")
	assertErr(err)
	employeeId1, err := employeeService.CreateEmployee(
		branchId1,
		"Ivan",
		"Petrov",
		"Sergeevich",
		"Programmer",
		"+79023524374",
		"ivan.petrov@email.com",
		model.Male,
		time.Now().Add(time.Hour*10),
		time.Now().Add(time.Hour*346),
		nil,
		nil,
	)
	assertErr(err)
	employeeId2, err := employeeService.CreateEmployee(
		branchId2,
		"Konstantin",
		"Olegov",
		"Petrovich",
		"Programmer",
		"+79023524374",
		"ivan.petrov@email.com",
		model.Male,
		time.Now().Add(time.Hour*10),
		time.Now().Add(time.Hour*346),
		nil,
		nil,
	)
	assertErr(err)
	employeeInfo1, err := employeeService.GetEmployeeInfo(employeeId1)
	assertErr(err)
	assert("employee first name", "Ivan", employeeInfo1.FirstName())
	assert("employee last name", "Petrov", employeeInfo1.LastName())

	employeeInfo2, err := employeeService.GetEmployeeInfo(employeeId2)
	assertErr(err)
	assert("employee first name", "Konstantin", employeeInfo2.FirstName())
	assert("employee last name", "Olegov", employeeInfo2.LastName())

	err = employeeService.DeleteEmployee(employeeId1)
	assertErr(err)

	branchEmployees1, err := employeeService.GetBranchEmployees(branchId1)
	assertErr(err)
	assert("branch employees count", 0, len(branchEmployees1))

	branchEmployees2, err := employeeService.GetBranchEmployees(branchId2)
	assertErr(err)
	assert("branch employees count", 1, len(branchEmployees2))
}

func assertBranchFields(referenceBranch model.Branch, expectedCity, expectedAddress string) {
	assert("branch city", expectedCity, referenceBranch.City())
	assert("branch address", expectedAddress, referenceBranch.Address())
}
