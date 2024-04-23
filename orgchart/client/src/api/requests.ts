import {fetchDeleteRequest, fetchPostRequest, fetchUpdateRequest} from "./fetchRequest";
import {createBranchAPIUrl, deleteBranchAPIUrl, deleteEmployeeAPIUrl, updateBranchInfoAPIUrl} from "./routes";

function createBranch(city: string, address: string) {
    return fetchPostRequest(
        createBranchAPIUrl,
        {
            city,
            address
        }
    )
}

function updateBranch(branchId: string,city: string, address: string) {
    return fetchUpdateRequest(
        updateBranchInfoAPIUrl.replace("BRANCH_ID", branchId), {
            city,
            address
        }
    )
}

function deleteBranch(branchId: string) {
    return fetchDeleteRequest(
        deleteBranchAPIUrl.replace("BRANCH_ID", branchId)
    )
}

function deleteEmployee(employeeId: string) {
    return fetchDeleteRequest(
        deleteEmployeeAPIUrl.replace("EMPLOYEE_ID", employeeId)
    )
}

export {createBranch, updateBranch, deleteBranch, deleteEmployee}