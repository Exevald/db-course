import styles from "./Popover.module.css"

import {ButtonTypes} from "../Button/Button";


export type PopoverProps = {
    Elems: PopoverElemProps[],
}

export type PopoverElemProps = {
    ElemType: ButtonTypes,
    Text?: string,
    onClick: () => void,
}


export const Popover = (props: PopoverProps) => {
    const elems = props.Elems.map((elem, i) =>
        <PopoverElem {...elem} key = {i}/>
    )
    return(
        <div className={styles.popoverContainer}>
            {elems}
        </div>
    )

}

export const PopoverElem = (props: PopoverElemProps) => {
    return(
        <div className={styles.popoverItemArea} onClick={props.onClick}>
            {props.Text}
        </div>

    )
}