import React, {ReactElement, useEffect, useState} from "react";
import {TopPanel} from "../../components/TopPanel/TopPanel";
import {fetchGetRequest} from "../../../api/fetchRequest";
import {crateEmployeeUrl, getBranchInfoAPIUrl} from "../../../api/routes";
import ReactDOM from "react-dom/client";
import styles from "./BranchInfo.module.css"
import {SideBarWrapper, SideBarWrapperProps} from "../../components/SideBar/SideBar";
import {Button, ButtonTypes} from "../../components/Button/Button";
import {AddIcon} from "../../components/Icons/Icons";
import {Breadcrumbs, BreadcrumbsItemProps} from "../../components/Breadcrumbs/Breadcrumbs";
import {Branch} from "../../../model/types";
import {TextArea} from "../../components/TextArea/TextArea";
import {updateBranchAction} from "../../../api/actionCreators";
import {EmployeeCard} from "../../components/EmployeeCard/EmployeeCard";

interface BranchInfoProps {
    branchId: string
}

const BranchInfo = (props: BranchInfoProps) => {
    const [data, setData] = useState<Branch>()
    const [loading, setLoading] = useState(true)
    useEffect(() => {
        setLoading(true)
        fetchGetRequest(getBranchInfoAPIUrl.replace("BRANCH_ID", props.branchId)).then(data => {
            setLoading(false)
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
                href: window.location.href,
                text: data?.city + ", " + data?.address,
                index: 0
            }
        ]
    }
    const sideBar: SideBarWrapperProps = {
        Elements: [{
            ElementType: ButtonTypes.Icon,
            icon: <AddIcon/>,
            active: true,
            linked: true,
            url: crateEmployeeUrl.replace("BRANCH_ID", props.branchId)
        }]
    }
    let employeeCards = data?.employees.map((employee, index) => {
        return <EmployeeCard
            key={index}
            employeeId={employee.employeeId}
            branchId={props.branchId}
            firstName={employee.firstName}
            lastName={employee.lastName}
            middleName={employee.middleName}
            jobTitle={employee.jobTitle}
        />
    })
;    return (
        <div>
            <TopPanel/>
            <div className={styles.branchInfoMainContentArea}>
                <SideBarWrapper {...sideBar}/>
                <div className={styles.branchInfoBlock}>
                    <Breadcrumbs items={breadcrumbsItems}/>
                    <div className={styles.branchInfoBlock__form}>
                        <h2 className={styles.branchInfoBlock__title}>Информация о филиале</h2>
                        <div className={styles.branchInfoBlock__paramRows}>
                            <div className={styles.branchInfoBlock__row}>
                                <p>Город</p>
                                <TextArea id={"city"} placeholder={data?.city}></TextArea>
                            </div>
                            <div className={styles.branchInfoBlock__row}>
                                <p>Aдрес</p>
                                <TextArea id={"address"} placeholder={data?.address}></TextArea>
                            </div>
                        </div>
                        <div className={styles.branchInfoBlock__employeesBlock}>
                            <p>Сотрудники:</p>
                            {employeeCards}
                        </div>
                        <Button text={"Сохранить"} viewStyle={"manageBranchPage"} onClick={() => {
                            updateBranchAction(props.branchId)
                        }}/>
                    </div>
                </div>
            </div>
        </div>
    )
}

function renderBranchInfoPage() {
    const branchId = window.location.search.replace("?branchId=", '');
    const root = ReactDOM.createRoot(
        document.getElementById('root') as HTMLElement
    );
    root.render(
        <React.StrictMode>
            <BranchInfo branchId={branchId}/>
        </React.StrictMode>
    )
}

export {renderBranchInfoPage}