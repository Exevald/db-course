enum Gender {
    Male,
    Female
}


type ViewBranch = {
    city: string,
    address: string,
    countOfEmployees: number
}

type Employee = {
    id?: string,
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

export type {Employee, Gender,ViewBranch}