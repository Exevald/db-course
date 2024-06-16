package tests

func testOrgchartPublic(api OrgchartPublicAPI) {
	testCreateAndUpdateBranchScenario(api)
	testCreateAndUpdateEmployeeScenario(api)
	testChangeBranchAndEmployeeScenario(api)
}

func testCreateAndUpdateBranchScenario(api OrgchartPublicAPI) {
	createBranchResponseData, err := api.CreateBranch(CreateBranchRequestData{
		City:    defaultBranchCity,
		Address: defaultBranchAddress,
	})
	assertErr(err)
	branchInfo, err := api.GetBranchInfo(GetBranchInfoRequestData{
		BranchID: createBranchResponseData.BranchID,
	})
	assertErr(err)
	// Вынести в вспомогательную фукнцию
	assert("branch city", defaultBranchCity, branchInfo.City)
	assert("branch address", defaultBranchAddress, branchInfo.Address)
	assert("count of branch employees", 0, len(branchInfo.Employees))

	err = api.UpdateBranchInfo(UpdateBranchRequestData{
		BranchID: createBranchResponseData.BranchID,
		Branch: Branch{
			Address: changedBranchAddress,
			City:    changedBranchCity,
		},
	})
	assertErr(err)
	branchInfo, err = api.GetBranchInfo(GetBranchInfoRequestData{
		BranchID: createBranchResponseData.BranchID,
	})
	assertErr(err)
	assert("branch city", changedBranchCity, branchInfo.City)
	assert("branch address", changedBranchAddress, branchInfo.Address)

	err = api.DeleteBranch(DeleteBranchRequestData{BranchID: createBranchResponseData.BranchID})
	assertErr(err)
	listBranchesResponse, err := api.ListBranches()
	assertErr(err)
	assert("count of branches", 0, len(listBranchesResponse.Branches))
}

func testCreateAndUpdateEmployeeScenario(api OrgchartPublicAPI) {
	createBranchResponseData, err := api.CreateBranch(CreateBranchRequestData{
		City:    defaultBranchCity,
		Address: defaultBranchAddress,
	})
	assertErr(err)
	createEmployeeResponseData, err := api.CreateEmployee(CreateEmployeeRequestData{
		BirthDate:  defaultEmployeeBirthDate,
		BranchId:   createBranchResponseData.BranchID,
		Comment:    defaultEmployeeComment,
		Email:      defaultEmployeeEmail,
		FirstName:  defaultEmployeeFirstName,
		Gender:     defaultEmployeeGender,
		HireDate:   defaultEmployeeHireDate,
		JobTitle:   defaultEmployeeJobTitle,
		LastName:   defaultEmployeeLastName,
		MiddleName: defaultEmployeeMiddleName,
		Phone:      defaultEmployeePhone,
		PhotoPath:  nil,
	})
	assertErr(err)
	branchInfo, err := api.GetBranchInfo(GetBranchInfoRequestData{
		BranchID: createBranchResponseData.BranchID,
	})
	assertErr(err)
	assert("count of employees", 1, len(branchInfo.Employees))
	employeeInfo, err := api.GetEmployeeInfo(GetEmployeeInfoRequestData{
		EmployeeId: createEmployeeResponseData.EmployeeId,
	})
	assertErr(err)
	assert("employee first name", defaultEmployeeFirstName, employeeInfo.FirstName)
	assert("employee last name", defaultEmployeeLastName, employeeInfo.LastName)
	assert("employee middle name", defaultEmployeeMiddleName, employeeInfo.MiddleName)
	assert("employee job title", defaultEmployeeJobTitle, employeeInfo.JobTitle)
	assert("employee phone number", defaultEmployeePhone, employeeInfo.Phone)
	assert("employee email", defaultEmployeeEmail, employeeInfo.Email)
	assert("employee gender", defaultEmployeeGender, employeeInfo.Gender)
	assert("employee comment", defaultEmployeeComment, employeeInfo.Comment)

	err = api.UpdateEmployeeInfo(UpdateEmployeeInfoRequestData{
		EmployeeId: createEmployeeResponseData.EmployeeId,
		Employee: Employee{
			BirthDate:  defaultEmployeeBirthDate,
			BranchId:   createBranchResponseData.BranchID,
			Comment:    defaultEmployeeComment,
			Email:      defaultEmployeeEmail,
			FirstName:  changedEmployeeFirstName,
			Gender:     defaultEmployeeGender,
			HireDate:   defaultEmployeeHireDate,
			JobTitle:   defaultEmployeeJobTitle,
			LastName:   changedEmployeeLastName,
			MiddleName: defaultEmployeeMiddleName,
			Phone:      defaultEmployeePhone,
			PhotoPath:  nil,
		},
	})
	assertErr(err)
	employeeInfo, err = api.GetEmployeeInfo(GetEmployeeInfoRequestData{
		EmployeeId: createEmployeeResponseData.EmployeeId,
	})
	assertErr(err)
	assert("employee first name", changedEmployeeFirstName, employeeInfo.FirstName)
	assert("employee last name", changedEmployeeLastName, employeeInfo.LastName)

	err = api.DeleteEmployee(DeleteEmployeeRequestData{EmployeeId: createEmployeeResponseData.EmployeeId})
	assertErr(err)
	branchInfo, err = api.GetBranchInfo(GetBranchInfoRequestData{
		BranchID: createBranchResponseData.BranchID,
	})
	assertErr(err)
	assert("count of branch employees", 0, len(branchInfo.Employees))
}

