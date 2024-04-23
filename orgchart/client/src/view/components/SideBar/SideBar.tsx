import styles from "./SideBar.module.css"
import {ButtonTypes} from "../Button/Button";


export type SideBarWrapperProps = {
    Elements: SideBarElementProps[]
}
export const SideBarWrapper = (props: SideBarWrapperProps) => {

    const Elements = props.Elements.map((Elem, i) =>
        <SideBarElement {...Elem} key={i}/>
    )

    return (<div className={styles.sidebarWrapper}>
        <div className={styles.sidebarWrapper__contentArea}>
            {Elements}
        </div>
    </div>)
}


type SideBarElementProps = {
    ElementType: ButtonTypes,
    icon?: JSX.Element,
    active?: boolean,
    linked: boolean,
    url?: string
}

const SideBarElement = (props: SideBarElementProps) => {
    switch (props.ElementType) {
        case ButtonTypes.Icon: {
            let activeMod = ""
            if (props.active) {
                activeMod = styles.sidebarElementWrapper__activePanel_active
            }
            return (
                <div className={styles.sidebarElementWrapper}>
                    <div className={`${styles.sidebarElementWrapper__activePanel} ${activeMod}`}>

                    </div>
                    <div className={styles.sidebarElementWrapper__iconArea}>
                        {props.linked && props.url &&
                            <a className={styles.linkWrapper} href={props.url}>{props.icon}</a>
                        }
                        {!props.linked && props.icon}
                    </div>
                </div>
            )
        }
        case ButtonTypes.IconText: {
            return (
                <div>

                </div>
            )
        }
        case ButtonTypes.Text: {
            return (
                <div>

                </div>
            )
        }
    }
}