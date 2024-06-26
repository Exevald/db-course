import React, {ReactElement, useEffect, useState} from 'react';
import styles from "./BranchList.module.css"
import {ViewBranch} from "../../../model/types";
import {TopPanel} from "../../components/TopPanel/TopPanel";
import {BranchCard} from "../../components/BranchCard/BranchCard";
import {SideBarWrapper, SideBarWrapperProps} from "../../components/SideBar/SideBar";
import {AddIcon} from "../../components/Icons/Icons";
import {ButtonTypes} from "../../components/Button/Button";
import {createBranchUrl, getBranchListAPIUrl} from "../../../api/routes";
import {fetchGetRequest} from "../../../api/fetchRequest";
import ReactDOM from "react-dom/client";

const sideBar: SideBarWrapperProps = {
    Elements: [{
        ElementType: ButtonTypes.Icon,
        icon: <AddIcon/>,
        active: true,
        linked: true,
        url: createBranchUrl
    }]
}

const BranchList = () => {
    const [data, setData] = useState<{ branches: ViewBranch[] }>({branches: []})
    const [branches, setBranches] = useState<ReactElement[]>([])
    const [loading, setLoading] = useState(true)
    useEffect(() => {
        setLoading(true)
        fetchGetRequest(getBranchListAPIUrl).then(data => {
            setLoading(false)
            setData(data)
        }).catch(error => console.log(error))
    }, [])
    useEffect(() => {
        setBranches(data.branches.map((branch, index) => (
            <BranchCard
                branchId={branch.branchId}
                key={index}
                city={branch.city}
                address={branch.address}
                countOfEmployees={branch.countOfEmployees}
            />
        )))
    }, [data, setData]);

    return (
        <div className={styles.branchList}>
            <TopPanel></TopPanel>
            <div className={styles.branchListMainContentArea}>
                <SideBarWrapper {...sideBar}/>
                <div className={styles.branchCardsWrapper}>
                    {loading ? <h1>LOADING...</h1> : branches}
                </div>
            </div>

        </div>
    )
}

function renderBranchListPage() {
    const root = ReactDOM.createRoot(
        document.getElementById('root') as HTMLElement
    );
    root.render(
        <React.StrictMode>
            <BranchList/>
        </React.StrictMode>
    )
}

export {renderBranchListPage}