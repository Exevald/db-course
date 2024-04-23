import styles from "./Button.module.css"
import {Link} from "react-router-dom";

export enum ButtonTypes {
    Icon = "Icon",
    IconText = "IconText",
    Text = "Text",
}

interface ButtonProps {
    viewStyle: "default" | "manageBranchPage"
    id?: string
    onClick?: () => void
    text: string
    to?: string
}

const Button = (props: ButtonProps) => {
    let buttonAdditionalStyle: string = ""

    switch (props.viewStyle) {
        case "manageBranchPage": {
            buttonAdditionalStyle = styles.buttonCreateUser
        }
    }

    if (props.to !== undefined) {
        return (
            <a href={props.to}>
                <button
                    type={"button"}
                    className={`${styles.button} ${buttonAdditionalStyle}`}
                    onClick={props.onClick}
                >
                    {props.text}
                </button>
            </a>
        )
    }
    return (
        <button
            type={"button"}
            className={`${styles.button} ${buttonAdditionalStyle}`}
            onClick={props.onClick}
        >
            {props.text}
        </button>
    )
}

export {Button}