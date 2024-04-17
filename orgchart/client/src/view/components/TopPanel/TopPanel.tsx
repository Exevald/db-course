import styles from "./TopPanel.module.css"
import logo from "./LogoArea.svg"
import {Link} from "react-router-dom";

const LogoArea = () => {
    return (
        <Link to={"/"}>
            <img src={logo} alt={"Main logo"} loading={"lazy"} className={styles.logoArea}/>
        </Link>
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