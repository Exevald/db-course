import styles from "./EmployeeCard.module.css"
import {useState} from "react";
import personDefaultAvatar from "../../images/icons/default_avatar_icon.svg"
import {Popover, PopoverElemProps} from "../Popover/Popover";
import {ButtonTypes} from "../Button/Button";
import {EllipsisIcon} from "../Icons/Icons";
import {deleteEmployeeAction} from "../../../api/actionCreators";

interface EmployeeCardProps {
    employeeId: string
    branchId: string
    firstName: string
    lastName: string
    middleName: string
    jobTitle: string
}

const EmployeeCard = (props: EmployeeCardProps) => {
    console.log(props.employeeId)
    const [openPopover, setOpened] = useState(false)
    const popoverElems: PopoverElemProps[] = [
        {
            Text: "Просмотреть",
            ElemType: ButtonTypes.Text,
            onClick: () => {
            }
        },
        {
            Text: "Удалить",
            ElemType: ButtonTypes.Text,
            onClick: () => {
                deleteEmployeeAction(props.employeeId, props.branchId)
            }
        },
    ]
    return (
        <div className={styles.employeeCard}>
            <div className={styles.employeeCard__info}>
                <img src={personDefaultAvatar} alt={"person avatar"}/>
                <p>ФИО: {props.lastName + " " + props.firstName + " " + props.middleName}</p>
                <p>Должность: {props.jobTitle}</p>
            </div>
            <div onClick={() => {
                setOpened(!openPopover)
            }} className={styles.iconContainer}>
                <EllipsisIcon/>
                {openPopover && <div className={styles.popover}>
                    <Popover Elems={popoverElems}/>
                </div>}
            </div>
        </div>
    )
}

export {EmployeeCard}