import React from 'react';
import styles from "./CreateBranch.module.css"
import {TopPanel} from "../../components/TopPanel/TopPanel";
import {Breadcrumbs, BreadcrumbsItemProps} from "../../components/Breadcrumbs/Breadcrumbs";
import {TextArea} from "../../components/TextArea/TextArea";
import {Button} from "../../components/Button/Button";
import {createBranchAction} from "../../../api/actionCreators";
import ReactDOM from "react-dom/client";

const CreateBranch = () => {
    const breadcrumbsItems: BreadcrumbsItemProps[] = [
        {
            href: "/",
            text: "Филиалы",
            index: 0
        },
        {
            href: "",
            text: "Создать филиал",
            index: 1
        }
    ]
    return (
        <div>
            <TopPanel></TopPanel>
            <div className={styles.branchInfoBlock__wrapper}>
                <Breadcrumbs items={breadcrumbsItems}/>
                <div className={styles.branchInfoBlock}>
                    <h2 className={styles.branchInfoBlock__title}>Информация о филиале</h2>
                    <div className={styles.branchInfoBlock__paramRows}>
                        <div className={styles.branchInfoBlock__row}>
                            <p>Город</p>
                            <TextArea id={"city"}></TextArea>
                        </div>
                        <div className={styles.branchInfoBlock__row}>
                            <p>Aдрес</p>
                            <TextArea id={"address"}></TextArea>
                        </div>
                    </div>
                    <Button text={"Создать"} viewStyle={"manageBranchPage"} onClick={createBranchAction}/>
                </div>
            </div>
        </div>
    )
}

function renderCreateBranchPage() {
    const root = ReactDOM.createRoot(
        document.getElementById('root') as HTMLElement
    );
    root.render(
        <React.StrictMode>
            <CreateBranch/>
        </React.StrictMode>
    )
}

export {renderCreateBranchPage}