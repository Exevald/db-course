import styles from "./BranchCard.module.css"
import {EllipsisIcon} from "../Icons/Icons";
import {PopoverContainer, PopoverProps} from "../PopoverContainer/PopoverContainer";
import {ButtonTypes} from "../Button/Button";
import {useState} from "react";
import {branchInfoUrl} from "../../../api/routes";

interface BranchCardProps {
    city: string,
    address: string,
    countOfEmployees: number
}

const popoverProps: PopoverProps = {
    Elems: [
        {
            Text: "Просмотреть",
            ElemType: ButtonTypes.Text,
            onClick: (url: string) => {
                window.location.href = url
            }
        },
        {
            Text: "Изменить",
            ElemType: ButtonTypes.Text,
            onClick: (url: string) => {
                window.location.href = url
            }
        },
        {
            Text: "Удалить",
            ElemType: ButtonTypes.Text,
            onClick: (url: string) => {
                window.location.href = url
            }
        },
    ]
}

const BranchCard = (props: BranchCardProps) => {
    const [openPopover, setOpened] = useState(false)
    return (
        <div className={styles.branchCard}>
            <div className={styles.branchInfoWrapper}>
                <div className={styles.branchInfoTitleArea}>
                    <h2>Подразделение</h2>
                    <div onClick={() => {
                        setOpened(!openPopover)
                    }} className={styles.iconContainer}>
                        <EllipsisIcon/>
                        {openPopover && <div className={styles.popover}>
                            <PopoverContainer {...popoverProps}/>
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