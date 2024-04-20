import styles from "./PopoverContainer.module.css"

import {ButtonTypes} from "../Button/Button";


export type PopoverProps = {
    Elems: PopoverElemProps[]
}

type PopoverElemProps = {
    ElemType: ButtonTypes,
    Text?: string,
    onClick: (url: string) => void
}


export const PopoverContainer = (props: PopoverProps) => {
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
        <div className={styles.popoverItemArea}>
            {props.Text}
        </div>

    )
}