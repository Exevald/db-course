import {createBranch, deleteBranch, deleteEmployee, updateBranch} from "./requests";
import {branchInfoUrl, branchListUrl} from "./routes";
import ToastManager from "../view/components/ToastManager/ToastManager";
import {TOAST_ANIMATION_TIME} from "../model/utilities";
import {responseStatus} from "./responseStatus";

function createBranchAction() {
    const city = document.querySelector("#city") as HTMLInputElement
    const address = document.querySelector("#address") as HTMLInputElement
    let errCode: number
    if (city && address) {
        createBranch(city.value, address.value).then(
            (response) => {
                if (!response.ok) {
                    errCode = response.status
                    throw new Error('Error occurred!')
                }
                window.location.href = branchListUrl
                ToastManager.add('Успешно создано', TOAST_ANIMATION_TIME)
            }
        ).catch(() => {
            if (errCode === responseStatus.badRequest) {
                ToastManager.add("Введены неверные данные", TOAST_ANIMATION_TIME)
            }
        })
    }
}

function updateBranchAction(branchId: string) {
    const city = document.querySelector("#city") as HTMLInputElement
    const address = document.querySelector("#address") as HTMLInputElement
    let errCode: number
    if (city.value !== "" && address.value !== "") {
        updateBranch(branchId, city.value, address.value).then(
            (response) => {
                if (!response.ok) {
                    errCode = response.status
                    throw new Error('Error occurred!')
                }
                window.location.href = branchInfoUrl.replace("BRANCH_ID", branchId)
                ToastManager.add('Успешно создано', TOAST_ANIMATION_TIME)
            }
        ).catch(() => {
            if (errCode === responseStatus.badRequest) {
                ToastManager.add("Введены неверные данные", TOAST_ANIMATION_TIME)
            }
        })
    }
}

function deleteBranchAction(branchId: string) {
    let errCode: number
    deleteBranch(branchId).then(
        (response) => {
            if (!response.ok) {
                errCode = response.status
                throw new Error('Error occurred!')
            }
            window.location.href = branchListUrl
            ToastManager.add('Успешно удалено', TOAST_ANIMATION_TIME)
        }
    ).catch(() => {
        if (errCode === responseStatus.badRequest) {
            ToastManager.add("Введены неверные данные", TOAST_ANIMATION_TIME)
        }
    })
}

function deleteEmployeeAction(employeeId: string, branchId: string) {
    let errCode: number
    deleteEmployee(employeeId).then(
        (response) => {
            if (!response.ok) {
                errCode = response.status
                throw new Error('Error occurred!')
            }
            window.location.href = branchInfoUrl.replace("BRANCH_ID", branchId)
            ToastManager.add('Успешно удалено', TOAST_ANIMATION_TIME)
        }
    ).catch(() => {
        if (errCode === responseStatus.badRequest) {
            ToastManager.add("Введены неверные данные", TOAST_ANIMATION_TIME)
        }
    })
}

export {createBranchAction, updateBranchAction, deleteBranchAction, deleteEmployeeAction}