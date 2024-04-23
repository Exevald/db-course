enum Gender {
    Male,
    Female
}

type Branch = {
    branchId: string
    city: string,
    address: string,
    employees: ViewEmployee[]
}

type ViewBranch = {
    branchId: string
    city: string,
    address: string,
    countOfEmployees: number
}

type ViewEmployee = {
    employeeId: string,
    branchId: string,
    firstName: string,
    lastName: string,
    middleName: string,
    jobTitle: string,
    phone: string,
    email: string,
    gender: Gender,
    birthDate: Date,
    hireDate: Date,
    comment?: string
    avatarPath?: string
}

export type {ViewEmployee, ViewBranch, Branch, Gender}