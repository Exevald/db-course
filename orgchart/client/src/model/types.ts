enum Gender {
    Male,
    Female
}

type Branch = {
    id?: string,
    city: string,
    address: string,
    employeesList: Array<Employee>
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

export type {Branch, Employee, Gender}