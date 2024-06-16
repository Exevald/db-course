package tests

type Branch struct {
	Address string
	City    string
}

type Employee struct {
	BirthDate  string
	BranchId   string
	Comment    string
	Email      string
	FirstName  string
	Gender     int
	HireDate   string
	JobTitle   string
	LastName   string
	MiddleName string
	Phone      string
	PhotoPath  *string
}

type BranchPreview struct {
	Address          string
	City             string
	CountOfEmployees int
}

type EmployeePreview struct {
	EmployeeId string
	FirstName  string
	JobTitle   string
	LastName   string
	MiddleName string
}

type CreateBranchRequestData struct {
	Address string
	City    string
}

type CreateBranchResponseData struct {
	BranchID string
}

type ListBranchesResponseData struct {
	Branches []BranchPreview
}

type GetBranchInfoRequestData struct {
	BranchID string
}

type GetBranchInfoResponseData struct {
	Address   string
	BranchId  string
	City      string
	Employees []EmployeePreview
}

type DeleteBranchRequestData struct {
	BranchID string
}

type UpdateBranchRequestData struct {
	BranchID string
	Branch   Branch
}

type CreateEmployeeRequestData struct {
	BirthDate  string
	BranchId   string
	Comment    string
	Email      string
	FirstName  string
	Gender     int
	HireDate   string
	JobTitle   string
	LastName   string
	MiddleName string
	Phone      string
	PhotoPath  *string
}

type CreateEmployeeResponseData struct {
	EmployeeId string
}

type GetEmployeeInfoRequestData struct {
	EmployeeId string
}

type GetEmployeeInfoResponseData struct {
	BirthDate  string
	BranchID   string
	Comment    string
	Email      string
	FirstName  string
	Gender     int
	HireDate   string
	JobTitle   string
	LastName   string
	MiddleName string
	Phone      string
	PhotoPath  *string
}

type DeleteEmployeeRequestData struct {
	EmployeeId string
}

type UpdateEmployeeInfoRequestData struct {
	EmployeeId string
	Employee   Employee
}

type OrgchartPublicAPI interface {
	CreateBranch(request CreateBranchRequestData) (CreateBranchResponseData, error)
	ListBranches() (ListBranchesResponseData, error)
	GetBranchInfo(request GetBranchInfoRequestData) (GetBranchInfoResponseData, error)
	DeleteBranch(request DeleteBranchRequestData) error
	UpdateBranchInfo(request UpdateBranchRequestData) error
	CreateEmployee(request CreateEmployeeRequestData) (CreateEmployeeResponseData, error)
	GetEmployeeInfo(request GetEmployeeInfoRequestData) (GetEmployeeInfoResponseData, error)
	DeleteEmployee(request DeleteEmployeeRequestData) error
	UpdateEmployeeInfo(request UpdateEmployeeInfoRequestData) error
}
