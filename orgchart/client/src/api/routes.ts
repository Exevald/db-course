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
const getBranchInfoAPIUrl = serviceUrl + "/api/v1/orgchart/branch/BRANCH_ID"
const getEmployeeInfoAPIUrl = serviceUrl + "/api/v1/orgchart/employee/EMPLOYEE_ID"

/**
 * API routes
 */
const createBranchUrl = serviceUrl + "/api/v1/orgchart/branch/create"
const createEmployeeUrl = serviceUrl + "/api/v1/orgchart/employee/create"
const updateBranchInfoUrl = serviceUrl + "/api/v1/orgchart/branch/BRANCH_ID/update"
const updateEmployeeInfoUrl = serviceUrl + "/api/v1/orgchart/employee/EMPLOYEE_ID/update"
const deleteBranchUrl = serviceUrl + "/api/v1/orgchart/branch/BRANCH_ID/delete"
const deleteEmployeeUrl = serviceUrl + "/api/v1/orgchart/employee/EMPLOYEE_ID/delete"

export {
    branchListUrl, branchInfoUrl, employeeInfoUrl, getBranchListAPIUrl, getBranchInfoAPIUrl, getEmployeeInfoAPIUrl,
    createBranchUrl, createEmployeeUrl, updateBranchInfoUrl, updateEmployeeInfoUrl, deleteBranchUrl, deleteEmployeeUrl
}