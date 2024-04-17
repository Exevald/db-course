import React, {useEffect, useState} from 'react';
import styles from "./BranchList.module.css"
import {Branch} from "../../../model/types";
import {TopPanel} from "../../components/TopPanel/TopPanel";
import {getBranchListAPIUrl} from "../../../api/routes";
import {fetchGetRequest} from "../../../api/fetchRequest";
import {convertAPIData} from "../../../api/convertAPIData";
import {BranchCard} from "../../components/BranchCard/BranchCard";

interface BranchListProps {
    branches: Array<Branch>
}

const mockBranchList: Branch[] = [
    {
        city: "Vologda",
        address: "Volkova 108",
        employeesList: []
    },
    {
        city: "Kazan",
        address: "Volkova 108",
        employeesList: []
    },
    {
        city: "Novgorod",
        address: "Volkova 108",
        employeesList: []
    },
    {
        city: "Moscow",
        address: "Volkova 108",
        employeesList: []
    },
]

const BranchList = (props: BranchListProps) => {
    const [data, setData] = useState(null)
    useEffect(() => {
        fetchGetRequest(getBranchListAPIUrl).then(data => {
            setData(data)
            console.log(data)
            convertAPIData(data)
        }).catch(error => console.log(error))
    }, [])
    let branches = []
    for (let i = 0; i < mockBranchList.length; i++) {
        let branch = mockBranchList[i]
        branches.push(
            <BranchCard
                city={branch.city}
                address={branch.address}
                countOfEmployees={branch.employeesList.length}
            />
        )
    }
    if (data !== null && data !== undefined) {
        return (
            <div className={styles.branchList}>
                <TopPanel></TopPanel>
                <div className={styles.branchCardsWrapper}>
                    {branches}
                </div>
            </div>
        )
    }
    return <div></div>
}

export {BranchList}