const serviceUrl = "http://localhost:8882"

/**
 *  Pages routes
 */
const branchListUrl = "/"
const branchInfoUrl = "/branch?branchId=BRANCH_ID"
const employeeInfoUrl = "/employee?employeeId=EMPLOYEE_ID"
const createBranchUrl = "/branch/create"
const crateEmployeeUrl = "/employee/create?branchId=BRANCH_ID"

/**
 * Page data API routes
 */
const getBranchListAPIUrl = serviceUrl + "/api/v1/orgchart/branch/list"
const getBranchInfoAPIUrl = serviceUrl + "/api/v1/orgchart/branch/BRANCH_ID"
const getEmployeeInfoAPIUrl = serviceUrl + "/api/v1/orgchart/employee/EMPLOYEE_ID"

/**
 * API routes
 */
const createBranchAPIUrl = serviceUrl + "/api/v1/orgchart/branch/create"
const createEmployeeAPIUrl = serviceUrl + "/api/v1/orgchart/employee/create"
const updateBranchInfoAPIUrl = serviceUrl + "/api/v1/orgchart/branch/BRANCH_ID/update"
const updateEmployeeInfoAPIUrl = serviceUrl + "/api/v1/orgchart/employee/EMPLOYEE_ID/update"
const deleteBranchAPIUrl = serviceUrl + "/api/v1/orgchart/branch/BRANCH_ID/delete"
const deleteEmployeeAPIUrl = serviceUrl + "/api/v1/orgchart/employee/EMPLOYEE_ID/delete"

export {
    branchListUrl, branchInfoUrl, employeeInfoUrl, getBranchListAPIUrl, getBranchInfoAPIUrl, getEmployeeInfoAPIUrl,
    createBranchAPIUrl, createEmployeeAPIUrl, updateBranchInfoAPIUrl, updateEmployeeInfoAPIUrl, deleteBranchAPIUrl,
    deleteEmployeeAPIUrl, createBranchUrl, crateEmployeeUrl
}