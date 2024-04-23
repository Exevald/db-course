import styles from "./TopPanel.module.css"
import logo from "../../images/icons/logo.svg"

const LogoArea = () => {
    return (
        <a href={"/"}>
            <img src={logo} alt={"Main logo"} loading={"lazy"} className={styles.logoArea}/>
        </a>
    )
}

const TopPanel = () => {
    return (
        <div className={styles.topPanel}>
            <LogoArea></LogoArea>
        </div>
    )
}

export {TopPanel}