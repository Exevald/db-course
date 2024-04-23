import styles from "./Breadcrumbs.module.css"
import chevron from "../../images/icons/chevron-right.svg"

export interface BreadcrumbsItemProps {
    index: number
    href?: string
    text: string
}

interface BreadcrumbsProps {
    items: BreadcrumbsItemProps[]
}

const BreadcrumbsItem = (props: BreadcrumbsItemProps) => {
    const linkStyles = props.index === 0 ? styles.breadcrumbsItem__link_first : styles.breadcrumbsItem__link
    return (
        <div className={styles.breadcrumbs__item}>
            {props.index !== 0 && <img src={chevron} alt={"chevron"} className={styles.breadcrumbs__chevron}/>}
            <a href={props.href}
               className={linkStyles}
            >
                {props.text}
            </a>
        </div>
    )
}

const Breadcrumbs = (props: BreadcrumbsProps) => {
    const items = props.items.map((item, index) => (
        <BreadcrumbsItem
            key={index}
            href={item.href}
            text={item.text}
            index={index}
        />
    ))
    return (
        <div className={styles.breadcrumbs}>
            {items}
        </div>
    )
}


export {Breadcrumbs}