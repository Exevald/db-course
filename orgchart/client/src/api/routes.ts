const serviceUrl = "http://localhost:8882"

/**
 *  Pages routes
 */
const branchListUrl = "/"
const branchInfoUrl = "/branch/BRANCH_ID"
const employeeInfoUrl = "/employee/EMPLOYEE_ID"

/**
 * Page data API routes
 */
const getBranchListAPIUrl = serviceUrl + "/api/v1/orgchart/branch/list"
const getBranchInfoAPIUrl = "/api/v1/orgchart/branch/BRANCH_ID"
const getEmployeeInfoAPIUrl = "/api/v1/orgchart/employee/EMPLOYEE_ID"

/**
 * API routes
 */
const createBranchUrl = "/api/v1/orgchart/branch/create"
const createEmployeeUrl = "/api/v1/orgchart/employee/create"
const updateBranchInfoUrl = "/api/v1/orgchart/branch/BRANCH_ID/update"
const updateEmployeeInfoUrl = "/api/v1/orgchart/employee/EMPLOYEE_ID/update"
const deleteBranchUrl = "/api/v1/orgchart/branch/BRANCH_ID/delete"
const deleteEmployeeUrl = "/api/v1/orgchart/employee/EMPLOYEE_ID/delete"

export {
    branchListUrl, branchInfoUrl, employeeInfoUrl, getBranchListAPIUrl, getBranchInfoAPIUrl, getEmployeeInfoAPIUrl,
    createBranchUrl, createEmployeeUrl, updateBranchInfoUrl, updateEmployeeInfoUrl, deleteBranchUrl, deleteEmployeeUrl
}