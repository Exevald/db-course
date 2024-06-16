import styles from "./CreateEmployee.module.css"
import ReactDOM from "react-dom/client";
import React, {useEffect, useState} from "react";
import {TopPanel} from "../../components/TopPanel/TopPanel";
import {Breadcrumbs, BreadcrumbsItemProps} from "../../components/Breadcrumbs/Breadcrumbs";
import {Branch} from "../../../model/types";
import {fetchGetRequest} from "../../../api/fetchRequest";
import {branchInfoUrl, getBranchInfoAPIUrl} from "../../../api/routes";
import {TextArea} from "../../components/TextArea/TextArea";

interface CreateEmployeeProps {
    branchId: string
}

const CreateEmployee = (props: CreateEmployeeProps) => {
    const [data, setData] = useState<Branch>()
    useEffect(() => {
        fetchGetRequest(getBranchInfoAPIUrl.replace("BRANCH_ID", props.branchId)).then(data => {
            setData(data)
        }).catch(error => console.log(error))
    }, [])
    let breadcrumbsItems: BreadcrumbsItemProps[] = []
    if (data) {
        breadcrumbsItems = [
            {
                href: "/",
                text: "Филиалы",
                index: 0
            },
            {
                href: branchInfoUrl.replace("BRANCH_ID", props.branchId),
                text: data?.city + ", " + data?.address,
                index: 1
            },
            {
                href: "",
                text: "Создать сотрудника",
                index: 2
            },
        ]
    }
    return (
        <div>
            <TopPanel/>
            <div className={styles.employeeInfoBlock__wrapper}>
                <Breadcrumbs items={breadcrumbsItems}/>
                <div className={styles.employeeInfoBlock}>
                    <h2 className={styles.employeeInfoBlock__title}>Информация о сотруднике</h2>
                    <div className={styles.employeeInfoBlock__paramRows}>
                        <div className={styles.employeeInfoBlock__row}>
                            <p>ФИО</p>
                            <TextArea id={"name"}/>
                        </div>
                        <div className={styles.employeeInfoBlock__row}>
                            <p>Должность</p>
                            <TextArea id={"jobTitle"}/>
                        </div>
                        <div className={styles.employeeInfoBlock__row}>
                            <p>Телефон</p>
                            <TextArea id={"phone"}/>
                        </div>
                        <div className={styles.employeeInfoBlock__row}>
                            <p>Почта</p>
                            <TextArea id={"email"}/>
                        </div>
                        <div className={styles.employeeInfoBlock__row}>
                            <p>Пол</p>
                            <TextArea id={"email"}/>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    )
}

function renderCreateEmployeePage() {
    const branchId = window.location.search.replace("?branchId=", '');
    const root = ReactDOM.createRoot(
        document.getElementById('root') as HTMLElement
    );
    root.render(
        <React.StrictMode>
            <CreateEmployee branchId={branchId}/>
        </React.StrictMode>
    )
}

export {renderCreateEmployeePage}