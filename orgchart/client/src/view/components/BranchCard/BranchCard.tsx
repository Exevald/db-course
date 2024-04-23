import styles from "./BranchCard.module.css"
import {EllipsisIcon} from "../Icons/Icons";
import {Popover, PopoverElemProps} from "../Popover/Popover";
import {ButtonTypes} from "../Button/Button";
import {useState} from "react";
import {branchInfoUrl, deleteBranchAPIUrl} from "../../../api/routes";
import {deleteBranchAction} from "../../../api/actionCreators";

interface BranchCardProps {
    branchId: string,
    city: string,
    address: string,
    countOfEmployees: number
}

const BranchCard = (props: BranchCardProps) => {
    const [openPopover, setOpened] = useState(false)
    const popoverElems: PopoverElemProps[] = [
        {
            Text: "Просмотреть",
            ElemType: ButtonTypes.Text,
            onClick: () => {
                window.location.href = branchInfoUrl.replace("BRANCH_ID", props.branchId)
            }
        },
        {
            Text: "Удалить",
            ElemType: ButtonTypes.Text,
            onClick: () => {
                deleteBranchAction(props.branchId)
            }
        },
    ]
    return (
        <div className={styles.branchCard}>
            <div className={styles.branchInfoWrapper}>
                <div className={styles.branchInfoTitleArea}>
                    <h2>Филиал</h2>
                    <div onClick={() => {
                        setOpened(!openPopover)
                    }} className={styles.iconContainer}>
                        <EllipsisIcon/>
                        {openPopover && <div className={styles.popover}>
                            <Popover Elems={popoverElems}/>
                        </div>}
                    </div>
                </div>
                <p>Город: {props.city}</p>
                <p>Адрес: {props.address}</p>
            </div>
        </div>
    )
}

export {BranchCard}