func testChangeBranchAndEmployeeScenario(api OrgchartPublicAPI) {
	createBranchResponseData1, err := api.CreateBranch(CreateBranchRequestData{
		City:    defaultBranchCity,
		Address: defaultBranchAddress,
	})
	assertErr(err)
	createBranchResponseData2, err := api.CreateBranch(CreateBranchRequestData{
		City:    changedBranchCity,
		Address: changedBranchAddress,
	})
	assertErr(err)
	createEmployeeResponseData1, err := api.CreateEmployee(CreateEmployeeRequestData{
		BirthDate:  defaultEmployeeBirthDate,
		BranchId:   createBranchResponseData1.BranchID,
		Comment:    defaultEmployeeComment,
		Email:      defaultEmployeeEmail,
		FirstName:  defaultEmployeeFirstName,
		Gender:     defaultEmployeeGender,
		HireDate:   defaultEmployeeHireDate,
		JobTitle:   defaultEmployeeJobTitle,
		LastName:   defaultEmployeeLastName,
		MiddleName: defaultEmployeeMiddleName,
		Phone:      defaultEmployeePhone,
		PhotoPath:  nil,
	})
	assertErr(err)
	createEmployeeResponseData2, err := api.CreateEmployee(CreateEmployeeRequestData{
		BirthDate:  defaultEmployeeBirthDate,
		BranchId:   createBranchResponseData2.BranchID,
		Comment:    defaultEmployeeComment,
		Email:      defaultEmployeeEmail,
		FirstName:  changedEmployeeFirstName,
		Gender:     defaultEmployeeGender,
		HireDate:   defaultEmployeeHireDate,
		JobTitle:   defaultEmployeeJobTitle,
		LastName:   changedEmployeeLastName,
		MiddleName: defaultEmployeeMiddleName,
		Phone:      defaultEmployeePhone,
		PhotoPath:  nil,
	})
	assertErr(err)
	branchesListResponseData, err := api.ListBranches()
	assertErr(err)
	assert("count of branches", 2, len(branchesListResponseData.Branches))

	branchInfo1, err := api.GetBranchInfo(GetBranchInfoRequestData{
		BranchID: createBranchResponseData1.BranchID,
	})
	assertErr(err)
	assert("branch1 city", defaultBranchCity, branchInfo1.City)
	assert("branch1 address", defaultBranchAddress, branchInfo1.Address)
	branchInfo2, err := api.GetBranchInfo(GetBranchInfoRequestData{
		BranchID: createBranchResponseData2.BranchID,
	})
	assertErr(err)
	assert("branch2 city", changedBranchCity, branchInfo2.City)
	assert("branch2 address", changedBranchAddress, branchInfo2.Address)
	employeeInfo1, err := api.GetEmployeeInfo(GetEmployeeInfoRequestData{
		EmployeeId: createEmployeeResponseData1.EmployeeId,
	})
	assertErr(err)
	assert("employee1 first name", defaultEmployeeFirstName, employeeInfo1.FirstName)
	assert("employee1 last name", defaultEmployeeLastName, employeeInfo1.LastName)
	employeeInfo2, err := api.GetEmployeeInfo(GetEmployeeInfoRequestData{
		EmployeeId: createEmployeeResponseData2.EmployeeId,
	})
	assertErr(err)
	assert("employee2 first name", changedEmployeeFirstName, employeeInfo2.FirstName)
	assert("employee2 last name", changedEmployeeLastName, employeeInfo2.LastName)

	assert("employee1 branch id", createBranchResponseData1.BranchID, employeeInfo1.BranchID)
	assert("employee2 branch id", createBranchResponseData2.BranchID, employeeInfo2.BranchID)

	err = api.DeleteEmployee(DeleteEmployeeRequestData{createEmployeeResponseData1.EmployeeId})
	assertErr(err)
	err = api.DeleteEmployee(DeleteEmployeeRequestData{createEmployeeResponseData2.EmployeeId})
	assertErr(err)
	branchInfo1, err = api.GetBranchInfo(GetBranchInfoRequestData{
		BranchID: createBranchResponseData1.BranchID,
	})
	assertErr(err)
	branchInfo2, err = api.GetBranchInfo(GetBranchInfoRequestData{
		BranchID: createBranchResponseData2.BranchID,
	})
	assertErr(err)
	assert("branch1 count of employees", 0, len(branchInfo1.Employees))
	assert("branch2 count of employees", 0, len(branchInfo2.Employees))

	err = api.DeleteBranch(DeleteBranchRequestData{createBranchResponseData1.BranchID})
	assertErr(err)
	err = api.DeleteBranch(DeleteBranchRequestData{createBranchResponseData2.BranchID})
	assertErr(err)
	branchesListResponseData, err = api.ListBranches()
	assertErr(err)
	assert("count of branches", 0, len(branchesListResponseData.Branches))
}
